package provider

import (
	"context"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/mcube/pager"
)

type HostOperator interface {
	PageQueryHost(req *QueryRequest) pager.Pager
	PageQueryDisk(req *QueryRequest) pager.Pager
	PageQueryEip(req *QueryRequest) pager.Pager
	DescribeHost(ctx context.Context, req *DescribeRequest) (*host.Host, error)
	DescribeDisk(ctx context.Context, req *DescribeRequest) (*disk.Disk, error)
}
