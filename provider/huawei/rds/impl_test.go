package rds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/rds"
)

var (
	operator provider.RdsOperator
	ctx      = context.Background()
)

func TestPageQueryRds(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryRds(req)

	for pager.Next() {
		set := rds.NewSet()
		if err := pager.Scan(ctx, set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeRds(t *testing.T) {
	req := provider.NewDescribeRequest("xxxx")
	ins, err := operator.DescribeRds(ctx, req)
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

	ec, err := connectivity.C().RdsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewRdsOperator(ec)
}
