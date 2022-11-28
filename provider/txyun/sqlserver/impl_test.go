package sqlserver_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	op "github.com/infraboard/cmdb/provider/txyun/sqlserver"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator *op.SQLServerOperator
	ctx      = context.Background()
)

func TestPageQueryRds(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryRds(req)

	for pager.Next() {
		set := rds.NewSet()
		if err := pager.Scan(ctx, set); err != nil {
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
	operator = op.NewSQLServerOperator(client.SQLServerClient())
}
