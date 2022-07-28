package cos_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.OssOperator
)

func TestQuery(t *testing.T) {
	pager := operator.QueryBucket(context.Background(), provider.NewQueryBucketRate(5))

	for pager.Next() {
		set := oss.NewBucketSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()

	err := txyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}

	operator = txyun.O().OssOperator()
}
