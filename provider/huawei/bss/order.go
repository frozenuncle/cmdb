package bss

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *BssOperator) QueryOrder(req *provider.QueryOrderRequest) pager.Pager {
	return nil
}

// 客户购买包年/包月资源后,可以查看待审核、处理中、已取消、已完成和待支付等状态的订单
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=BSS&api=ListCustomerOrders
func (o *BssOperator) doQueryOrder(req *model.ListCustomerOrdersRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()

	resp, err := o.client.ListCustomerOrders(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	return set, nil
}

// 客户可以在伙伴销售平台查看订单详情
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=BSS&api=ShowCustomerOrderDetails
func (o *BssOperator) doDetailOroder(req *model.ShowCustomerOrderDetailsRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()

	resp, err := o.client.ShowCustomerOrderDetails(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	return set, nil
}

// 客户在伙伴销售平台查询某个或所有的包年/包月资源(ListPayPerUseCustomerResources)
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=BSS&api=ListPayPerUseCustomerResources