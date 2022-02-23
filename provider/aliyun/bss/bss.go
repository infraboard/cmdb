package bss

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBssOperater(client *bssopenapi.Client) *BssOperater {
	return &BssOperater{
		client: client,
		log:    zap.L().Named("ALI BSS"),
	}
}

type BssOperater struct {
	client *bssopenapi.Client
	log    logger.Logger
}

func (o *BssOperater) transferSet(list bssopenapi.DataInQueryInstanceBill) *bill.BillSet {
	set := bill.NewBillSet()
	items := list.Items.Item
	for i := range items {
		ins := o.transferOne(items[i])
		ins.Vendor = resource.Vendor_ALIYUN
		ins.Month = list.BillingCycle
		set.Add(ins)
	}
	return set
}

func (o *BssOperater) transferOne(ins bssopenapi.Item) *bill.Bill {
	b := bill.NewDefaultBill()
	b.OwnerId = ins.OwnerID
	b.OwnerName = ins.OwnerName
	b.ProductType = ins.ProductType
	b.ProductCode = ins.ProductCode
	b.ProductDetail = ins.ProductDetail
	b.PayMode = ins.Item
	b.PayModeDetail = ins.BillingType
	b.OrderId = ins.SubOrderId
	b.InstanceId = ins.InstanceID
	b.InstanceName = ins.NickName
	b.PublicIp = ins.InternetIP
	b.PrivateIp = ins.IntranetIP
	b.InstanceConfig = ins.InstanceConfig
	b.RegionCode = ins.RegionNo
	b.RegionName = ins.Region

	cost := b.Cost
	cost.SalePrice = ins.PretaxGrossAmount
	cost.SaveCost = ins.InvoiceDiscount
	cost.RealCost = ins.PretaxAmount
	cost.StoredcardPay = ins.DeductedByPrepaidCard
	cost.VoucherPay = ins.DeductedByCashCoupons
	cost.CashPay = ins.PaymentAmount
	cost.OutstandingAmount = ins.OutstandingAmount
	return b
}
