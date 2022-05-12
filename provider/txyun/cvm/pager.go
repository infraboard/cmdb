package cvm

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func newPager(pageSize int, operater *CVMOperator, reqPs int) *pager {
	req := cvm.NewDescribeInstancesRequest()
	req.Limit = common.Int64Ptr(int64(pageSize))

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		total:    -1,
		log:      zap.L().Named("tx.cvm"),
		tb:       tokenbucket.NewBucketWithRate(1/float64(reqPs), 1),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *CVMOperator
	req      *cvm.DescribeInstancesRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Scan(ctx context.Context, set *host.HostSet) error {
	resp, err := p.operater.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.Items...)
	p.total = resp.Total
	p.log.Debugf("get %d hosts", len(resp.Items))

	p.number++
	return nil
}

func (p *pager) nextReq() *cvm.DescribeInstancesRequest {
	// 生成请求的时候, 现获取速率令牌, 等待一个可用的令牌
	p.tb.Wait(1)

	p.log.Debugf("请求第%d页数据", p.number)
	p.req.Offset = common.Int64Ptr(p.offset())
	return p.req
}

func (p *pager) Next() bool {
	if p.total == -1 {
		return true
	}
	return int64(p.number*p.size) < p.total
}

func (p *pager) offset() int64 {
	return int64(p.size * (p.number - 1))
}
