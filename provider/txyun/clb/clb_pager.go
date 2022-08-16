package clb

import (
	"context"

	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *CLBOperator) pager.Pager {
	return &clbPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       clb.NewDescribeLoadBalancersRequest(),
		log:       zap.L().Named("tx.clb"),
	}
}

type clbPager struct {
	*pager.BasePager
	operator *CLBOperator
	req      *clb.DescribeLoadBalancersRequest
	log      logger.Logger
}

func (p *clbPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryLoadBalancer(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d clb", len(resp.Items))

	p.CheckHasNext(resp)
	return nil
}

func (p *clbPager) nextReq() *clb.DescribeLoadBalancersRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Int64Ptr(int64(p.Offset()))
	p.req.Limit = common.Int64Ptr(int64(p.PageSize()))
	return p.req
}
