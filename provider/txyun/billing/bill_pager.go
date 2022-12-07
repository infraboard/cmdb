package billing

import (
	"context"

	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/pager"
)

// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。
// 您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)
// 页面顶部查看确认您的账单统计周期类型。
func newPager(operator *BillOperator, r *provider.QueryBillRequest) pager.Pager {
	req := billing.NewDescribeBillResourceSummaryRequest()
	req.Month = common.StringPtr(r.Month)
	req.PeriodType = common.StringPtr("byPayTime")

	return &billPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       operator.log,
	}
}

type billPager struct {
	*pager.BasePager
	operator *BillOperator
	req      *billing.DescribeBillResourceSummaryRequest
	log      logger.Logger
}

func (p *billPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.doQueryBill(ctx, p.nextReq())
	if err != nil {
		return err
	}
	// 由于账单接口并没有返回Total总数量, 无法通过梳理判断是否数据已经拉起完成, 改而判断是否有数据
	p.CheckHasNext(resp)
	p.log.Debugf("get %d hosts", resp.Length())

	set.Add(resp.ToAny()...)
	return nil
}

func (p *billPager) nextReq() *billing.DescribeBillResourceSummaryRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Uint64Ptr(uint64(p.Offset()))
	p.req.Limit = common.Uint64Ptr(uint64(p.PageSize()))
	return p.req
}
