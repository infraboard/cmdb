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
	ctx      = context.Background()
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryRds(req)

	for pager.Next() {
		set := rds.NewSet()
		if err := pager.Scan(ctx, set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribe(t *testing.T) {
	req := provider.NewDescribeRequest("xxx")
	ins, err := operator.DescribeRds(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().RdsOperator()
}
