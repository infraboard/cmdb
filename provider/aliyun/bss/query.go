package bss

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *BssOperator) Query(req *bssopenapi.QueryInstanceBillRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()
	resp, err := o.client.QueryInstanceBill(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	set.Total = int64(resp.Data.TotalCount)
	set.Items = o.transferSet(resp.Data).Items
	return set, nil
}

func (o *BssOperator) QueryBill(req *provider.QueryBillRequest) pager.Pager {
	p := newPager(o, req)
	p.SetRate(req.Rate)
	return p
}
