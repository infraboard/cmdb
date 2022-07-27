package ecs

import (
	"context"
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

// 查询云服务器详情列表
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=ECS&api=ListServersDetails
func (o *EcsOperator) query(req *model.ListServersDetailsRequest) (*host.HostSet, error) {
	set := host.NewHostSet()

	resp, err := o.client.ListServersDetails(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.Count)
	set.Items = o.transferSet(resp.Servers).Items

	return set, nil
}

func (o *EcsOperator) QueryHost(req *provider.QueryHostRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}

func (o *EcsOperator) QueryDisk(req *provider.QueryDiskRequest) pager.Pager {
	panic("not imple")
}

func (o *EcsOperator) DescribeHost(ctx context.Context, req *provider.DescribeHostRequest) (*host.Host, error) {
	return nil, fmt.Errorf("not impl")
}
