package dcs

import (
	"context"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/huawei/mapping"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *DcsOperator) DescribeRedis(ctx context.Context, r *provider.DescribeRequest) (
	*redis.Redis, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := &model.ListInstancesRequest{
		InstanceId: tea.String(r.Id),
		Limit:      tea.Int32(1),
	}

	set, err := o.Query(req)
	if err != nil {
		return nil, err
	}

	if set.Length() == 0 {
		return nil, exception.NewNotFound("redis %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *DcsOperator) PageQueryRedis(req *provider.QueryRedisRequest) pager.Pager {
	return newPager(o)
}

// 查询租户的缓存实例列表，支持按照条件查询
// 参考: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=DCS&api=ListInstances
func (o *DcsOperator) Query(req *model.ListInstancesRequest) (*redis.Set, error) {
	set := redis.NewSet()

	resp, err := o.client.ListInstances(req)
	if err != nil {
		return nil, err
	}

	// 华为云返回的TotalCount不准确
	set.Total = int64(*resp.InstanceNum)
	set.Items = o.transferSet(resp.Instances).Items

	return set, nil
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
	b.Vendor = resource.VENDOR_HUAWEI
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreatedAt))
	b.Id = tea.StringValue(ins.InstanceId)

	info := r.Information
	info.Name = tea.StringValue(ins.Name)
	info.Category = tea.StringValue(ins.SpecCode)
	info.PayMode = mapping.PrasePayMode(ins.ChargingMode)
	info.Status = praseStatus(ins.Status)

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
