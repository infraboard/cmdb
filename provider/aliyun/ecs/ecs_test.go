package ecs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider/aliyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/aliyun/ecs"
)

var (
	operater *op.EcsOperater
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	pager := operater.PageQuery(req)
	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribe(t *testing.T) {
	ins, err := operater.Describe(&op.DescribeRequest{Id: "i-bp1f6d1sbq8s9mm59jeu"})
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
	operater = op.NewEcsOperater(ec)
}
