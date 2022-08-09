package bss

import (
	"context"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newOrderPager(operator *BssOperator, r *provider.QueryOrderRequest) pager.Pager {
	req := &bssopenapi.QueryOrdersRequest{}
	req.CreateTimeEnd = tea.String(r.EndTime.Format(utils.DEFAULT_TIME_SECOND_FORMAT))
	req.CreateTimeStart = tea.String(r.StartTime.Format(utils.DEFAULT_TIME_SECOND_FORMAT))
	return &orderPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.order"),
	}
}

type orderPager struct {
	*pager.BasePager
	operator  *BssOperator
	req       *bssopenapi.QueryOrdersRequest
	log       logger.Logger
	nextToken string
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

func (p *orderPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *orderPager) nextReq() *bssopenapi.QueryOrdersRequest {
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	p.req.PageNum = tea.Int32(int32(p.PageNumber()))
	p.log.Debugf("请求参数: %s", p.req.String())
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	return p.req
}
