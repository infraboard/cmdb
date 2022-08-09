package bss

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/utils"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newOrderPager(operator *BssOperator) pager.Pager {
	req := &model.ListCustomerOrdersRequest{}

	return &orderPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("huawei.order"),
	}
}

type orderPager struct {
	*pager.BasePager
	operator *BssOperator
	req      *model.ListCustomerOrdersRequest
	log      logger.Logger
}

func (p *orderPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.doQueryOrder(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(set)
	return nil
}

func (p *orderPager) nextReq() *model.ListCustomerOrdersRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.PageNumber()))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
