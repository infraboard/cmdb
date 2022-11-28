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
	ctx      = context.Background()
)

func TestQuery(t *testing.T) {
	pager := operator.QueryBucket(ctx, provider.NewQueryRequestWithRate(5))

	for pager.Next() {
		set := oss.NewBucketSet()
		if err := pager.Scan(ctx, set); err != nil {
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
