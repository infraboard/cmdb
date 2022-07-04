package billing

import (
	"context"

	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/provider"
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

func (o *Billingoperator) QueryBill(req *provider.QueryBillRequest) pager.Pager {
	p := newPager(o, req.Month)
	p.SetRate(req.Rate)
	return p
}

func (o *Billingoperator) QuerySummary(ctx context.Context, req *provider.QueryBillSummaryRequeset) (
	*bill.SummaryRecordSet, error) {
	return nil, nil
}
