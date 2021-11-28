package billing

import (
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/infraboard/cmdb/app/bill"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。
// 您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)
// 页面顶部查看确认您的账单统计周期类型。
func newPager(pageSize int, operater *BillingOperater, rate int, month string) *pager {
	req := billing.NewDescribeBillResourceSummaryRequest()
	req.Month = common.StringPtr(month)
	req.Limit = common.Uint64Ptr(uint64(pageSize))
	req.PeriodType = common.StringPtr("byPayTime")
	rateFloat := 1 / float64(rate)
	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		log:      zap.L().Named("Pagger"),
		tb:       tokenbucket.NewBucketWithRate(rateFloat, 1),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *BillingOperater
	req      *billing.DescribeBillResourceSummaryRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Next() *bill.PagerResult {
	result := bill.NewPagerResult()

	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		result.Err = err
		return result
	}
	p.total = resp.Total
	p.log.Debugf("get %d hosts", len(resp.Items))

	result.Data = resp
	result.HasNext = p.hasNext()

	p.number++
	return result
}

func (p *pager) nextReq() *billing.DescribeBillResourceSummaryRequest {
	p.tb.Wait(1)

	p.log.Debugf("请求第%d页数据", p.number)
	p.req.Offset = common.Uint64Ptr(uint64(p.offset()))
	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}

func (p *pager) offset() int64 {
	return int64(p.size * (p.number - 1))
}
