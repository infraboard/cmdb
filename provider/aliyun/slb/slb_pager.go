package slb

import (
	"context"

	slb "github.com/alibabacloud-go/slb-20140515/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *SLBOperator) pager.Pager {
	return &slbPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req: &slb.DescribeLoadBalancersRequest{
			RegionId: operator.client.RegionId,
		},
		log: zap.L().Named("ali.slb"),
	}
}

type slbPager struct {
	*pager.BasePager
	operator *SLBOperator
	req      *slb.DescribeLoadBalancersRequest
	log      logger.Logger
}

func (p *slbPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryLoadBalancer(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *slbPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *slbPager) nextReq() *slb.DescribeLoadBalancersRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNumber = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	return p.req
}
