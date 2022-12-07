package cloudaudit

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	cloudaudit "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit/v20190319"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *Operator, r *provider.QueryEventRequest) pager.Pager {
	req := cloudaudit.NewDescribeEventsRequest()
	req.StartTime = tea.Uint64(uint64(r.StartTime.Unix()))
	req.EndTime = tea.Uint64(uint64(r.EndTime.Unix()))
	req.LookupAttributes = []*cloudaudit.LookupAttribute{
		{AttributeKey: tea.String("EventName"), AttributeValue: tea.String("StartInstances")},
	}
	return &clbPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("tx.cloudaudit"),
	}
}

type clbPager struct {
	*pager.BasePager
	operator *Operator
	req      *cloudaudit.DescribeEventsRequest
	log      logger.Logger
}

func (p *clbPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryEvents(ctx, p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)
	p.log.Debugf("get %d clb", len(resp.Items))

	set.Add(resp.ToAny()...)
	return nil
}

func (p *clbPager) nextReq() *cloudaudit.DescribeEventsRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.MaxResults = tea.Uint64(uint64(p.PageSize()))
	return p.req
}
