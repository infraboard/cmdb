package slb_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.LoadBalancerOperator
	ctx      = context.Background()
)

func TestPageQueryLoadBalancer(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryLoadBalancer(req)

	for pager.Next() {
		set := lb.NewLoadBalancerSet()
		if err := pager.Scan(ctx, set); err != nil {
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
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().LbOperator()
}
