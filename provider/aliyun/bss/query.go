package bss

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/infraboard/cmdb/apps/bill"
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

func NewPageQueryRequest() *PageQueryRequest {
	return &PageQueryRequest{
		Rate: 1,
	}
}

type PageQueryRequest struct {
	Rate  int
	Month string
}

func (o *BssOperator) PageQuery(req *PageQueryRequest) bill.Pager {
	return newPager(20, o, req.Rate, req.Month)
}
