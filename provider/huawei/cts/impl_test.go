package cts_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/cts"
)

var (
	operator provider.EventOperator
	ctx      = context.Background()
)

func TestPageQueryRedis(t *testing.T) {
	req := provider.NewQueryEventRequest()
	req.StartTime = time.Now().Add(-24 * 1 * time.Hour)
	pager := operator.PageQueryEvent(req)

	set := redis.NewSet()
	for pager.Next() {
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

	ec, err := connectivity.C().CtsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewOperator(ec)
}
