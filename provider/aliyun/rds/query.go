package rds

import (
	"context"

	rds "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

// 查询RDS实例列表
// 参考文档: https://next.api.aliyun.com/api/Rds/2014-08-15/DescribeDBInstances?params={}&lang=GO
func (o *RdsOperator) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.Set, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}
	req.NextToken = resp.Body.NextToken

	set := cmdbRds.NewSet()
	// 更详细的数据 需要通过DescribeDBInstanceAttribute获取，比如cpu和内存信息
	for _, ins := range resp.Body.Items.DBInstance {
		descReq := &rds.DescribeDBInstanceAttributeRequest{}
		descReq.DBInstanceId = ins.DBInstanceId
		detail, err := o.client.DescribeDBInstanceAttribute(descReq)
		if err != nil {
			return nil, err
		}
		set.AddSet(o.transferSet(detail.Body.Items))
	}

	set.Total = int64(tea.Int32Value(resp.Body.TotalRecordCount))
	return set, nil
}

func (o *RdsOperator) QueryRds(req *provider.QueryRdsRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

func (o *RdsOperator) DescribeRds(ctx context.Context, req *provider.DescribeRdsRequest) (*cmdbRds.Rds, error) {
	descReq := &rds.DescribeDBInstanceAttributeRequest{
		DBInstanceId: &req.Id,
	}

	detail, err := o.client.DescribeDBInstanceAttribute(descReq)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(detail.Body.Items)
	if set.Length() == 0 {
		return nil, exception.NewNotFound("ins %s not found", req.Id)
	}

	return set.Items[0], nil
}
