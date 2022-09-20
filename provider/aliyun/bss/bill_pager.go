package bss

import (
	"context"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *BssOperator, r *provider.QueryBillRequest) pager.Pager {
	req := &bssopenapi.DescribeInstanceBillRequest{
		BillingCycle:  tea.String(r.Month),
		IsBillingItem: tea.Bool(false),
		Granularity:   tea.String("DAILY"),
	}

	if r.ProductCode != "" {
		req.ProductCode = tea.String(r.ProductCode)
	}

	return &bssPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.bss"),
	}
}

type bssPager struct {
	*pager.BasePager
	operator  *BssOperator
	req       *bssopenapi.DescribeInstanceBillRequest
	log       logger.Logger
	nextToken string
}

func (p *bssPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.doQueryBill(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(set)
	return nil
}

func (p *bssPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *bssPager) nextReq() *bssopenapi.DescribeInstanceBillRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.MaxResults = tea.Int32(int32(p.PageSize()))
	return p.req
}
