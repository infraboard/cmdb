package cdb

import (
	"context"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/mcube/pager"
)

func (o *CDBOperator) Query(ctx context.Context, req *cdb.DescribeDBInstancesRequest) (*rds.Set, error) {
	resp, err := o.client.DescribeDBInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response.Items), nil
}

func (o *CDBOperator) PageQuery() pager.Pager {
	return newPager(20, o)
}
