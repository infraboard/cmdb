package cvm

import (
	"context"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

// 查看实例列表: https://cloud.tencent.com/document/api/213/15728
func (o *CVMOperator) Query(ctx context.Context, req *cvm.DescribeInstancesRequest) (*host.HostSet, error) {
	resp, err := o.client.DescribeInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Response.InstanceSet)
	set.Total = utils.PtrInt64(resp.Response.TotalCount)

	return set, nil
}

func NewPageQueryRequest(rate float64) *PageQueryRequest {
	return &PageQueryRequest{
		Rate: rate,
	}
}

type PageQueryRequest struct {
	Rate float64
}

func (o *CVMOperator) PageQuery(req *PageQueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}
