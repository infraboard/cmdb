package mongo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/mongodb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	op "github.com/infraboard/cmdb/provider/txyun/mongo"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator *op.MongoOperator
	ctx      = context.Background()
)

func TestPageQueryMongoDB(t *testing.T) {
	req := provider.NewQueryRequestWithRate(5)
	pager := operator.PageQueryMongoDB(req)

	for pager.Next() {
		set := mongodb.NewMongoDBSet()
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

	client := connectivity.C()
	operator = op.NewMongoOperator(client.MongoClient())
}
