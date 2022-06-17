package rds

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *RdsOperator) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.Set, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	set := cmdbRds.NewSet()

	// 更详细的数据 需要通过DescribeDBInstanceAttribute获取，比如cpu和内存信息
	for _, ins := range resp.Items.DBInstance {
		descReq := rds.CreateDescribeDBInstanceAttributeRequest()
		descReq.DBInstanceId = ins.DBInstanceId
		detail, err := o.client.DescribeDBInstanceAttribute(descReq)
		if err != nil {
			return nil, err
		}
		set.AddSet(o.transferSet(detail.Items.DBInstanceAttribute))
	}

	set.Total = int64(resp.TotalRecordCount)

	return set, nil
}

func (o *RdsOperator) QueryRds(req *provider.QueryRdsRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

func (o *RdsOperator) DescribeRds(ctx context.Context, req *provider.DescribeRdsRequest) (*cmdbRds.Rds, error) {
	return nil, nil
}
