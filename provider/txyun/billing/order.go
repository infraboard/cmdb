package billing

import (
	"context"
	"encoding/json"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/order"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
)

func (o *BillOperator) QueryOrder(req *provider.QueryOrderRequest) pager.Pager {
	p := newOrderPager(o, req)
	p.SetRate(req.Rate)
	return p
}

// 查询订单数据
// 参考: https://console.cloud.tencent.com/api/explorer?Product=billing&Version=2018-07-09&Action=DescribeDealsByCond&SignVersion=
func (o *BillOperator) doQueryOrder(ctx context.Context, req *billing.DescribeDealsByCondRequest) (*order.OrderSet, error) {
	resp, err := o.client.DescribeDealsByCondWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	set := o.transferOrderSet(resp.Response.Deals)
	set.Total = utils.PtrInt64(resp.Response.TotalCount)
	return set, nil
}

func (o *BillOperator) transferOrderSet(items []*billing.Deal) *order.OrderSet {
	set := order.NewOrderSet()
	for i := range items {
		ins := o.transferOrder(items[i])
		ins.Vendor = resource.VENDOR_TENCENT
		set.Add(ins)
	}
	return set
}

func (o *BillOperator) transferOrder(ins *billing.Deal) *order.Order {
	b := order.NewDefaultOrder()
	b.BigOrderId = tea.StringValue(ins.BigDealId)
	b.Id = tea.StringValue(ins.OrderId)
	b.OrderType = praseOrderType(ins.Action)
	b.Status = praseOrderStatus(ins.Status)
	b.Payer = tea.StringValue(ins.Payer)
	b.CreateAt = utils.ParseSecondMod1Time(tea.StringValue(ins.CreateTime))
	b.CreateBy = tea.StringValue(ins.Creator)
	b.PayMode = tea.StringValue(ins.PayMode)
	b.ProductCode = tea.StringValue(ins.ProductCode)
	b.ProductName = tea.StringValue(ins.ProductName)
	b.SubProductCode = tea.StringValue(ins.SubProductCode)
	b.SubProductName = tea.StringValue(ins.SubProductName)
	pi, _ := json.Marshal(ins.ProductInfo)
	b.ProductInfo = string(pi)
	b.ResourceId = tea.StringSliceValue(ins.ResourceId)

	cost := b.Cost
	cost.SalePrice = tea.Float64Value(ins.TotalCost)
	cost.Policy = tea.Float64Value(ins.Policy)
	cost.TimeSpan = tea.Float64Value(ins.TimeSpan)
	cost.TimeUnit = tea.StringValue(ins.TimeUnit)
	cost.RealCost = float64(tea.Int64Value(ins.RealTotalCost))
	cost.VoucherPay = float64(tea.Int64Value(ins.VoucherDecline))

	return b
}
