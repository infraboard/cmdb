package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	cmdbRds "github.com/infraboard/cmdb/apps/rds"
)

func newPager(pageSize int, operater *RdsOperater, rate int) *pager {
	req := rds.CreateDescribeDBInstancesRequest()
	req.PageSize = requests.NewInteger(pageSize)
	rateFloat := 1 / float64(rate)

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		total:    -1,
		req:      req,
		log:      zap.L().Named("Pagger"),
		tb:       tokenbucket.NewBucketWithRate(rateFloat, 1),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *RdsOperater
	req      *rds.DescribeDBInstancesRequest
	log      logger.Logger
	tb       *tokenbucket.Bucket
}

func (p *pager) Next() *cmdbRds.PagerResult {
	result := cmdbRds.NewPagerResult()

	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		result.Err = err
		return result
	}

	p.total = int64(resp.Total)

	result.Data = resp
	result.HasNext = p.HasNext()

	p.number++
	return result
}

func (p *pager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *pager) nextReq() *rds.DescribeDBInstancesRequest {
	p.log.Debug("请求第%d页数据", p.number)
	p.req.PageNumber = requests.NewInteger(p.number)
	return p.req
}

func (p *pager) HasNext() bool {
	if p.total == -1 {
		return true
	}
	return int64(p.number*p.size) < p.total
}
