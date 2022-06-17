package provider

import (
	"github.com/infraboard/mcube/pager"
)

type BillOperator interface {
	QueryBill(req *QueryBillRequest) pager.Pager
}

func NewQueryBillRequest() *QueryBillRequest {
	return &QueryBillRequest{
		Rate: 5,
	}
}

func NewQueryBillRequestWithRate(rate int32) *QueryBillRequest {
	return &QueryBillRequest{
		Rate: float64(rate),
	}
}

type QueryBillRequest struct {
	Rate        float64
	Month       string
	ProductCode string
}
