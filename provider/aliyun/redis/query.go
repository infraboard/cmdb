package redis

import (
	"context"

	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	cmdbRedis "github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/provider"
)

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

func (o *RedisOperator) QueryRedis(req *provider.QueryRedisRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

func (o *RedisOperator) DescribeRedis(ctx context.Context, req *provider.DescribeRedisRequest) (*cmdbRedis.Redis, error) {
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
