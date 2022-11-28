package actiontrail_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.EventOperator
	ctx      = context.Background()
)

func TestPageQueryEvent(t *testing.T) {
	req := provider.NewQueryEventRequest()
	req.StartTime = time.Now().Add(-24 * 10 * time.Hour)
	pager := operator.PageQueryEvent(req)

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

	operator = aliyun.O().EventOperator()
}
