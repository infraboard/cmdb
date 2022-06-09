package provider

import (
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/mcube/pager"
)

type HostOperator interface {
	PageQuery(req *PageQueryRequest) pager.Pager
	Describe(req *DescribeRequest) (*host.Host, error)
}

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate float64
}

type DescribeRequest struct {
	Id string `json:"id"`
}
