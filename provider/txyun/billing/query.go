package billing

import (
	"context"

	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *Billingoperator) Query(ctx context.Context, req *billing.DescribeBillResourceSummaryRequest) (*bill.BillSet, error) {
	resp, err := o.client.DescribeBillResourceSummaryWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Response.ResourceSummarySet, *req.Month)
	set.Total = utils.PtrInt64(resp.Response.Total)

	return set, nil
}

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate  float64
	Month string
}

func (o *Billingoperator) PageQuery(req *PageQueryRequest) pager.Pager {
	p := newPager(o, req.Month)
	p.SetRate(req.Rate)
	return p
}
