package oss_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider/aliyun"
)

var (
	operator provider.OssOperator
	ctx      = context.Background()
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRequestWithRate(5)
	pager := operator.QueryBucket(ctx, req)

	set := oss.NewBucketSet()
	for pager.Next() {
		if err := pager.Scan(ctx, set); err != nil {
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
	operator = aliyun.O().OssOperator()
}
