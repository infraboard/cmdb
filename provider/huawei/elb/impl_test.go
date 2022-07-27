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
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryLBRequest()
	pager := operator.QueryLB(req)

	for pager.Next() {
		set := lb.NewLBSet()
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

	ec, err := connectivity.C().ElbClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewELBOperator(ec)
}
