package ecs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/eip"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/huawei/ecs"
)

var (
	operator *op.EcsOperator
)

func TestQueryEcs(t *testing.T) {
	pager := operator.QueryHost(provider.NewQueryHostRequest())

	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestQueryDisk(t *testing.T) {
	pager := operator.QueryDisk(provider.NewQueryDiskRequest())

	for pager.Next() {
		set := disk.NewDiskSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestQueryEip(t *testing.T) {
	pager := operator.QueryEip(provider.NewQueryEipRequest())

	for pager.Next() {
		set := eip.NewEIPSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeHost(t *testing.T) {
	req := &provider.DescribeHostRequest{Id: "5f55a412-cd3c-4144-82ce-5001c2b2f08c"}
	ins, err := operator.DescribeHost(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	zap.DevelopmentSetup()
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	ec, err := connectivity.C().EcsClient()
	if err != nil {
		panic(err)
	}
	ev, err := connectivity.C().EvsClient()
	if err != nil {
		panic(err)
	}
	ep, err := connectivity.C().EipClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewEcsOperator(ec, ev, ep)
}
