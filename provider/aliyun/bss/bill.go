package bss

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
	"github.com/shopspring/decimal"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
)

func (o *BssOperator) PageQueryBill(req *provider.QueryBillRequest) pager.Pager {
	var p pager.Pager

	if req.IsSplite {
		p = newSpliteBillPager(o, req)
	} else {
		p = newResourceBillPager(o, req)
	}

	p.SetRate(req.Rate)
	return p
}

// 查询用户某个账期内所有商品实例或计费项的消费汇总
// 参考: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/DescribeInstanceBill?params={}
func (o *BssOperator) doQueryBill(req *bssopenapi.DescribeInstanceBillRequest) (*bill.BillSet, error) {
	set := bill.NewBillSet()
	resp, err := o.client.DescribeInstanceBill(req)
	if err != nil {
		return nil, err
	}

	code := tea.StringValue(resp.Body.Code)
	if strings.ToLower(code) != "success" {
		return nil, fmt.Errorf("%s: %s", tea.StringValue(resp.Body.Code), tea.StringValue(resp.Body.Message))
	}

	data := resp.Body.Data
	if data != nil {
		set.Total = int64(tea.Int32Value(data.TotalCount))
		req.NextToken = data.NextToken
		set.Items = o.transferSet(data).Items
	}

	return set, nil
}

func (o *BssOperator) transferSet(list *bssopenapi.DescribeInstanceBillResponseBodyData) *bill.BillSet {
	set := bill.NewBillSet()
	items := list.Items
	for i := range items {
		ins := o.transferBill(items[i])
		set.Add(ins)
	}
	return set
}

func (o *BssOperator) transferBill(ins *bssopenapi.DescribeInstanceBillResponseBodyDataItems) *bill.Bill {
	b := bill.NewDefaultBill()
	b.Vendor = resource.VENDOR_ALIYUN
	b.OwnerId = tea.StringValue(ins.OwnerID)
	b.OwnerName = tea.StringValue(ins.BillAccountName)
	b.ProductType = tea.StringValue(ins.ProductType)
	b.ProductCode = tea.StringValue(ins.ProductCode)
	b.ProductDetail = tea.StringValue(ins.ProductDetail)
	b.PayMode = mapping.PrasePAY_MODE(ins.Item)
	b.PayModeDetail = tea.StringValue(ins.Item) + "_" + tea.StringValue(ins.BillingType)
	b.InstanceId = tea.StringValue(ins.InstanceID)
	b.InstanceName = tea.StringValue(ins.NickName)
	b.InstanceConfig = tea.StringValue(ins.InstanceConfig)
	b.RegionName = tea.StringValue(ins.Region)
	b.Day = tea.StringValue(ins.BillingDate)
	b.Id = utils.Hash(fmt.Sprintf("%s_%s_%s_%s", b.OwnerId, b.ProductCode, b.InstanceId, b.Day))
	b.ResourceType = parseResourceType(ins.ProductCode)

	// 获取实例日账单日期
	if b.Day != "" {
		bdArray := strings.Split(b.Day, "-")
		if len(bdArray) >= 3 {
			b.Year = bdArray[0]
			b.Month = bdArray[1]
			b.Day = bdArray[2]
		}
	}

	cost := b.Cost
	cost.SalePrice = utils.Float32ToFloat64(ins.PretaxGrossAmount)
	cost.SaveCost = utils.Float32ToFloat64(ins.InvoiceDiscount)
	cost.RealCost = utils.Float32ToFloat64(ins.PretaxAmount)
	cost.StoredcardPay = utils.Float32ToFloat64(ins.DeductedByPrepaidCard)
	cost.VoucherPay = utils.Float32ToFloat64(ins.DeductedByCashCoupons)
	cost.CashPay = utils.Float32ToFloat64(ins.PaymentAmount)
	cost.OutstandingAmount = utils.Float32ToFloat64(ins.OutstandingAmount)
	return b
}

func ToFloat64(v float32) float64 {
	dv := decimal.NewFromFloat32(v)
	t, _ := dv.Float64()
	return t
}

// 查询用户某个账期内账单总览信息
// 参考: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/QueryBillOverview?params={}
func (o *BssOperator) QuerySummary(ctx context.Context, req *provider.QueryBillSummaryRequeset) (
	*bill.SummaryRecordSet, error) {
	r := &bssopenapi.QueryBillOverviewRequest{}

	ownerId, err := strconv.ParseInt(req.OwnerId, 10, 64)
	if err != nil {
		return nil, err
	}

	r.BillingCycle = tea.String(req.Month)
	r.BillOwnerId = tea.Int64(ownerId)
	resp, err := o.client.QueryBillOverview(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.String())
	return nil, nil
}

// 导出日账单到Oss
// 参考文档: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/SubscribeBillToOSS?params={}

// 查询分账账单
// 参考文档: https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/DescribeSplitItemBill
func (o *BssOperator) doDescribeSplitItemBill(req *bssopenapi.DescribeSplitItemBillRequest) (
	*bill.BillSet, error) {

	set := bill.NewBillSet()
	resp, err := o.client.DescribeSplitItemBill(req)
	if err != nil {
		return nil, err
	}

	code := tea.StringValue(resp.Body.Code)
	if strings.ToLower(code) != "success" {
		return nil, fmt.Errorf("%s: %s", tea.StringValue(resp.Body.Code), tea.StringValue(resp.Body.Message))
	}

	data := resp.Body.Data
	if data != nil {
		set.Total = int64(tea.Int32Value(data.TotalCount))
		req.NextToken = data.NextToken
		set.Items = o.transferSpliteSet(data).Items
	}

	return set, nil
}

func (o *BssOperator) transferSpliteSet(list *bssopenapi.DescribeSplitItemBillResponseBodyData) *bill.BillSet {
	set := bill.NewBillSet()
	items := list.Items
	for i := range items {
		ins := o.transferSpliteBill(items[i])
		set.Add(ins)
	}
	return set
}

func (o *BssOperator) transferSpliteBill(ins *bssopenapi.DescribeSplitItemBillResponseBodyDataItems) *bill.Bill {
	b := bill.NewDefaultBill()
	b.Vendor = resource.VENDOR_ALIYUN
	b.OwnerId = tea.StringValue(ins.OwnerID)
	b.OwnerName = tea.StringValue(ins.BillAccountName)
	b.ProductType = tea.StringValue(ins.ProductType)
	b.ProductCode = tea.StringValue(ins.ProductCode)
	b.ProductDetail = tea.StringValue(ins.ProductDetail)
	b.PayMode = mapping.PrasePAY_MODE(ins.Item)
	b.PayModeDetail = tea.StringValue(ins.Item) + "_" + tea.StringValue(ins.BillingType)
	b.InstanceId = tea.StringValue(ins.InstanceID)
	b.InstanceName = tea.StringValue(ins.NickName)
	b.InstanceConfig = tea.StringValue(ins.InstanceConfig)
	b.RegionName = tea.StringValue(ins.Region)
	b.Day = tea.StringValue(ins.BillingDate)
	b.Id = utils.Hash(fmt.Sprintf("%s_%s_%s_%s", b.OwnerId, b.ProductCode, b.InstanceId, b.Day))
	b.ResourceType = parseResourceType(ins.ProductCode)

	// 获取实例日账单日期
	if b.Day != "" {
		bdArray := strings.Split(b.Day, "-")
		if len(bdArray) >= 3 {
			b.Year = bdArray[0]
			b.Month = bdArray[1]
			b.Day = bdArray[2]
		}
	}

	cost := b.Cost
	cost.SalePrice = utils.Float32ToFloat64(ins.PretaxGrossAmount)
	cost.SaveCost = utils.Float32ToFloat64(ins.InvoiceDiscount)

	cost.RealCost = utils.Float32ToFloat64(ins.PretaxAmount)
	cost.StoredcardPay = utils.Float32ToFloat64(ins.DeductedByPrepaidCard)
	cost.VoucherPay = utils.Float32ToFloat64(ins.DeductedByCashCoupons)
	cost.CashPay = utils.Float32ToFloat64(ins.PaymentAmount)
	cost.OutstandingAmount = utils.Float32ToFloat64(ins.OutstandingAmount)

	// 处理流量包抵扣
	if ins.DeductedByResourcePackage != nil {
		d, err := decimal.NewFromString(*ins.DeductedByResourcePackage)
		if err != nil {
			zap.L().Errorf("load DeductedByResourcePackage error, %s", err)
		} else {
			f32 := tea.Float32Value(ins.PretaxAmount)
			total := d.Add(decimal.NewFromFloat32(f32))
			cost.RealCost, _ = total.Float64()
		}
	}

	return b
}
