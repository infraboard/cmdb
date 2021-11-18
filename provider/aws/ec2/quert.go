package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/infraboard/cmdb/app/host"
)

type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

func GetInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

func (o *Ec2Operater) Query() (*host.HostSet, error) {
	set := host.NewHostSet()
	input := &ec2.DescribeInstancesInput{}
	result, err := GetInstances(context.TODO(), o.client, input)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return nil, err
	}
	set.Items = o.transferSet(result.Reservations).Items
	return set, nil
}
