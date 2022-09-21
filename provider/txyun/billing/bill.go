package billing

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alibabacloud-go/tea/tea"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *BillOperator) PageQueryBill(req *provider.QueryBillRequest) pager.Pager {
	p := newPager(o, req)
	p.SetRate(req.Rate)
	return p
}

// 查询账单资源汇总数据
//参考文档: https://console.cloud.tencent.com/api/explorer?Product=billing&Version=2018-07-09&Action=DescribeBillResourceSummary&SignVersion=
func (o *BillOperator) doQueryBill(ctx context.Context, req *billing.DescribeBillResourceSummaryRequest) (*bill.BillSet, error) {
	resp, err := o.client.DescribeBillResourceSummaryWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Response.ResourceSummarySet, *req.Month)
	set.Total = utils.PtrInt64(resp.Response.Total)

	return set, nil
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
	b.ProductCode = utils.PtrStrV(ins.BusinessCodeName)
	b.ProductType = utils.PtrStrV(ins.BusinessCode)
	b.ProductDetail = utils.PtrStrV(ins.ProductCodeName)
	b.PayMode = mapping.PrasePayMode(tea.StringValue(ins.PayModeName))
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

func (o *BillOperator) QuerySummary(ctx context.Context, req *provider.QueryBillSummaryRequeset) (
	*bill.SummaryRecordSet, error) {
	return nil, nil
}
