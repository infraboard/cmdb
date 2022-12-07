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

func newResourceBillPager(operator *BssOperator, r *provider.QueryBillRequest) pager.Pager {
	req := &bssopenapi.DescribeInstanceBillRequest{
		BillingCycle: tea.String(r.Month),
	}

	if r.ProductCode != "" {
		req.ProductCode = tea.String(r.ProductCode)
	}

	return &resourceBillPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.resource.bill"),
	}
}

type resourceBillPager struct {
	*pager.BasePager
	operator *BssOperator
	req      *bssopenapi.DescribeInstanceBillRequest
	log      logger.Logger
}

func (p *resourceBillPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.doQueryBill(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *resourceBillPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *resourceBillPager) nextReq() *bssopenapi.DescribeInstanceBillRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.MaxResults = tea.Int32(int32(p.PageSize()))
	return p.req
}

func newSpliteBillPager(operator *BssOperator, r *provider.QueryBillRequest) pager.Pager {
	req := &bssopenapi.DescribeSplitItemBillRequest{
		BillingCycle: tea.String(r.Month),
	}

	if r.ProductCode != "" {
		req.ProductCode = tea.String(r.ProductCode)
	}

	return &spliteBillPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.resource.bill_splite"),
	}
}

type spliteBillPager struct {
	*pager.BasePager
	operator *BssOperator
	req      *bssopenapi.DescribeSplitItemBillRequest
	log      logger.Logger
}

func (p *spliteBillPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.doDescribeSplitItemBill(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *spliteBillPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *spliteBillPager) nextReq() *bssopenapi.DescribeSplitItemBillRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.MaxResults = tea.Int32(int32(p.PageSize()))
	return p.req
}
