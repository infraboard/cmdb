package ecs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/eip"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider/aliyun"
)

var (
	operator provider.HostOperator
)

func TestQueryEcs(t *testing.T) {
	req := provider.NewQueryHostRequest()
	pager := operator.QueryHost(req)
	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeEcs(t *testing.T) {
	req := &provider.DescribeHostRequest{Id: "i-bp1ieov2ieftvjai48nf"}
	ins, err := operator.DescribeHost(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryDisk(t *testing.T) {
	req := provider.NewQueryDiskRequest()
	pager := operator.QueryDisk(req)
	for pager.Next() {
		set := disk.NewDiskSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}

		fmt.Println(set)
	}
}

func TestQueryEip(t *testing.T) {
	req := provider.NewQueryEipRequest()
	pager := operator.QueryEip(req)
	for pager.Next() {
		set := eip.NewEIPSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}

		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().HostOperator()
}
