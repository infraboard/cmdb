package cdb

import (
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/infraboard/cmdb/apps/rds"
)

func (o *CDBOperater) Query(req *cdb.DescribeDBInstancesRequest) (*rds.Set, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response.Items), nil
}

func (o *CDBOperater) PageQuery() rds.Pager {
	return newPager(20, o)
}
