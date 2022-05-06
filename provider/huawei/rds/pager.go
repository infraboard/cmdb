package rds

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/utils"
)

func newPager(pageSize int, operater *RdsOperater) *pager {
	req := &model.ListInstancesRequest{}
	req.Limit = utils.Int32Ptr(int32(pageSize))

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		total:    -1,
		log:      zap.L().Named("Pagger"),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *RdsOperater
	req      *model.ListInstancesRequest
	log      logger.Logger
}

func (p *pager) Scan(ctx context.Context, set *rds.Set) error {
	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.Items...)
	p.total = resp.Total

	p.number++
	return nil
}

func (p *pager) nextReq() *model.ListInstancesRequest {
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
