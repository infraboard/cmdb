package rds

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"

	"github.com/infraboard/cmdb/apps/rds"
	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

// 查询数据库实例列表
// 参考: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=RDS&api=ListInstances
func (o *RdsOperator) Query(req *model.ListInstancesRequest) (*rds.Set, error) {
	set := rds.NewSet()

	resp, err := o.client.ListInstances(req)
	if err != nil {
		return nil, err
	}

	// 华为云返回的TotalCount不准确
	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferSet(resp.Instances).Items

	return set, nil
}

func (o *RdsOperator) QueryRds(req *provider.QueryRdsRequest) pager.Pager {
	return newPager(o)
}

func (o *RdsOperator) DescribeRds(ctx context.Context, req *provider.DescribeRdsRequest) (*cmdbRds.Rds, error) {
	return nil, nil
}
