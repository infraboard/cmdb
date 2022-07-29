package clb_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/txyun/clb"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator *op.CLBOperator
)

func TestQueryCLB(t *testing.T) {
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

	client := connectivity.C()
	operator = op.NewCLBOperator(client.ClbClient())
}
