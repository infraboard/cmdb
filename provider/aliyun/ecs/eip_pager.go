package ecs

import (
	"context"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newEipPager(operator *EcsOperator) pager.Pager {
	req := &ecs.DescribeEipAddressesRequest{
		RegionId: operator.client.RegionId,
	}

	return &eipPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.eip"),
	}
}

type eipPager struct {
	*pager.BasePager
	operator *EcsOperator
	req      *ecs.DescribeEipAddressesRequest
	log      logger.Logger
}

func (p *eipPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryEip(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *eipPager) nextReq() *ecs.DescribeEipAddressesRequest {
	p.req.PageNumber = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	p.log.Debugf("请求第%d页数据, 分页大小%d", *p.req.PageNumber, *p.req.PageSize)
	return p.req
}
