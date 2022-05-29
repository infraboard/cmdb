package ecs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/mcube/pager"
)

func (o *EcsOperator) Query(req *model.ListServersDetailsRequest) (*host.HostSet, error) {
	set := host.NewHostSet()

	resp, err := o.client.ListServersDetails(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.Count)
	set.Items = o.transferSet(resp.Servers).Items

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

func (o *EcsOperator) PageQuery(req *PageQueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}
