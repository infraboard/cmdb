package elb_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/elb"
)

var (
	operator provider.LBOperator
	ctx      = context.Background()
)

func TestPageQueryLoadBalancer(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryLoadBalancer(req)

	for pager.Next() {
		set := lb.NewLoadBalancerSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeLoadBalancer(t *testing.T) {
	req := provider.NewDescribeRequest("xxx")
	ins, err := operator.DescribeLoadBalancer(ctx, req)
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

	ec, err := connectivity.C().ElbClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewELBOperator(ec)
}
