package bss

import (
	"context"
	"fmt"
	"strconv"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
)

func (o *BssOperator) PageQueryBill(req *provider.QueryBillRequest) pager.Pager {
	p := newPager(o, req)
	p.SetRate(req.Rate)
	return p
}

// 查询用户某个账期内所有商品实例或计费项的消费汇总
// 参考: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/DescribeInstanceBill?params={}
func (o *BssOperator) doQueryBill(req *bssopenapi.DescribeInstanceBillRequest) (*bill.BillSet, error) {
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

func (o *BssOperator) transferSet(list *bssopenapi.DescribeInstanceBillResponseBodyData) *bill.BillSet {
	set := bill.NewBillSet()
	items := list.Items
	for i := range items {
		ins := o.transferBill(items[i])
		ins.Vendor = resource.VENDOR_ALIYUN
		ins.Month = tea.StringValue(list.BillingCycle)
		set.Add(ins)
	}
	return set
}

func (o *BssOperator) transferBill(ins *bssopenapi.DescribeInstanceBillResponseBodyDataItems) *bill.Bill {
	b := bill.NewDefaultBill()
	b.OwnerId = tea.StringValue(ins.OwnerID)
	b.OwnerName = tea.StringValue(ins.BillAccountName)
	b.ProductType = tea.StringValue(ins.ProductType)
	b.ProductCode = tea.StringValue(ins.ProductCode)
	b.ProductDetail = tea.StringValue(ins.ProductDetail)
	b.PayMode = mapping.PrasePayMode(ins.Item)
	b.PayModeDetail = tea.StringValue(ins.Item) + "_" + tea.StringValue(ins.BillingType)
	b.InstanceId = tea.StringValue(ins.InstanceID)
	b.InstanceName = tea.StringValue(ins.NickName)
	b.PublicIp = tea.StringValue(ins.InternetIP)
	b.PrivateIp = tea.StringValue(ins.IntranetIP)
	b.InstanceConfig = tea.StringValue(ins.InstanceConfig)
	b.RegionName = tea.StringValue(ins.Region)

	cost := b.Cost
	cost.SalePrice = float64(tea.Float32Value(ins.PretaxGrossAmount))
	cost.SaveCost = float64(tea.Float32Value(ins.InvoiceDiscount))
	cost.RealCost = float64(tea.Float32Value(ins.PretaxAmount))
	cost.StoredcardPay = float64(tea.Float32Value(ins.DeductedByPrepaidCard))
	cost.VoucherPay = float64(tea.Float32Value(ins.DeductedByCashCoupons))
	cost.CashPay = float64(tea.Float32Value(ins.PaymentAmount))
	cost.OutstandingAmount = float64(tea.Float32Value(ins.OutstandingAmount))
	return b
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
