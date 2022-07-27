package provider

import "github.com/infraboard/mcube/pager"

type LBOperator interface {
	QueryLB(req *QueryLBRequest) pager.Pager
}

func NewQueryLBRequest() *QueryLBRequest {
	return &QueryLBRequest{
		Rate: 5,
	}
}

type QueryLBRequest struct {
	Rate float64 `json:"rate"`
}
