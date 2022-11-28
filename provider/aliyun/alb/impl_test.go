package alb_test

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

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().ALbOperator()
}
