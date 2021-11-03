package impl

import (
	"context"

	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/pkg/task"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
)

// 通过回调更新任务状态
func (s *service) SyncTaskCallback(t *task.Task) {
	err := s.save(context.Background(), t)
	if err != nil {
		s.log.Error("save task error, %s", err)
	}
}

func (s *service) CreatTask(ctx context.Context, req *task.CreateTaskRequst) (
	*task.Task, error) {
	t, err := task.NewTaskFromReq(req)
	if err != nil {
		return nil, err
	}

	secret, err := s.secret.DescribeSecret(ctx, secret.NewDescribeSecretRequest(req.SecretId))
	if err != nil {
		return nil, err
	}

	// 如果不是vsphere 需要检查region
	if !secret.Vendor.Equal(resource.VendorVsphere) {
		if req.Region == "" {
			return nil, exception.NewBadRequest("region required")
		}
		if !secret.IsAllowRegion(req.Region) {
			return nil, exception.NewBadRequest("this secret not allow sync region %s", req.Region)
		}
	}

	// 资源同步
	switch req.ResourceType {
	case resource.HostResource:
		go s.syncHost(ctx, secret, t, s.SyncTaskCallback)
	}

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *service) QueryTask(ctx context.Context, req *task.QueryTaskRequest) (*task.TaskSet, error) {
	query := sqlbuilder.NewQuery(queryTaskSQL)

	querySQL, args := query.Order("start_at").Desc().Limit(req.OffSet(), uint(req.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query task error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := task.NewTaskSet()
	for rows.Next() {
		ins := task.NewDefaultTask()
		err := rows.Scan(
			&ins.Id, &ins.Region, &ins.ResourceType, &ins.SecretId, &ins.Timeout, &ins.Status,
			&ins.Message, &ins.StartAt, &ins.EndAt, &ins.TotolSucceed, &ins.TotalFailed,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query task error, %s", err.Error())
		}
		set.Add(ins)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	return set, nil
}