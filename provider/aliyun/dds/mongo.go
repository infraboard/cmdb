package dds

import (
	"fmt"

	dds "github.com/alibabacloud-go/dds-20151201/v3/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/mongodb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *Operator) PageQueryMongo(req *provider.QueryRequest) pager.Pager {
	p := newMongoPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询MongoDB实例列表
// 参考: https://next.api.aliyun.com/api/Dds/2015-12-01/DescribeDBInstances?params={}
func (o *Operator) Query(req *dds.DescribeDBInstancesRequest) (*mongodb.MongoDBSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	fmt.Print(resp.String())

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
	b := r.Resource.Meta
	b.CreateAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.CreationTime))
	b.Id = tea.StringValue(ins.DBInstanceId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_ALIYUN
	info.Region = tea.StringValue(ins.RegionId)
	info.Zone = tea.StringValue(ins.ZoneId)
	info.ExpireAt = utils.ParseDefaultMiniteTime(tea.StringValue(ins.ExpireTime))
	info.Name = tea.StringValue(ins.DBInstanceDescription)
	info.Type = tea.StringValue(ins.DBInstanceClass)
	info.Category = tea.StringValue(ins.KindCode)
	r.Resource.Status.Phase = tea.StringValue(ins.DBInstanceStatus)
	r.Resource.Cost.PayMode = mapping.PrasePayMode(ins.ChargeType)

	return r
}
