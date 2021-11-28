package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infraboard/cmdb/app/host"
)

func (o *Ec2Operater) Query(ctx context.Context, req *ec2.DescribeInstancesInput) (*host.HostSet, error) {
	set := host.NewHostSet()
	result, err := o.client.DescribeInstances(ctx, req)
	if err != nil {
		o.log.Errorf("Got an error retrieving information about your Amazon EC2 instances: %s", err)
		return nil, err
	}
	// 设置Next Page token
	req.NextToken = req.NextToken
	set.Items = o.transferSet(result.Reservations).Items
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

func (o *Ec2Operater) PageQuery(req *PageQueryRequest) host.Pager {
	return newPager(20, o, req.Rate)
}
