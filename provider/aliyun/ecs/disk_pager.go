package ecs

import (
	"context"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newDiskPager(operator *EcsOperator) pager.Pager {
	req := &ecs.DescribeDisksRequest{
		RegionId: operator.client.RegionId,
	}

	return &diskPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.disk"),
	}
}

type diskPager struct {
	*pager.BasePager
	operator *EcsOperator
	req      *ecs.DescribeDisksRequest
	log      logger.Logger
}

func (p *diskPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryDisk(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *diskPager) nextReq() *ecs.DescribeDisksRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNumber = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	return p.req
}
