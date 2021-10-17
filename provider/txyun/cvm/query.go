package cvm

import (
	"github.com/infraboard/cmdb/pkg/host"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func (o *CVMOperater) Query() (*host.HostSet, error) {
	req := cvm.NewDescribeInstancesRequest()
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response.InstanceSet), nil
}
