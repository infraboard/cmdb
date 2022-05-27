package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/mcube/exception"
)

// 阿里云ECS实例字段描述: https://next.api.aliyun.com/document/Ecs/2014-05-26/DescribeInstances
func (o *EcsOperator) Query(req *ecs.DescribeInstancesRequest) (*host.HostSet, error) {
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

func (o *EcsOperator) PageQuery(req *PageQueryRequest) host.Pager {
	return newPager(20, o, req.Rate)
}

type DescribeRequest struct {
	Id string `json:"id"`
}

func (o *EcsOperator) Describe(req *DescribeRequest) (*host.Host, error) {
	r := ecs.CreateDescribeInstancesRequest()
	r.InstanceIds = `["` + req.Id + `"]`
	hs, err := o.Query(r)
	if err != nil {
		return nil, err
	}
	if hs.Length() == 0 {
		return nil, exception.NewNotFound("instance %s not found", err)
	}

	return hs.Items[0], nil
}
