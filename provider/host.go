package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/mcube/pager"
)

type HostOperator interface {
	PageQueryHost(req *QueryHostRequest) pager.Pager
	PageQueryDisk(req *QueryDiskRequest) pager.Pager
	PageQueryEip(req *QueryEipRequest) pager.Pager
	DescribeHost(ctx context.Context, req *DescribeHostRequest) (*host.Host, error)
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		Rate: 5,
	}
}

func NewQueryHostRequestWithRate(rate int32) *QueryHostRequest {
	return &QueryHostRequest{
		Rate: float64(rate),
	}
}

type QueryHostRequest struct {
	Rate float64 `json:"rate"`
}

func NewDescribeHostRequest(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

type DescribeHostRequest struct {
	Id string `json:"id"`
}

func NewQueryDiskRequest() *QueryDiskRequest {
	return &QueryDiskRequest{
		Rate: 5,
	}
}

type QueryDiskRequest struct {
	Rate float64 `json:"rate"`
}

func NewQueryEipRequest() *QueryEipRequest {
	return &QueryEipRequest{
		Rate: 5,
	}
}

type QueryEipRequest struct {
	Rate float64 `json:"rate"`
}
