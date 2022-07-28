package bss

import (
	"context"
	"fmt"
	"strconv"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/provider"
)

// 查询用户某个账期内所有商品实例或计费项的消费汇总
// 参考: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/DescribeInstanceBill?params={}
func (o *BssOperator) Query(req *bssopenapi.DescribeInstanceBillRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()
	resp, err := o.client.DescribeInstanceBill(req)
	if err != nil {
		return nil, err
	}
	data := resp.Body.Data
	set.Total = int64(*data.TotalCount)
	req.NextToken = data.NextToken
	set.Items = o.transferSet(data).Items
	return set, nil
}

func (o *BssOperator) QueryBill(req *provider.QueryBillRequest) pager.Pager {
	p := newPager(o, req)
	p.SetRate(req.Rate)
	return p
}

// 查询用户某个账期内账单总览信息
// 参考: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/QueryBillOverview?params={}
func (o *BssOperator) QuerySummary(ctx context.Context, req *provider.QueryBillSummaryRequeset) (
	*bill.SummaryRecordSet, error) {
	r := &bssopenapi.QueryBillOverviewRequest{}

	ownerId, err := strconv.ParseInt(req.OwnerId, 10, 64)
	if err != nil {
		return nil, err
	}

	r.BillingCycle = tea.String(req.Month)
	r.BillOwnerId = tea.Int64(ownerId)
	resp, err := o.client.QueryBillOverview(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())
	return nil, nil
}
