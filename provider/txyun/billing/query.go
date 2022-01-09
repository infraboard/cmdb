package billing

import (
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/utils"
)

func (o *BillingOperater) Query(req *billing.DescribeBillResourceSummaryRequest) (*bill.BillSet, error) {
	resp, err := o.client.DescribeBillResourceSummary(req)
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
	Rate  int
	Month string
}

func (o *BillingOperater) PageQuery(req *PageQueryRequest) bill.Pager {
	return newPager(20, o, req.Rate, req.Month)
}
