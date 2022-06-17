package rds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/rds"
)

var (
	operator provider.RdsOperator
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRdsRequest()
	pager := operator.QueryRds(req)

	for pager.Next() {
		set := rds.NewSet()
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

	ec, err := connectivity.C().RdsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewRdsOperator(ec)
}
