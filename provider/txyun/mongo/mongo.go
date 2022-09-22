package mongo

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/mongodb"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
	mongo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"
)

func (o *MongoOperator) PageQueryMongoDB(req *provider.QueryRequest) pager.Pager {
	return newPager(20, o)
}

// 查询云数据库实例列表
// 参考文档: https://console.cloud.tencent.com/api/explorer?Product=mongodb&Version=2019-07-25&Action=DescribeDBInstances&SignVersion=
func (o *MongoOperator) Query(ctx context.Context, req *mongo.DescribeDBInstancesRequest) (*mongodb.MongoDBSet, error) {
	resp, err := o.client.DescribeDBInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.ToJsonString())

	return o.transferMongoDBSet(resp.Response.InstanceDetails), nil
}

func (o *MongoOperator) transferMongoDBSet(items []*mongo.InstanceDetail) *mongodb.MongoDBSet {
	set := mongodb.NewMongoDBSet()
	for i := range items {
		set.Add(o.transferMongoDBOne(items[i]))
	}
	return set
}

func (o *MongoOperator) transferMongoDBOne(ins *mongo.InstanceDetail) *mongodb.MongoDB {
	r := mongodb.NewDefaultMongoDB()

	return r
}
