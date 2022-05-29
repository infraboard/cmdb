package ecs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/huawei/ecs"
)

var (
	operator *op.EcsOperator
)

func TestQuery(t *testing.T) {
	pager := operator.PageQuery(op.NewPageQueryRequest(1))

	for pager.Next() {
		set := host.NewHostSet()
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

	ec, err := connectivity.C().EcsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewEcsOperator(ec)
}
