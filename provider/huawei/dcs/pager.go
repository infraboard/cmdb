package dcs

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/utils"
)

func newPager(operator *DcsOperator) pager.Pager {
	return &rdsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       &model.ListInstancesRequest{},
		log:       zap.L().Named("huawei.redis"),
	}
}

type rdsPager struct {
	*pager.BasePager
	operator *DcsOperator
	req      *model.ListInstancesRequest
	log      logger.Logger
}

func (p *rdsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *rdsPager) nextReq() *model.ListInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.PageNumber()))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
