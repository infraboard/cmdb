package redis

import (
	"context"

	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"

	cmdbRedis "github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

// 查询Redis实例列表
// 参考: https://console.cloud.tencent.com/api/explorer?Product=redis&Version=2018-04-12&Action=DescribeInstances&SignVersion=
func (o *RedisOperator) Query(ctx context.Context, req *redis.DescribeInstancesRequest) (*cmdbRedis.Set, error) {
	resp, err := o.client.DescribeInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response), nil
}

func (o *RedisOperator) QueryRedis(req *provider.QueryRedisRequest) pager.Pager {
	return newPager(20, o)
}

func (o *RedisOperator) DescribeRedis(ctx context.Context, req *provider.DescribeRedisRequest) (*cmdbRedis.Redis, error) {
	return nil, nil
}
