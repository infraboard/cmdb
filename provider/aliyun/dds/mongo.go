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

	set := mongodb.NewMongoDBSet()
	for i := range resp.Body.DBInstances.DBInstance {
		item := resp.Body.DBInstances.DBInstance[i]
		resp, err := o.client.DescribeDBInstanceAttribute(&dds.DescribeDBInstanceAttributeRequest{
			DBInstanceId: item.DBInstanceId,
		})
		if err != nil {
			o.log.Errorf("describe instance %s[%s] error, %s", tea.StringValue(item.DBInstanceDescription), tea.StringValue(item.DBInstanceId), err)
			continue
		}
		set.Add(o.transferMongoDBSet(resp.Body.DBInstances).ToAny()...)
	}

	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	return set, nil
}

func (o *Operator) transferMongoDBSet(items *dds.DescribeDBInstanceAttributeResponseBodyDBInstances) *mongodb.MongoDBSet {
	set := mongodb.NewMongoDBSet()
	for i := range items.DBInstance {
		set.Add(o.transferMongoDB(items.DBInstance[i]))
	}
	return set
}

func (o *Operator) transferMongoDB(ins *dds.DescribeDBInstanceAttributeResponseBodyDBInstancesDBInstance) *mongodb.MongoDB {
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
	info.Type = tea.StringValue(ins.DBInstanceType) + "|" + tea.StringValue(ins.DBInstanceClass) + "|" + tea.StringValue(ins.ReplicationFactor)
	info.Category = parseKindCode(ins.KindCode)
	info.Storage = tea.Int32Value(ins.DBInstanceStorage)

	r.Resource.Status.Phase = praseStatus(tea.StringValue(ins.DBInstanceStatus))
	r.Resource.Status.LockMode = tea.StringValue(ins.LockMode)
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(ins.ChargeType)

	connections := []string{}
	for i := range ins.MongosList.MongosAttribute {
		item := ins.MongosList.MongosAttribute[i]
		connections = append(connections, fmt.Sprintf("%s:%d", tea.StringValue(item.ConnectSting), tea.Int32Value(item.Port)))
	}
	for i := range ins.ReplicaSets.ReplicaSet {
		item := ins.ReplicaSets.ReplicaSet[i]
		connections = append(connections, fmt.Sprintf("%s:%s", tea.StringValue(item.ConnectionDomain), tea.StringValue(item.ConnectionPort)))
	}
	r.Resource.Status.PrivateAddress = connections

	desc := r.Describe
	desc.Engine = tea.StringValue(ins.Engine)
	desc.MaxConnections = tea.Int32Value(ins.MaxConnections)
	desc.MaxIops = tea.Int32Value(ins.MaxIOPS)
	desc.EngineVersion = tea.StringValue(ins.EngineVersion)
	return r
}

func parseKindCode(c *string) string {
	if c == nil {
		return ""
	}

	switch *c {
	case "0":
		return "物理机"
	case "1":
		return "ECS"
	case "2":
		return "DOCKER"
	case "18":
		return "k8s新架构实例"
	}

	return ""
}
