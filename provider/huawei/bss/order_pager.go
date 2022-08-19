package bss

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newOrderPager(operator *BssOperator, r *provider.QueryOrderRequest) pager.Pager {
	req := &model.ListCustomerOrdersRequest{
		CreateTimeBegin: tea.String(r.StartTime.UTC().Format(utils.DEFAULT_TIME_SECOND_FORMAT)),
		CreateTimeEnd:   tea.String(r.EndTime.UTC().Format(utils.DEFAULT_TIME_SECOND_FORMAT)),
	}

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
	p.req.Offset = utils.Int32Ptr(int32(p.Offset()))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	p.log.Debugf("请求第%d页数据, 开始时间: %s, 结束时间: %s", p.PageNumber(),
		tea.StringValue(p.req.CreateTimeBegin),
		tea.StringValue(p.req.CreateTimeEnd),
	)
	return p.req
}
