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
	ctx      = context.Background()
)

func TestPageQueryHost(t *testing.T) {
	pager := operator.PageQueryHost(provider.NewQueryRequest())

	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestPageQueryDisk(t *testing.T) {
	pager := operator.PageQueryDisk(provider.NewQueryRequest())

	for pager.Next() {
		set := disk.NewDiskSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeDisk(t *testing.T) {
	req := provider.NewDescribeRequest("xxx")
	ins, err := operator.DescribeDisk(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestPageQueryEip(t *testing.T) {
	pager := operator.PageQueryEip(provider.NewQueryRequest())

	for pager.Next() {
		set := eip.NewEIPSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeHost(t *testing.T) {
	req := &provider.DescribeRequest{Id: "xxxx"}
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
