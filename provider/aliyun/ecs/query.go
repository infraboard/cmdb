package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/infraboard/cmdb/pkg/host"
)

func (o *EcsOperater) Query(req *ecs.DescribeInstancesRequest) (*host.HostSet, error) {
	set := host.NewHostSet()

	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	set.Total = int64(resp.TotalCount)
	set.Items = o.transferSet(resp.Instances.Instance).Items

	return set, nil
}

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate int
}

func (o *EcsOperater) PageQuery(req *PageQueryRequest) host.Pager {
	return newPager(20, o, req.Rate)
}
