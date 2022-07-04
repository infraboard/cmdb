package bss

import (
	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBssOperator(client *bssopenapi.Client) *BssOperator {
	return &BssOperator{
		client: client,
		log:    zap.L().Named("ALI BSS"),
	}
}

type BssOperator struct {
	client *bssopenapi.Client
	log    logger.Logger
}

func (o *BssOperator) transferSet(list *bssopenapi.DescribeInstanceBillResponseBodyData) *bill.BillSet {
	set := bill.NewBillSet()
	items := list.Items
	for i := range items {
		ins := o.transferOne(items[i])
		ins.Vendor = resource.Vendor_ALIYUN
		ins.Month = tea.StringValue(list.BillingCycle)
		set.Add(ins)
	}
	return set
}

func (o *BssOperator) transferOne(ins *bssopenapi.DescribeInstanceBillResponseBodyDataItems) *bill.Bill {
	b := bill.NewDefaultBill()
	b.OwnerId = tea.StringValue(ins.OwnerID)
	b.OwnerName = tea.StringValue(ins.BillAccountName)
	b.ProductType = tea.StringValue(ins.ProductType)
	b.ProductCode = tea.StringValue(ins.ProductCode)
	b.ProductDetail = tea.StringValue(ins.ProductDetail)
	b.PayMode = tea.StringValue(ins.Item)
	b.PayModeDetail = tea.StringValue(ins.BillingType)
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
