package cts

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cts/v3/model"
)

func newPager(operator *Operator, r *provider.QueryEventRequest) pager.Pager {
	req := &model.ListTracesRequest{}
	req.From = tea.Int64(r.StartTime.UnixMilli())
	req.To = tea.Int64(r.EndTime.UnixMilli())
	return &clbPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("huawei.event"),
	}
}

type clbPager struct {
	*pager.BasePager
	operator *Operator
	req      *model.ListTracesRequest
	log      logger.Logger
}

func (p *clbPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryEvents(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d clb", len(resp.Items))

	p.CheckHasNext(resp)
	return nil
}

func (p *clbPager) nextReq() *model.ListTracesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Limit = tea.Int32(int32(p.PageSize()))
	return p.req
}
