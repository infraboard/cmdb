package ec2

import (
	"context"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *Ec2operator) DescribeDisk(ctx context.Context, req *provider.DescribeRequest) (
	*disk.Disk, error) {
	return nil, nil
}

func (o *Ec2operator) PageQueryDisk(req *provider.QueryRequest) pager.Pager {
	return nil
}
