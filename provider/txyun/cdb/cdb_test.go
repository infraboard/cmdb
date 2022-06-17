package cdb_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/txyun/cdb"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator *op.CDBOperator
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

	client := connectivity.C()
	operator = op.NewCDBOperator(client.CDBClient())
}
