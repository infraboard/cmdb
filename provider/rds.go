package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/mcube/pager"
)

type RdsOperator interface {
	PageQueryRds(req *QueryRequest) pager.Pager
	DescribeRds(ctx context.Context, req *DescribeRequest) (*rds.Rds, error)
}
