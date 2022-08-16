package dcs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/dcs"
)

var (
	operator provider.RedisOperator
	ctx      = context.Background()
)

func TestPageQueryRedis(t *testing.T) {
	req := provider.NewQueryRedisWithRate(5)
	pager := operator.PageQueryRedis(req)

	set := redis.NewSet()
	for pager.Next() {
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribeRedis(t *testing.T) {
	req := provider.NewDescribeRequest("xxx")
	ins, err := operator.DescribeRedis(ctx, req)
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

	ec, err := connectivity.C().DcsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewDcsOperator(ec)
}
