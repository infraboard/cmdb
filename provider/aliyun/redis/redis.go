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
