package billing

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/order"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
)

// 默认查询最近10年内订单
func (o *BillOperator) DescribeOrder(ctx context.Context, r *provider.DescribeRequest) (*order.Order, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := billing.NewDescribeDealsByCondRequest()
	req.StartTime = tea.String(time.Now().Add(-10 * 365 * 25 * time.Hour).Format(utils.TIME_SECOND_FORMAT_MOD1))
	req.EndTime = tea.String(time.Now().Format(utils.TIME_SECOND_FORMAT_MOD1))
	req.Limit = tea.Int64(1)
	set, err := o.QueryOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	if set.Length() == 0 {
		return nil, exception.NewNotFound("order %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *BillOperator) PageQueryOrder(req *provider.QueryOrderRequest) pager.Pager {
	p := newOrderPager(o, req)
	p.SetRate(req.Rate)
	return p
}

// 查询订单数据
// 参考: https://console.cloud.tencent.com/api/explorer?Product=billing&Version=2018-07-09&Action=DescribeDealsByCond&SignVersion=
func (o *BillOperator) QueryOrder(ctx context.Context, req *billing.DescribeDealsByCondRequest) (*order.OrderSet, error) {
	resp, err := o.client.DescribeDealsByCondWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.ToJsonString())

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
	b.PayMode = mapping.PrasePayMode(ins.PayMode)
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
