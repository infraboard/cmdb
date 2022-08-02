package ecs

import (
	"context"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newEcsPager(operator *EcsOperator) pager.Pager {
	req := &ecs.DescribeInstancesRequest{
		RegionId: operator.client.RegionId,
	}

	return &ecsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.ecs"),
	}
}

type ecsPager struct {
	*pager.BasePager
	operator *EcsOperator
	req      *ecs.DescribeInstancesRequest
	log      logger.Logger
}

func (p *ecsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryInstance(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *ecsPager) nextReq() *ecs.DescribeInstancesRequest {
	p.req.PageNumber = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	p.log.Debugf("请求第%d页数据, 分页大小%d", *p.req.PageNumber, *p.req.PageSize)
	return p.req
}
