package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/utils"
)

func newPager(pageSize int, operater *Ec2Operater, rate int) *pager {
	req := &ec2.DescribeInstancesInput{}
	req.MaxResults = utils.Int32Ptr(int32(pageSize))
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
	operater *Ec2Operater
	req      *ec2.DescribeInstancesInput
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Scan(ctx context.Context, set *host.HostSet) error {
	resp, err := p.operater.Query(ctx, p.nextReq())
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

func (p *pager) nextReq() *ec2.DescribeInstancesInput {
	// 等待一个可用token
	p.tb.Wait(1)

	p.log.Debugf("请求第%d页数据", p.number)
	return p.req
}

func (p *pager) Next() bool {
	if p.total == -1 {
		return true
	}
	return int64(p.number*p.size) < p.total
}
