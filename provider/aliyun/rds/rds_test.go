package rds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/mcube/logger/zap"

	op "github.com/infraboard/cmdb/provider/aliyun/rds"
)

var (
	operater *op.RdsOperator
)

func TestQuery(t *testing.T) {
	req := op.NewPageQueryRequest()
	pager := operater.PageQuery(req)

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
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operater = aliyun.O().RdsOperator()
}
