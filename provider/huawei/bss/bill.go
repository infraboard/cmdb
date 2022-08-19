package bss

import (
	"context"
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *BssOperator) PageQueryBill(req *provider.QueryBillRequest) pager.Pager {
	p := newPager(o, req.Month)
	p.SetRate(req.Rate)
	return p
}

// 查询资源消费记录
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=BSS&api=ListCustomerselfResourceRecords
func (o *BssOperator) Query(req *model.ListCustomerselfResourceRecordsRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()

	resp, err := o.client.ListCustomerselfResourceRecords(req)
	if err != nil {
		return nil, err
	}
	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferSet(resp.FeeRecords).Items
	return set, nil
}

func (o *BssOperator) transferSet(list *[]model.ResFeeRecordV2) *bill.BillSet {
	set := bill.NewBillSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *BssOperator) transferOne(ins model.ResFeeRecordV2) *bill.Bill {
	b := bill.NewDefaultBill()
	b.Vendor = resource.VENDOR_HUAWEI
	b.Month = utils.PtrStrV(ins.BillDate)
	b.OwnerId = utils.PtrStrV(ins.CustomerId)
	b.ProductType = utils.PtrStrV(ins.CloudServiceType)
	b.ProductCode = utils.PtrStrV(ins.ProductId)
	b.ProductDetail = utils.PtrStrV(ins.ProductSpecDesc)
	b.PayMode = o.ParsePayMode(ins.ChargeMode)
	b.PayModeDetail = fmt.Sprintf("%d", utils.PtrInt32(ins.BillType))
	b.OrderId = utils.PtrStrV(ins.TradeId)
	b.InstanceId = utils.PtrStrV(ins.ResourceId)
	b.InstanceName = utils.PtrStrV(ins.ResourceName)
	b.RegionCode = utils.PtrStrV(ins.Region)
	b.RegionName = utils.PtrStrV(ins.RegionName)

	// 金额信息
	cost := b.Cost
	cost.SalePrice = utils.PtrFloat64(ins.OfficialAmount)
	cost.SalePrice = utils.PtrFloat64(ins.DiscountAmount)
	cost.RealCost = utils.PtrFloat64(ins.Amount)
	cost.CreditPay = utils.PtrFloat64(ins.CreditAmount)
	cost.VoucherPay = utils.PtrFloat64(ins.CouponAmount)
	cost.CashPay = utils.PtrFloat64(ins.CashAmount)
	cost.StoredcardPay = utils.PtrFloat64(ins.StoredCardAmount)
	cost.OutstandingAmount = utils.PtrFloat64(ins.DebtAmount)
	return b
}

// 	计费模式。 1：包年/包月3：按需10：预留实例
func (o *BssOperator) ParsePayMode(m *string) resource.PayMode {
	if m != nil {
		switch *m {
		case "1":
			return resource.PayMode_PRE_PAY
		case "3":
			return resource.PayMode_POST_PAY
		case "10":
			return resource.PayMode_RESERVED_PAY
		}
		return resource.PayMode_NULL
	}

	return resource.PayMode_NULL
}

func (o *BssOperator) QuerySummary(ctx context.Context, req *provider.QueryBillSummaryRequeset) (
	*bill.SummaryRecordSet, error) {
	return nil, nil
}
