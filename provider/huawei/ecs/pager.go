package ecs

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/utils"
)

func newPager(operator *EcsOperator) pager.Pager {
	req := &model.ListServersDetailsRequest{}

	return &ecsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("huawei.ecs"),
	}
}

type ecsPager struct {
	*pager.BasePager
	operator *EcsOperator
	req      *model.ListServersDetailsRequest
	log      logger.Logger
}

func (p *ecsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	if int64(resp.Length()) < p.PageSize {
		p.HasNext = false
	}

	p.PageNumber++
	return nil
}

func (p *ecsPager) nextReq() *model.ListServersDetailsRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber)

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.PageNumber))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize))
	return p.req
}
