package bss

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func newPager(pageSize int, operater *BssOperater, rate int, month string) *pager {
	req := bssopenapi.CreateQueryInstanceBillRequest()
	req.IsHideZeroCharge = requests.NewBoolean(true)
	req.PageSize = requests.NewInteger(pageSize)
	req.BillingCycle = month
	rateFloat := 1 / float64(rate)

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		total:    -1,
		log:      zap.L().Named("Pagger"),
		tb:       tokenbucket.NewBucketWithRate(rateFloat, 1),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *BssOperater
	req      *bssopenapi.QueryInstanceBillRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Scan(ctx context.Context, set *bill.BillSet) error {
	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		return err
	}

	set.Add(resp.Items...)
	p.total = int64(resp.Total)

	p.number++
	return nil
}

func (p *pager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *pager) nextReq() *bssopenapi.QueryInstanceBillRequest {
	// 等待一个可用token
	p.tb.Wait(1)

	p.log.Debugf("请求第%d页数据", p.number)
	p.req.PageNum = requests.NewInteger(p.number)
	return p.req
}

func (p *pager) Next() bool {
	if p.total == -1 {
		return true
	}
	return int64(p.number*p.size) < p.total
}
