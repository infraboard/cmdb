package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/mcube/pager"
)

type RedisOperator interface {
	PageQueryRedis(req *QueryRedisRequest) pager.Pager
	DescribeRedis(context.Context, *DescribeRequest) (*redis.Redis, error)
}

func NewQueryRedisWithRate(rate int32) *QueryRedisRequest {
	return &QueryRedisRequest{
		Rate: float64(rate),
	}
}

type QueryRedisRequest struct {
	Rate float64 `json:"rate"`
}
