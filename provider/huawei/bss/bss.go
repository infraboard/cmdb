package bss

import (
	"fmt"

	v2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"

	"github.com/infraboard/cmdb/app/bill"
	"github.com/infraboard/cmdb/app/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBssOperater(client *v2.BssClient) *BssOperater {
	return &BssOperater{
		client: client,
		log:    zap.L().Named("Huawei BSS"),
	}
}

type BssOperater struct {
	client *v2.BssClient
	log    logger.Logger
}

func (o *BssOperater) transferSet(list *[]model.ResFeeRecordV2) *bill.BillSet {
	set := bill.NewBillSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *BssOperater) transferOne(ins model.ResFeeRecordV2) *bill.Bill {
	b := bill.NewDefaultBill()
	b.Vendor = resource.Vendor_HUAWEI
	b.Month = utils.PtrStrV(ins.BillDate)
	b.OwnerId = utils.PtrStrV(ins.CustomerId)
	b.ProductType = utils.PtrStrV(ins.CloudServiceType)
	b.ProductCode = utils.PtrStrV(ins.ProductId)
	b.ProductDetail = utils.PtrStrV(ins.ProductSpecDesc)
	b.PayMode = utils.PtrStrV(ins.ChargeMode)
	b.PayModeDetail = fmt.Sprintf("%d", utils.PtrInt32(ins.BillType))
	b.OrderId = utils.PtrStrV(ins.TradeId)
	b.InstanceId = utils.PtrStrV(ins.ResourceId)
	b.InstanceName = utils.PtrStrV(ins.ResourceName)
	b.RegionCode = utils.PtrStrV(ins.Region)
	b.RegionName = utils.PtrStrV(ins.RegionName)

	// 金额信息
	b.SalePrice = utils.PtrFloat64(ins.OfficialAmount)
	b.SalePrice = utils.PtrFloat64(ins.DiscountAmount)
	b.RealCost = utils.PtrFloat64(ins.Amount)
	b.CreditPay = utils.PtrFloat64(ins.CreditAmount)
	b.VoucherPay = utils.PtrFloat64(ins.CouponAmount)
	b.CashPay = utils.PtrFloat64(ins.CashAmount)
	b.StoredcardPay = utils.PtrFloat64(ins.StoredCardAmount)
	b.OutstandingAmount = utils.PtrFloat64(ins.DebtAmount)
	return b
}
