package ecs

import (
	"context"

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
		total:    -1,
		operater: operater,
		req:      req,
		log:      zap.L().Named("huawei.ecs"),
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

func (p *pager) Scan(ctx context.Context, set *host.HostSet) error {
	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		return err
	}

	set.Add(resp.Items...)
	p.total = resp.Total

	p.number++
	return nil
}

func (p *pager) nextReq() *model.ListServersDetailsRequest {
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
