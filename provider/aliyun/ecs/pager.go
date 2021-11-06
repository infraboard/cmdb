package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/pkg/host"
)

func newPager(pageSize int, operater *EcsOperater, rate int) *pager {
	req := ecs.CreateDescribeInstancesRequest()
	req.PageSize = requests.NewInteger(pageSize)
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
	operater *EcsOperater
	req      *ecs.DescribeInstancesRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Next() *host.PagerResult {
	result := host.NewPagerResult()

	resp, err := p.operater.Query(p.nextReq())
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

func (p *pager) nextReq() *ecs.DescribeInstancesRequest {
	// 等待一个可用token
	p.tb.Wait(1)

	p.log.Debugf("请求第%d页数据", p.number)
	p.req.PageNumber = requests.NewInteger(p.number)
	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}
