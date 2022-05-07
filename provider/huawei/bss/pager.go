package bss

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/utils"

	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func newPager(pageSize int, operater *BssOperater, rate int, month string) *pager {
	req := &model.ListCustomerselfResourceRecordsRequest{}
	req.Cycle = month
	req.Limit = utils.Int32Ptr(int32(pageSize))
	rateFloat := 1 / float64(rate)

	return &pager{
		size:     pageSize,
		number:   1,
		total:    -1,
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
	operater *BssOperater
	req      *model.ListCustomerselfResourceRecordsRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Scan(ctx context.Context, set *bill.BillSet) error {
	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.Items...)
	p.total = resp.Total

	p.number++
	return nil
}

func (p *pager) nextReq() *model.ListCustomerselfResourceRecordsRequest {
	p.log.Debugf("请求第%d页数据", p.number)

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.number))
	return p.req
}

func (p *pager) Next() bool {
	if p.total == -1 {
		return true
	}
	return int64(p.number*p.size) < p.total
}
