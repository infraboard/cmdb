package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/infraboard/cmdb/pkg/host"
)

func (o *EcsOperater) Query() (*host.HostSet, error) {
	req := ecs.CreateDescribeInstancesRequest()
	req.PageSize = requests.NewInteger(50)
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Instances.Instance), nil
}
