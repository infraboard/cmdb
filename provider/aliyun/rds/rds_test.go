package rds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/mcube/logger/zap"
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
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().RdsOperator()
}
