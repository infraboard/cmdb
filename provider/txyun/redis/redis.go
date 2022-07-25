package redis

import (
	"time"

	"github.com/alibabacloud-go/tea/tea"
	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"

	cmdbRedis "github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRedisOperator(client *redis.Client) *RedisOperator {
	return &RedisOperator{
		client: client,
		log:    zap.L().Named("tx.redis"),
	}
}

type RedisOperator struct {
	client *redis.Client
	log    logger.Logger
}

func (o *RedisOperator) transferSet(items *redis.DescribeInstancesResponseParams) *cmdbRedis.Set {
	set := cmdbRedis.NewSet()
	for i := range items.InstanceSet {
		set.Add(o.transferOne(items.InstanceSet[i]))
	}
	return set
}

func (o *RedisOperator) transferOne(ins *redis.InstanceSet) *cmdbRedis.Redis {
	r := cmdbRedis.NewDefaultRedis()
	b := r.Base
	b.Vendor = resource.Vendor_TENCENT
	b.Region = tea.StringValue(ins.Region)
	b.CreateAt = o.parseTime(tea.StringValue(ins.Createtime))
	b.Id = tea.StringValue(ins.InstanceId)

	info := r.Information
	info.ExpireAt = o.parseTime(tea.StringValue(ins.DeadlineTime))
	info.Category = tea.StringValue(ins.ProductType)
	info.Type = o.ParseType(ins.Type)
	info.Name = tea.StringValue(ins.InstanceName)
	info.Status = praseStatus(ins.Status)
	info.PayType = o.ParseBillMode(ins.BillingMode)
	info.PrivateIp = []string{tea.StringValue(ins.WanIp)}

	desc := r.Describe
	desc.Memory = int64(tea.Float64Value(ins.Size))
	desc.BandWidth = tea.Int64Value(ins.NetLimit)
	desc.MaxConnection = tea.Int64Value(ins.ClientLimitMax)
	desc.EngineType = tea.StringValue(ins.Engine)
	desc.EngineVersion = o.ParseType(ins.Type)
	desc.ConnectAddr = tea.StringValue(ins.WanIp)
	desc.ConnectPort = tea.Int64Value(ins.Port)
	return r
}

var (
	typeMap = map[int64]string{
		1: "Redis2.8内存版（集群架构）",
		2: "Redis2.8内存版（标准架构）",
		3: "CKV 3.2内存版(标准架构)",
		4: "CKV 3.2内存版(集群架构)",
		5: "Redis2.8内存版（单机）",
		6: "Redis4.0内存版（标准架构）",
		7: "Redis4.0内存版（集群架构）",
		8: "Redis5.0内存版（标准架构）",
		9: "Redis5.0内存版（集群架构）",
	}
)

func (o *RedisOperator) ParseType(t *int64) string {
	if t == nil {
		return ""
	}
	return typeMap[*t]
}

var (
	billModMap = map[int64]string{
		0: "按量计费",
		1: "包年包月",
	}
)

func (o *RedisOperator) ParseBillMode(t *int64) string {
	if t == nil {
		return ""
	}

	return billModMap[*t]
}

func (o *RedisOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixMilli()
}
