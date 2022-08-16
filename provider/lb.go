package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/lb"
	"github.com/infraboard/mcube/pager"
)

type LBOperator interface {
	DescribeLoadBalancer(context.Context, *DescribeRequest) (*lb.LoadBalancer, error)
	PageQueryLoadBalancer(*QueryRequest) pager.Pager
}
