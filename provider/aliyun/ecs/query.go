package ecs

import (
	"context"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

// 阿里云ECS实例字段描述: https://next.api.aliyun.com/document/Ecs/2014-05-26/DescribeInstances
func (o *EcsOperator) query(req *ecs.DescribeInstancesRequest) (*host.HostSet, error) {
	set := host.NewHostSet()

	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}
	req.NextToken = req.NextToken

	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	set.Items = o.transferSet(resp.Body.Instances.Instance).Items

	return set, nil
}

func (o *EcsOperator) QueryHost(req *provider.QueryHostRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}

func (o *EcsOperator) DescribeHost(ctx context.Context, req *provider.DescribeHostRequest) (*host.Host, error) {
	r := &ecs.DescribeInstancesRequest{}
	r.InstanceIds = tea.String(`["` + req.Id + `"]`)
	hs, err := o.query(r)
	if err != nil {
		return nil, err
	}
	if hs.Length() == 0 {
		return nil, exception.NewNotFound("instance %s not found", err)
	}

	return hs.Items[0], nil
}
