package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/mcube/pager"
)

type RdsOperator interface {
	QueryRds(req *QueryRdsRequest) pager.Pager
	DescribeRds(ctx context.Context, req *DescribeRdsRequest) (*rds.Rds, error)
}

func NewQueryRdsRequest() *QueryRdsRequest {
	return &QueryRdsRequest{
		Rate: 5,
	}
}

func NewQueryRdsRequestWithRate(rate int32) *QueryRdsRequest {
	return &QueryRdsRequest{
		Rate: float64(rate),
	}
}

type QueryRdsRequest struct {
	Rate float64 `json:"rate"`
}

type DescribeRdsRequest struct {
	Id string `json:"id"`
}
