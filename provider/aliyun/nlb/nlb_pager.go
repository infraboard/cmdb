package nlb

import (
	"context"

	nlb "github.com/alibabacloud-go/nlb-20220430/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *NLBOperator) pager.Pager {
	return &nlbPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       &nlb.ListLoadBalancersRequest{},
		log:       zap.L().Named("ali.alb"),
	}
}

type nlbPager struct {
	*pager.BasePager
	operator *NLBOperator
	req      *nlb.ListLoadBalancersRequest
	log      logger.Logger
}

func (p *nlbPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryLoadBalancer(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *nlbPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *nlbPager) nextReq() *nlb.ListLoadBalancersRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.MaxResults = tea.Int32(int32(p.PageSize()))
	return p.req
}
