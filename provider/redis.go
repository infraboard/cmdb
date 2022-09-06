package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/mcube/pager"
)

type RedisOperator interface {
	PageQueryRedis(req *QueryRequest) pager.Pager
	DescribeRedis(context.Context, *DescribeRequest) (*redis.Redis, error)
}
