package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/mcube/pager"
)

type RedisOperator interface {
	QueryRedis(req *QueryRedisRequest) pager.Pager
	DescribeRedis(context.Context, *DescribeRedisRequest) (*redis.Redis, error)
}

func NewDescribeRedisRequest(id string) *DescribeRedisRequest {
	return &DescribeRedisRequest{
		Id: id,
	}
}

type DescribeRedisRequest struct {
	Id string `json:"id"`
}

func NewQueryRedisWithRate(rate int32) *QueryRedisRequest {
	return &QueryRedisRequest{
		Rate: float64(rate),
	}
}

type QueryRedisRequest struct {
	Rate float64 `json:"rate"`
}
