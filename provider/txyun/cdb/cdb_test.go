package cdb_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	op "github.com/infraboard/cmdb/provider/txyun/cdb"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operater *op.CDBOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery()

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
	err = client.Check()
	if err != nil {
		panic(err)
	}

	operater = op.NewCDBOperater(client.CDBClient())
}
