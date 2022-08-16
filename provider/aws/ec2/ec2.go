package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *Ec2operator) PageQueryHost(req *provider.QueryHostRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}

func (o *Ec2operator) Query(ctx context.Context, req *ec2.DescribeInstancesInput) (*host.HostSet, error) {
	set := host.NewHostSet()
	result, err := o.client.DescribeInstances(ctx, req)
	if err != nil {
		o.log.Errorf("Got an error retrieving information about your Amazon EC2 instances: %s", err)
		return nil, err
	}
	// 设置Next Page token
	req.NextToken = result.NextToken
	set.Items = o.transferSet(result.Reservations).Items
	return set, nil
}

func (o *Ec2operator) QueryDisk(req *provider.QueryDiskRequest) pager.Pager {
	panic("not imple")
}

func (o *Ec2operator) DescribeHost(ctx context.Context, req *provider.DescribeHostRequest) (*host.Host, error) {
	return nil, fmt.Errorf("not impl")
}
