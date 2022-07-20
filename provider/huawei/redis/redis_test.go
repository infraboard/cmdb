package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/redis"
)

var (
	operator provider.RedisOperator
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRedisWithRate(5)
	pager := operator.QueryRedis(req)

	for pager.Next() {
		set := redis.NewSet()
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

	ec, err := connectivity.C().DcsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewDcsOperator(ec)
}
