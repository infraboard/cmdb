package dds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/mongodb"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/dds"
)

var (
	operator provider.MongoOperator
	ctx      = context.Background()
)

func TestPageQueryRedis(t *testing.T) {
	req := provider.NewQueryRequestWithRate(5)
	pager := operator.PageQueryMongo(req)

	set := mongodb.NewMongoDBSet()
	for pager.Next() {
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		for i := range set.Items {
			fmt.Println(set.Items[i])
		}
	}
}

func init() {
	zap.DevelopmentSetup()
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	ec, err := connectivity.C().DdsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewDdsOperator(ec)
}
