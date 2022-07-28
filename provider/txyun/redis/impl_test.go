package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	op "github.com/infraboard/cmdb/provider/txyun/redis"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator *op.RedisOperator
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRedisWithRate(5)
	pager := operator.QueryRedis(req)

	for pager.Next() {
		set := redis.NewSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		for i := range set.Items {
			fmt.Println(set.Items[i])
		}
	}
}

func init() {
	zap.DevelopmentSetup()
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	client := connectivity.C()
	operator = op.NewRedisOperator(client.RedisClient())
}
