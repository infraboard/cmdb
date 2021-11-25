package bss

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/app/bill"
)

func (o *BssOperater) Query(req *model.ListCustomerselfResourceRecordsRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()

	resp, err := o.client.ListCustomerselfResourceRecords(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferSet(resp.FeeRecords).Items
	return set, nil
}

func (o *BssOperater) PageQuery() bill.Pager {
	return newPager(20, o)
}
