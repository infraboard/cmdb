package redis

import (
	"time"

	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	cmdbRedis "github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRedisOperator(client *redis.Client) *RedisOperator {
	return &RedisOperator{
		client: client,
		log:    zap.L().Named("ALI Redis"),
	}
}

type RedisOperator struct {
	client *redis.Client
	log    logger.Logger
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
	b.Vendor = resource.Vendor_ALIYUN
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
	b.Vendor = resource.Vendor_ALIYUN
	b.Region = tea.StringValue(ins.RegionId)
	b.Zone = tea.StringValue(ins.ZoneId)
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreateTime))
	b.Id = tea.StringValue(ins.InstanceId)

	return r
}
