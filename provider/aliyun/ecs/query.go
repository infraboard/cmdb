package ecs

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func (o *EcsOperater) Query() error {
	req := ecs.CreateDescribeInstancesRequest()
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return err
	}

	fmt.Println(resp.String())
	return nil
}
