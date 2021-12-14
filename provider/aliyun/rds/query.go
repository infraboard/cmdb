package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/app/rds"
)

func (o *RdsOperater) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.Set, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Items.DBInstance)
	set.Total = int64(resp.TotalRecordCount)

	return set, nil
}

func (o *RdsOperater) PageQuery() cmdbRds.Pager {
	return newPager(20, o)
}
