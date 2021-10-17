package cvm

import (
	"fmt"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func (o *CVMOperater) Query() error {
	req := cvm.NewDescribeInstancesRequest()
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return err
	}

	fmt.Println(resp.ToJsonString())
	return nil
}
