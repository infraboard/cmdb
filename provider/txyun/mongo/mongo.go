package mongo

import (
	"context"
	"fmt"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/mongodb"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	mongo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"
)

func (o *MongoOperator) PageQueryMongoDB(req *provider.QueryRequest) pager.Pager {
	return newPager(20, o)
}

// 查询实例列表 https://console.cloud.tencent.com/api/explorer?Product=mongodb&Version=2019-07-25&Action=DescribeDBInstances&SignVersion=
// 查询实例客户端连接信息 https://console.cloud.tencent.com/api/explorer?Product=mongodb&Version=2019-07-25&Action=DescribeClientConnections
func (o *MongoOperator) Query(ctx context.Context, req *mongo.DescribeDBInstancesRequest) (*mongodb.MongoDBSet, error) {
	resp, err := o.client.DescribeDBInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.ToJsonString())
	set := o.transferMongoDBSet(ctx, resp.Response.InstanceDetails)

	// 补充连接信息
	for i := range set.Items {
		item := set.Items[i]
		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := mongo.NewDescribeClientConnectionsRequest()
		request.InstanceId = &item.Resource.Meta.Id
		resp, err := o.client.DescribeClientConnectionsWithContext(ctx, request)
		if err != nil {
			return nil, err
		}
		fmt.Println(resp.ToJsonString())
		// connections := []string{}
		// for i := range ins.MongosList.MongosAttribute {
		// 	item := ins.MongosList.MongosAttribute[i]
		// 	connections = append(connections, fmt.Sprintf("%s:%d", tea.StringValue(item.ConnectSting), tea.Int32Value(item.Port)))
		// }
		// for i := range ins.ReplicaSets.ReplicaSet {
		// 	item := ins.ReplicaSets.ReplicaSet[i]
		// 	connections = append(connections, fmt.Sprintf("%s:%s", tea.StringValue(item.ConnectionDomain), tea.StringValue(item.ConnectionPort)))
		// }
		// r.Resource.Status.PrivateAddress = connections
	}

	return set, nil
}

func (o *MongoOperator) transferMongoDBSet(ctx context.Context, items []*mongo.InstanceDetail) *mongodb.MongoDBSet {
	set := mongodb.NewMongoDBSet()
	for i := range items {
		set.Add(o.transferMongoDBOne(ctx, items[i]))
	}
	return set
}

func (o *MongoOperator) transferMongoDBOne(ctx context.Context, ins *mongo.InstanceDetail) *mongodb.MongoDB {
	r := mongodb.NewDefaultMongoDB()
	b := r.Resource.Meta
	b.CreateAt = utils.ParseSecondMod1Time(tea.StringValue(ins.CreateTime))
	b.Id = tea.StringValue(ins.InstanceId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_TENCENT
	info.Region = tea.StringValue(ins.Region)
	info.Zone = tea.StringValue(ins.Zone)
	info.ExpireAt = utils.ParseSecondMod1Time(tea.StringValue(ins.DeadLine))
	info.Name = tea.StringValue(ins.InstanceName)
	// 实例分片数 * (实例从节点数 + 1)
	rs := tea.Uint64Value(ins.ReplicationSetNum) * (tea.Uint64Value(ins.SecondaryNum) + 1)

	info.Type = ParseClusterType(ins.ClusterType) + "|" + tea.StringValue(ins.MachineType) + "|" + fmt.Sprintf("%d", rs)
	info.Storage = int32(tea.Uint64Value(ins.Volume))
	info.Cpu = int32(tea.Uint64Value(ins.CpuNum))
	info.Memory = int32(tea.Uint64Value(ins.Memory))

	r.Resource.Status.Phase = praseStatus(tea.Int64Value(ins.Status))
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(fmt.Sprintf("%d", tea.Uint64Value(ins.PayMode)))

	desc := r.Describe
	mv := tea.StringValue(ins.MongoVersion)
	desc.Engine = strings.Split(mv, "_")[0]
	desc.EngineVersion = mv
	return r
}

func ParseClusterType(t *uint64) string {
	if t != nil {
		switch v := tea.Uint64Value((t)); v {
		case 0:
			return "副本集实例"
		case 1:
			return "分片实例"
		}
	}

	return ""
}
