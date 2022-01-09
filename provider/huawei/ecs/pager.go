package ecs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/utils"
)

func newPager(pageSize int, operater *EcsOperater) *pager {
	req := &model.ListServersDetailsRequest{}
	req.Limit = utils.Int32Ptr(int32(pageSize))

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
	operater *EcsOperater
	req      *model.ListServersDetailsRequest
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

func (p *pager) nextReq() *model.ListServersDetailsRequest {
	p.log.Debugf("请求第%d页数据", p.number)

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.number))
	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}
