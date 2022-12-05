package obs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/obs"
)

var (
	operator provider.OssOperator
	ctx      = context.Background()
)

func TestPageQueryRds(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.QueryBucket(ctx, req)

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
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	ec, err := connectivity.C().ObsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewObsOperator(ec)
}
