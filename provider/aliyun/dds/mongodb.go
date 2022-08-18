package dds

import (
	dds "github.com/alibabacloud-go/dds-20151201/v3/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/mongodb"
)

// 查询MongoDB实例列表
// 参考: https://next.api.aliyun.com/api/Dds/2015-12-01/DescribeDBInstances?params={}
func (o *Operator) Query(req *dds.DescribeDBInstancesRequest) (*mongodb.MongoDBSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	set := mongodb.NewMongoDBSet()
	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	set.Items = o.transferMongoDBSet(resp.Body.DBInstances).Items
	return set, nil
}

func (o *Operator) transferMongoDBSet(items *dds.DescribeDBInstancesResponseBodyDBInstances) *mongodb.MongoDBSet {
	set := mongodb.NewMongoDBSet()
	for i := range items.DBInstance {
		set.Add(o.transferMongoDB(items.DBInstance[i]))
	}

	return set
}

func (o *Operator) transferMongoDB(ins *dds.DescribeDBInstancesResponseBodyDBInstancesDBInstance) *mongodb.MongoDB {
	r := mongodb.NewDefaultMongoDB()

	return r
}
