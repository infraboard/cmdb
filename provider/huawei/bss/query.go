package bss

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/apps/bill"
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

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate  float64
	Month string
}

func (o *BssOperator) PageQuery(req *PageQueryRequest) pager.Pager {
	p := newPager(o, req.Month)
	p.SetRate(req.Rate)
	return p
}
