package ecs

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/utils"
)

func newEipPager(operator *EcsOperator) pager.Pager {
	req := &model.ListPublicipsRequest{}

	return &eipPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("huawei.eip"),
	}
}

type eipPager struct {
	*pager.BasePager
	operator *EcsOperator
	req      *model.ListPublicipsRequest
	log      logger.Logger
}

func (p *eipPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryEip(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *eipPager) nextReq() *model.ListPublicipsRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())

	// 注意: 华为云的Offse表示的是页码
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
