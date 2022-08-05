package billing

import (
	"fmt"
	"strconv"

	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBillingoperator(client *billing.Client) *BillOperator {
	return &BillOperator{
		client: client,
		log:    zap.L().Named("Tx Billing"),
	}
}

type BillOperator struct {
	client *billing.Client
	log    logger.Logger
}

func (o *BillOperator) transferSet(items []*billing.BillResourceSummary, month string) *bill.BillSet {
	set := bill.NewBillSet()
	for i := range items {
		ins := o.transferOne(items[i])
		ins.Vendor = resource.VENDOR_TENCENT
		ins.Month = month
		set.Add(ins)
	}
	return set
}

func (o *BillOperator) transferOne(ins *billing.BillResourceSummary) *bill.Bill {
	b := bill.NewDefaultBill()
	b.OwnerId = utils.PtrStrV(ins.OwnerUin)
	b.ProductCode = utils.PtrStrV(ins.ProductCode)
	b.ProductType = utils.PtrStrV(ins.ProductCodeName)
	b.PayMode = utils.PtrStrV(ins.PayModeName)
	b.PayModeDetail = utils.PtrStrV(ins.ActionTypeName)
	b.OrderId = utils.PtrStrV(ins.OrderId)
	b.InstanceId = utils.PtrStrV(ins.ResourceId)
	b.InstanceName = utils.PtrStrV(ins.ResourceName)
	b.InstanceConfig = utils.PtrStrV(ins.ConfigDesc)
	b.RegionCode = fmt.Sprintf("%d", utils.PtrInt64(ins.RegionId))
	b.RegionName = utils.PtrStrV(ins.RegionName)

	cost := b.Cost
	cost.SalePrice, _ = strconv.ParseFloat(utils.PtrStrV(ins.TotalCost), 64)
	cost.RealCost, _ = strconv.ParseFloat(utils.PtrStrV(ins.RealTotalCost), 64)
	cost.VoucherPay, _ = strconv.ParseFloat(utils.PtrStrV(ins.VoucherPayAmount), 64)
	cost.CashPay, _ = strconv.ParseFloat(utils.PtrStrV(ins.CashPayAmount), 64)
	return b
}
