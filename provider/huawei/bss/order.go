package bss

import (
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/order"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
)

func (o *BssOperator) QueryOrder(req *provider.QueryOrderRequest) pager.Pager {
	p := newOrderPager(o)
	p.SetRate(req.Rate)
	return p
}

// 客户购买包年/包月资源后,可以查看待审核、处理中、已取消、已完成和待支付等状态的订单
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=BSS&api=ListCustomerOrders
func (o *BssOperator) doQueryOrder(req *model.ListCustomerOrdersRequest) (*order.OrderSet, error) {
	set := order.NewOrderSet()

	resp, err := o.client.ListCustomerOrders(req)
	if err != nil {
		return nil, err
	}

	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferOrderSet(resp.OrderInfos).Items
	return set, nil
}

func (o *BssOperator) transferOrderSet(list *[]model.CustomerOrderV2) *order.OrderSet {
	set := order.NewOrderSet()
	items := *list
	for i := range items {
		set.Add(o.transferOrder(items[i]))
	}
	return set
}

func (o *BssOperator) transferOrder(ins model.CustomerOrderV2) *order.Order {
	b := order.NewDefaultOrder()
	b.Vendor = resource.VENDOR_HUAWEI
	b.Id = tea.StringValue(ins.OrderId)
	b.OrderType = praseOrderType(ins.OrderType)
	b.Status = praseOrderStatus(ins.Status)
	b.ProductCode = tea.StringValue(ins.ServiceTypeCode)
	b.ProductName = tea.StringValue(ins.ServiceTypeName)
	b.CreateAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.CreateTime))
	b.PayAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.PaymentTime))

	// 金额信息
	cost := b.Cost
	cost.SalePrice = utils.PtrFloat64(ins.OfficialAmount)
	cost.RealCost = utils.PtrFloat64(ins.AmountAfterDiscount)
	return b
}

// 客户在伙伴销售平台查询某个或所有的包年/包月资源(ListPayPerUseCustomerResources)
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=BSS&api=ListPayPerUseCustomerResources
func (o *BssOperator) doOrderResource(req *model.ListPayPerUseCustomerResourcesRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()

	resp, err := o.client.ListPayPerUseCustomerResources(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	return set, nil
}

// 客户可以在伙伴销售平台查看订单详情
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=BSS&api=ShowCustomerOrderDetails
// func (o *BssOperator) doDetailOroder(req *model.ShowCustomerOrderDetailsRequest) (*bill.BillSet, error) {
// 	set := bill.NewBillSet()

// 	resp, err := o.client.ShowCustomerOrderDetails(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println(resp.String())

// 	return set, nil
// }
