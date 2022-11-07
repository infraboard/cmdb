package redis

import (
	"context"
	"fmt"
	"time"

	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	cmdbRedis "github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/provider"
)

func (o *RedisOperator) DescribeRedis(ctx context.Context, r *provider.DescribeRequest) (*cmdbRedis.Redis, error) {
	req := &redis.DescribeInstancesRequest{
		RegionId:    o.client.RegionId,
		InstanceIds: tea.String(r.Id),
		PageSize:    tea.Int32(1),
	}

	set, err := o.query(req)
	if err != nil {
		return nil, err
	}

	if set.Length() == 0 {
		return nil, fmt.Errorf("redis %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *RedisOperator) PageQueryRedis(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询一个或多个Redis实例的信息
// 参考文档: https://next.api.aliyun.com/api/R-kvstore/2015-01-01/DescribeInstances?params={}
func (o *RedisOperator) query(req *redis.DescribeInstancesRequest) (*cmdbRedis.Set, error) {
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	set := cmdbRedis.NewSet()
	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	set.Items = o.transferSet(resp.Body.Instances).Items
	return set, nil
}

func (o *RedisOperator) transferSet(items *redis.DescribeInstancesResponseBodyInstances) *cmdbRedis.Set {
	set := cmdbRedis.NewSet()
	for i := range items.KVStoreInstance {
		set.Add(o.transferOne(items.KVStoreInstance[i]))
	}
	return set
}

func (o *RedisOperator) transferOne(ins *redis.DescribeInstancesResponseBodyInstancesKVStoreInstance) *cmdbRedis.Redis {
	r := cmdbRedis.NewDefaultRedis()

	b := r.Resource.Meta
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreateTime))
	b.Id = tea.StringValue(ins.InstanceId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_ALIYUN
	info.Region = tea.StringValue(ins.RegionId)
	info.Zone = tea.StringValue(ins.ZoneId)
	info.ExpireAt = o.parseTime(tea.StringValue(ins.EndTime))
	info.Name = tea.StringValue(ins.InstanceName)
	info.Type = tea.StringValue(ins.EditionType)
	info.Category = tea.StringValue(ins.ArchitectureType)
	info.Memory = int32(tea.Int64Value(ins.Capacity))
	info.BandWidth = int32(tea.Int64Value(ins.Bandwidth))

	r.Resource.Status.PrivateAddress = []string{tea.StringValue(ins.PrivateIp)}
	r.Resource.Status.Phase = praseStatus(ins.InstanceStatus)
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(ins.ChargeType)

	desc := r.Describe
	desc.ConnectAddr = tea.StringValue(ins.ConnectionDomain)
	desc.ConnectPort = tea.Int64Value(ins.Port)
	desc.ArchitectureType = tea.StringValue(ins.ArchitectureType)
	desc.Qps = tea.Int64Value(ins.QPS)
	desc.MaxConnection = tea.Int64Value(ins.Connections)
	desc.Config = tea.StringValue(ins.Config)
	desc.NodeType = tea.StringValue(ins.NodeType)
	desc.NetworkType = tea.StringValue(ins.NetworkType)
	desc.EngineType = tea.StringValue(ins.InstanceType)
	desc.EngineVersion = tea.StringValue(ins.EngineVersion)
	desc.ReplicaId = tea.StringValue(ins.ReplacateId)
	return r
}

func (o *RedisOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
