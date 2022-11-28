package dds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/mongodb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.MongoOperator
	ctx      = context.Background()
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryMongo(req)

	for pager.Next() {
		set := mongodb.NewMongoDBSet()
		if err := pager.Scan(ctx, set); err != nil {
			panic(err)
		}
		for i := range set.Items {
			fmt.Println(set.Items[i])
		}
	}
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().MongoOperator()
}
