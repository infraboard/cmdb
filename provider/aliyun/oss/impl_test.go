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
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryBucketRate(5)
	pager := operator.QueryBucket(context.Background(), req)

	set := oss.NewBucketSet()
	for pager.Next() {
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
	operator = aliyun.O().OssOperator()
}
