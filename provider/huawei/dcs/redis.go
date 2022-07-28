package dcs

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

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

func (o *DcsOperator) QueryRedis(req *provider.QueryRedisRequest) pager.Pager {
	return newPager(o)
}

func (o *DcsOperator) DescribeRedis(ctx context.Context, req *provider.DescribeRedisRequest) (*redis.Redis, error) {
	return nil, nil
}
