package cvm

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func newPager(pageSize int, operater *CVMOperater) *pager {
	req := cvm.NewDescribeInstancesRequest()
	req.Limit = common.Int64Ptr(int64(pageSize))

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		log:      zap.L().Named("Pagger"),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *CVMOperater
	req      *cvm.DescribeInstancesRequest
	log      logger.Logger
}

func (p *pager) Next() *host.PagerResult {
	result := host.NewPagerResult()

	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		result.Err = err
		return result
	}
	p.total = resp.Total

	result.Data = resp
	result.HasNext = p.hasNext()

	p.number++
	return result
}

func (p *pager) nextReq() *cvm.DescribeInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.number)
	p.req.Offset = common.Int64Ptr(p.offset())
	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}

func (p *pager) offset() int64 {
	return int64(p.size * (p.number - 1))
}
