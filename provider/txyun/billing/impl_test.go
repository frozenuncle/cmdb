package billing_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/order"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/txyun/billing"
)

var (
	operator provider.BillOperator
)

func TestQueryBill(t *testing.T) {
	req := provider.NewQueryBillRequest()
	req.Month = "2022-04"

	pager := operator.QueryBill(req)
	for pager.Next() {
		set := bill.NewBillSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestQueryOrder(t *testing.T) {
	req := provider.NewQueryOrderRequest()
	req.StartTime = time.Now().Add(-24 * time.Hour)
	pager := operator.QueryOrder(req)
	for pager.Next() {
		set := order.NewOrderSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()

	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	client := connectivity.C()
	operator = op.NewBillingoperator(client.BillingClient())
}
