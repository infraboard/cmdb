package bss

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *BssOperator) Query(req *model.ListCustomerselfResourceRecordsRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()

	resp, err := o.client.ListCustomerselfResourceRecords(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferSet(resp.FeeRecords).Items
	return set, nil
}

func (o *BssOperator) QueryBill(req *provider.QueryBillRequest) pager.Pager {
	p := newPager(o, req.Month)
	p.SetRate(req.Rate)
	return p
}

func (o *BssOperator) QuerySummary(ctx context.Context, req *provider.QueryBillSummaryRequeset) (
	*bill.SummaryRecordSet, error) {
	return nil, nil
}
