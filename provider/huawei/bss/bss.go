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

func (o *BssOperater) transferSet(list *[]model.NvlCostAnalysedBillDetail) *bill.BillSet {
	set := bill.NewBillSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *BssOperater) transferOne(ins model.NvlCostAnalysedBillDetail) *bill.Bill {
	fmt.Println(ins)
	b := bill.NewDefaultBill()
	b.Vendor = resource.Vendor_HUAWEI
	b.Month = utils.PtrStrV(ins.SharedMonth)
	b.OwnerId = utils.PtrStrV(ins.CustomerId)
	b.ProductType = utils.PtrStrV(ins.ServiceTypeCode)
	b.ProductCode = utils.PtrStrV(ins.ResourceTypeCode)
	b.ProductDetail = utils.PtrStrV(ins.ProductSpecDesc)
	b.PayMode = fmt.Sprintf("%d", utils.PtrInt32(ins.ChargingMode))
	b.PayModeDetail = fmt.Sprintf("%d", utils.PtrInt32(ins.BillType))
	b.OrderId = utils.PtrStrV(ins.OrderId)
	b.InstanceId = utils.PtrStrV(ins.ResourceId)
	b.InstanceName = utils.PtrStrV(ins.ResourceName)
	b.RegionCode = utils.PtrStrV(ins.RegionCode)
	b.RegionName = utils.PtrStrV(ins.RegionName)

	// // |参数名称：消费金额（应付金额）| |参数的约束及描述：|
	// ConsumeAmount *float64 `json:"consume_amount,omitempty"`

	// // |参数名称：期初已分摊金额（包周期和预留实例预付时有效，计费类型为按需，预留实例按时计费时为0）| |参数的约束及描述：|
	// PastMonthsAmortizedAmount *float64 `json:"past_months_amortized_amount,omitempty"`

	// // |参数名称：当月分摊金额| |参数的约束及描述：|
	// CurrentMonthAmortizedAmount *float64 `json:"current_month_amortized_amount,omitempty"`

	// // |参数名称：期末未分摊金额（包周期和预留实例预付时有效，计费类型为按需，预留实例按时计费时为0）| |参数的约束及描述：|
	// FutureMonthsAmortizedAmount *float64 `json:"future_months_amortized_amount,omitempty"`

	b.SalePrice = utils.PtrFloat64(ins.ConsumeAmount)
	b.RealCost = utils.PtrFloat64(ins.CurrentMonthAmortizedAmount)
	b.VoucherPay = utils.PtrFloat64(ins.AmortizedCouponAmount)
	b.CashPay = utils.PtrFloat64(ins.AmortizedCashAmount)
	b.BankcardPay = utils.PtrFloat64(ins.AmortizedStoredValueCardAmount)
	return b
}
