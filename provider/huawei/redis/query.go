package redis

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

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
