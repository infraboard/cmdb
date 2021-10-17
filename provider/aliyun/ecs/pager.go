package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func newPagger(pageSize int, client *ecs.Client) *pager {
	req := ecs.CreateDescribeInstancesRequest()
	req.PageSize = requests.NewInteger(pageSize)

	return &pager{
		size:   pageSize,
		number: 1,
		client: client,
		req:    req,
		log:    zap.L().Named("Pagger"),
	}
}

type pager struct {
	size   int
	number int
	total  int64
	client *ecs.Client
	req    *ecs.DescribeInstancesRequest
	log    logger.Logger
}

func (p *pager) Next() *host.PagerResult {
	result := host.NewPagerResult()

	resp, err := p.client.DescribeInstances(p.nextReq())
	if err != nil {
		result.Err = err
		return result
	}
	p.total = int64(resp.TotalCount)

	result.Data = transferSet(resp.Instances.Instance)
	result.Data.Total = p.total
	result.HasNext = p.hasNext()

	p.number++
	return result
}

func (p *pager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *pager) nextReq() *ecs.DescribeInstancesRequest {
	p.log.Debug("请求第%d页数据", p.number)
	p.req.PageNumber = requests.NewInteger(p.number)
	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}
