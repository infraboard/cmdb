package ecs

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/utils"
)

func newDiskPager(operator *EcsOperator) pager.Pager {
	req := &model.ListVolumesRequest{}

	return &diskPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("huawei.disk"),
	}
}

type diskPager struct {
	*pager.BasePager
	operator *EcsOperator
	req      *model.ListVolumesRequest
	log      logger.Logger
}

func (p *diskPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryDisk(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *diskPager) nextReq() *model.ListVolumesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.Offset()))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
