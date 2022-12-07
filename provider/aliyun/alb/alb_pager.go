package alb

import (
	"context"

	alb "github.com/alibabacloud-go/alb-20200616/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *ALBOperator) pager.Pager {
	return &albPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       &alb.ListLoadBalancersRequest{},
		log:       zap.L().Named("ali.alb"),
	}
}

type albPager struct {
	*pager.BasePager
	operator *ALBOperator
	req      *alb.ListLoadBalancersRequest
	log      logger.Logger
}

func (p *albPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryLoadBalancer(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *albPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *albPager) nextReq() *alb.ListLoadBalancersRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.MaxResults = tea.Int32(int32(p.PageSize()))
	return p.req
}
