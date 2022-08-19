package provider

import (
	"context"
	"time"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/order"
	"github.com/infraboard/mcube/pager"
)

type BillOperator interface {
	PageQueryBill(*QueryBillRequest) pager.Pager
	QuerySummary(context.Context, *QueryBillSummaryRequeset) (*bill.SummaryRecordSet, error)
	PageQueryOrder(*QueryOrderRequest) pager.Pager
	DescribeOrder(context.Context, *DescribeRequest) (*order.Order, error)
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

func NewQueryBillSummaryRequeset() *QueryBillSummaryRequeset {
	return &QueryBillSummaryRequeset{}
}

type QueryBillSummaryRequeset struct {
	// 子账号ID, 设置可查看财务云子账号账单，不填默认查看当前调用账号
	OwnerId string
	// 账单月份
	Month string
}

func NewQueryOrderRequest() *QueryOrderRequest {
	now := time.Now()
	return &QueryOrderRequest{
		Rate:      5,
		StartTime: now.Add(-1 * time.Hour),
		EndTime:   now,
	}
}

type QueryOrderRequest struct {
	Rate      float64
	StartTime time.Time
	EndTime   time.Time
}
