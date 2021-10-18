package cvm

import (
	"github.com/infraboard/cmdb/pkg/host"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func (o *CVMOperater) Query(req *cvm.DescribeInstancesRequest) (*host.HostSet, error) {
	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response.InstanceSet), nil
}

func (o *CVMOperater) PageQuery() host.Pager {
	return newPager(20, o)
}
