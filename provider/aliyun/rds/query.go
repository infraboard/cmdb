package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/pkg/rds"
)

func (o *RdsOperater) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.RdsSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Items.DBInstance), nil
}
