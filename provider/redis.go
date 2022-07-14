package provider

import "github.com/infraboard/mcube/pager"

type RedisOperator interface {
	QueryRedis(req *QueryRedisRequest) pager.Pager
}

func NewQueryRedisWithRate(rate int32) *QueryRedisRequest {
	return &QueryRedisRequest{
		Rate: float64(rate),
	}
}

type QueryRedisRequest struct {
	Rate float64 `json:"rate"`
}
