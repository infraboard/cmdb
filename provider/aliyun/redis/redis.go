package redis

import (
	"context"
	"time"

	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	cmdbRedis "github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/provider"
)

func (o *RedisOperator) DescribeRedis(ctx context.Context, req *provider.DescribeRequest) (
	*cmdbRedis.Redis, error) {
	descReq := &redis.DescribeInstanceAttributeRequest{
		InstanceId: &req.Id,
	}

	detail, err := o.client.DescribeInstanceAttribute(descReq)
	if err != nil {
		return nil, err
	}

	set := o.transferAttrSet(detail.Body.Instances)
	if set.Length() == 0 {
		return nil, exception.NewNotFound("ins %s not found", req.Id)
	}

	return set.Items[0], nil
}

func (o *RedisOperator) PageQueryRedis(req *provider.QueryRedisRequest) pager.Pager {
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

	b := r.Base
	b.Vendor = resource.VENDOR_ALIYUN
	b.Region = tea.StringValue(ins.RegionId)
	b.Zone = tea.StringValue(ins.ZoneId)
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreateTime))
	b.Id = tea.StringValue(ins.InstanceId)

	info := r.Information
	info.ExpireAt = o.parseTime(tea.StringValue(ins.EndTime))
	info.Name = tea.StringValue(ins.InstanceName)
	info.Type = tea.StringValue(ins.EditionType)
	info.Category = tea.StringValue(ins.ArchitectureType)
	info.Status = praseStatus(ins.InstanceStatus)
	info.PayType = tea.StringValue(ins.ChargeType)
	info.PrivateIp = []string{tea.StringValue(ins.PrivateIp)}

	desc := r.Describe
	desc.Memory = tea.Int64Value(ins.Capacity)
	desc.ConnectAddr = tea.StringValue(ins.ConnectionDomain)
	desc.ConnectPort = tea.Int64Value(ins.Port)
	desc.ArchitectureType = tea.StringValue(ins.ArchitectureType)
	desc.Qps = tea.Int64Value(ins.QPS)
	desc.BandWidth = tea.Int64Value(ins.Bandwidth)
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

func (o *RedisOperator) transferAttrSet(items *redis.DescribeInstanceAttributeResponseBodyInstances) *cmdbRedis.Set {
	set := cmdbRedis.NewSet()
	for i := range items.DBInstanceAttribute {
		set.Add(o.transferAttrOne(items.DBInstanceAttribute[i]))
	}
	return set
}

func (o *RedisOperator) transferAttrOne(ins *redis.DescribeInstanceAttributeResponseBodyInstancesDBInstanceAttribute) *cmdbRedis.Redis {
	r := cmdbRedis.NewDefaultRedis()

	b := r.Base
	b.Vendor = resource.VENDOR_ALIYUN
	b.Region = tea.StringValue(ins.RegionId)
	b.Zone = tea.StringValue(ins.ZoneId)
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreateTime))
	b.Id = tea.StringValue(ins.InstanceId)

	return r
}
