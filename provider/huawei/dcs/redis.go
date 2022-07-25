package dcs

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
	info.PayType = o.parseChargeMod(ins.ChargingMode)
	info.Status = tea.StringValue(ins.Status)

	d := r.Describe
	d.Memory = int64(tea.Int32Value(ins.Capacity))
	d.ConnectAddr = tea.StringValue(ins.Ip)
	d.ConnectPort = int64(tea.Int32Value(ins.Port))
	d.ArchitectureType = tea.StringValue(ins.SpecCode)
	d.EngineType = tea.StringValue(ins.Engine)
	d.EngineVersion = tea.StringValue(ins.EngineVersion)
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

var (
	chargeModeMap = map[int32]string{
		0: "按需计费",
		1: "包年/包月计费",
	}
)

func (o *DcsOperator) parseChargeMod(t *int32) string {
	if t == nil {
		return ""
	}
	return chargeModeMap[*t]
}
