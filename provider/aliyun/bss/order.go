package bss

import (
	"encoding/json"
	"strconv"
	"sync"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/order"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *BssOperator) QueryOrder(req *provider.QueryOrderRequest) pager.Pager {
	p := newOrderPager(o, req)
	p.SetRate(req.Rate)
	return p
}

// 查询用户或者分销客户订单列表情况。默认查询当前时间最近1小时范围内订单，如需查询更长时间范围的订单数据，
// 请设**CreateTimeStart** 和**CreateTimeEnd**参数
// 参考文档: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/QueryOrders?params={}
func (o *BssOperator) doQueryOrder(req *bssopenapi.QueryOrdersRequest) (*order.OrderSet, error) {
	set := order.NewOrderSet()
	resp, err := o.client.QueryOrders(req)
	if err != nil {
		return nil, err
	}

	if resp.Body.Data.OrderList == nil {
		return set, nil
	}

	set.Total = int64(*resp.Body.Data.TotalCount)
	// 获取Order Id列表
	orderIds := []string{}
	orders := resp.Body.Data.OrderList.Order
	for i := range orders {
		orderIds = append(orderIds, tea.StringValue(orders[i].OrderId))
	}

	// 查询订单详情
	wg := &sync.WaitGroup{}
	for _, oid := range orderIds {
		wg.Add(1)
		go func(oid string) {
			defer wg.Done()
			ins, err := o.doDescribeOrder(&bssopenapi.GetOrderDetailRequest{
				OrderId: tea.String(oid),
			})
			if err != nil {
				o.log.Errorf("describe order %s error, %s", oid, err)
				return
			}
			set.Add(ins.ToAny()...)
		}(oid)
	}
	wg.Wait()
	return set, nil
}

// 查询用户或分销客户某个订单详情信息
// 参考文档: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/GetOrderDetail?params={}
func (o *BssOperator) doDescribeOrder(req *bssopenapi.GetOrderDetailRequest) (*order.OrderSet, error) {
	o.tb.Wait(1)
	o.log.Debugf("query order: %s detail", tea.StringValue(req.OrderId))

	set := order.NewOrderSet()
	resp, err := o.client.GetOrderDetail(req)
	if err != nil {
		return nil, err
	}
	data := resp.Body.Data
	set.Items = o.transferOrderSet(data).Items
	return set, nil
}

func (o *BssOperator) transferOrderSet(list *bssopenapi.GetOrderDetailResponseBodyData) *order.OrderSet {
	set := order.NewOrderSet()
	items := list.OrderList.Order
	for i := range items {
		ins := o.transferOrder(items[i])
		ins.Vendor = resource.VENDOR_ALIYUN
		set.Add(ins)
	}
	return set
}

func (o *BssOperator) transferOrder(ins *bssopenapi.GetOrderDetailResponseBodyDataOrderListOrder) *order.Order {
	b := order.NewDefaultOrder()
	b.BigOrderId = tea.StringValue(ins.OrderId)
	b.Id = tea.StringValue(ins.SubOrderId)
	b.OrderType = praseOrderType(ins.OrderType)
	b.Status = praseOrderStatus(ins.PaymentStatus)
	b.CreateAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.CreateTime))
	b.PayAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.PaymentTime))
	b.PayMode = mapping.PrasePayMode(ins.SubscriptionType)
	b.ProductCode = tea.StringValue(ins.ProductCode)
	b.ProductInfo = tea.StringValue(ins.OriginalConfig)

	if err := json.Unmarshal([]byte(tea.StringValue(ins.InstanceIDs)), &b.ResourceId); err != nil {
		o.log.Warnf("unmarshal resource id error, %s", err)
	}

	cost := b.Cost
	cost.SalePrice, _ = strconv.ParseFloat(tea.StringValue(ins.PretaxGrossAmount), 64)
	cost.RealCost, _ = strconv.ParseFloat(tea.StringValue(ins.PretaxAmount), 64)
	return b
}
