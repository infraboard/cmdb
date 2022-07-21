package redis

import (
	"time"

	"github.com/alibabacloud-go/tea/tea"
	dcs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewDcsOperator(client *dcs.DcsClient) *DcsOperator {
	return &DcsOperator{
		client: client,
		log:    zap.L().Named("huawei.redis"),
	}
}

type DcsOperator struct {
	client *dcs.DcsClient
	log    logger.Logger
}

func (o *DcsOperator) transferSet(list *[]model.InstanceListInfo) *redis.Set {
	set := redis.NewSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *DcsOperator) transferOne(ins model.InstanceListInfo) *redis.Redis {
	r := redis.NewDefaultRedis()
	b := r.Base
	b.Vendor = resource.Vendor_HUAWEI
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreatedAt))
	b.Id = tea.StringValue(ins.InstanceId)

	info := r.Information
	info.Name = tea.StringValue(ins.Name)
	info.Category = tea.StringValue(ins.SpecCode)
	return r
}

func (o *DcsOperator) parseTime(t string) int64 {
	if t == "" {
		return 0
	}

	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixMilli()
}
