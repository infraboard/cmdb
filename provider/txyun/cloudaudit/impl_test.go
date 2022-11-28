package cloudaudit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/txyun/cloudaudit"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator *op.Operator
	ctx      = context.Background()
)

func TestPageQueryLoadBalancer(t *testing.T) {
	req := provider.NewQueryEventRequest()
	req.StartTime = time.Now().Add(-24 * 30 * time.Hour)
	pager := operator.PageQueryEvents(req)

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
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	client := connectivity.C()
	operator = op.NewOperator(client.AuditClient())
}
