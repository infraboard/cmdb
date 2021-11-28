package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/app/host"
	"github.com/infraboard/cmdb/utils"
)

func newPager(pageSize int, operater *Ec2Operater, rate int) *pager {
	req := &ec2.DescribeInstancesInput{}
	req.MaxResults = utils.Int32Ptr(int32(pageSize))
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
	operater *Ec2Operater
	req      *ec2.DescribeInstancesInput
	log      logger.Logger
	tb       *tokenbucket.Bucket
	timeout  time.Duration
}

func (p *pager) Next() *host.PagerResult {
	result := host.NewPagerResult()

	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()

	resp, err := p.operater.Query(ctx, p.nextReq())
	if err != nil {
		result.Err = err
		return result
	}

	p.total = int64(resp.Total)

	result.Data = resp
	result.HasNext = p.hasNext()

	p.number++
	return result
}

func (p *pager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *pager) nextReq() *ec2.DescribeInstancesInput {
	// 等待一个可用token
	p.tb.Wait(1)

	p.log.Debugf("请求第%d页数据", p.number)
	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}
