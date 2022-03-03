package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/cmdb/apps/task"

	aliConn "github.com/infraboard/cmdb/provider/aliyun/connectivity"
	ecsOp "github.com/infraboard/cmdb/provider/aliyun/ecs"
	awsConn "github.com/infraboard/cmdb/provider/aws/connectivity"
	ec2Op "github.com/infraboard/cmdb/provider/aws/ec2"
	hwConn "github.com/infraboard/cmdb/provider/huawei/connectivity"
	hwEcsOp "github.com/infraboard/cmdb/provider/huawei/ecs"
	txConn "github.com/infraboard/cmdb/provider/txyun/connectivity"
	cvmOp "github.com/infraboard/cmdb/provider/txyun/cvm"
	vsConn "github.com/infraboard/cmdb/provider/vsphere/connectivity"
	vmOp "github.com/infraboard/cmdb/provider/vsphere/vm"
)

func (s *service) syncHost(ctx context.Context, secretIns *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager host.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		if err := recover(); err != nil {
			t.Failed(fmt.Sprintf("pannic, %v", err))
		} else {
			t.Completed()
		}
		cb(t)
	}()

	secret := secretIns.Data
	switch secret.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun ecs ...")
		client := aliConn.NewAliCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		if err := client.Check(); err != nil {
			t.Failed(err.Error())
			return
		}
		ec, err := client.EcsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := ecsOp.NewEcsOperater(ec)
		operater.WithAccountId(client.AccountID())
		req := ecsOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun cvm ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		if err := client.Check(); err != nil {
			t.Failed(err.Error())
			return
		}
		operater := cvmOp.NewCVMOperater(client.CvmClient())
		operater.WithAccountId(client.AccountID())
		req := cvmOp.NewPageQueryRequest(int(secret.RequestRate))
		pager = operater.PageQuery(req)
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun ecs ...")
		client := hwConn.NewHuaweiCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		if err := client.Check(); err != nil {
			t.Failed(err.Error())
			return
		}
		ec, err := client.EcsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := hwEcsOp.NewEcsOperater(ec)
		operater.WithAccountId(client.AccountID())
		pager = operater.PageQuery()
	case resource.Vendor_AMAZON:
		s.log.Debugf("sync aws ec2 ...")
		client := awsConn.NewAwsCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		ec, err := client.Ec2Client()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := ec2Op.NewEc2Operator(ec)
		req := ec2Op.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_VSPHERE:
		s.log.Debugf("sync vshpere vm ...")
		client := vsConn.NewVsphereClient(secret.Address, secret.ApiKey, secret.ApiSecret)
		ec, err := client.VimClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := vmOp.NewVmOperater(ec)
		// 通过回调直接保存
		err = operater.Query(func(h *host.Host) {
			// 补充管理信息
			h.Base.SecretId = secretIns.Id
			s.doSyncHost(ctx, h, t)
		})
		if err != nil {
			t.Failed(err.Error())
			return
		}
	default:
		t.Failed(fmt.Sprintf("unsuport vendor %s", secret.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		for pager.HasNext() {
			p := pager.Next()
			if p.Err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", p.Err))
				return
			}

			// 调用host服务保持数据
			for i := range p.Data.Items {
				target := p.Data.Items[i]
				// 补充管理信息
				target.Base.SecretId = secretIns.Id
				s.doSyncHost(ctx, target, t)
			}
		}
	}
}

// Host主机数据入库
func (s *service) doSyncHost(ctx context.Context, ins *host.Host, t *task.Task) {
	h, err := s.host.SyncHost(ctx, ins)

	var detail *task.Record
	if err != nil {
		s.log.Warnf("save host error, %s", err)
		detail = task.NewSyncFailedRecord(t.Id, ins.Base.Id, ins.Information.Name, err.Error())
	} else {
		s.log.Debugf("save host %s to db", h.ShortDesc())
		detail = task.NewSyncSucceedRecord(t.Id, ins.Base.Id, ins.Information.Name)
	}

	t.AddDetail(detail)
	if err := s.insertTaskDetail(ctx, detail); err != nil {
		s.log.Errorf("update detail error, %s", err)
	}
}
