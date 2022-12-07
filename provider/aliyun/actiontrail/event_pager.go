package actiontrail

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/pager"

	actiontrail "github.com/alibabacloud-go/actiontrail-20200706/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func newPager(operator *Operator, r *provider.QueryEventRequest, logger logger.Logger) pager.Pager {
	req := &actiontrail.LookupEventsRequest{}
	req.StartTime = tea.String(r.StartTime.UTC().Format(utils.ISO8601_FORMAT))
	req.EndTime = tea.String(r.EndTime.UTC().Format(utils.ISO8601_FORMAT))
	req.LookupAttribute = []*actiontrail.LookupEventsRequestLookupAttribute{
		{
			Key:   tea.String("EventName"),
			Value: tea.String("DeleteInstance"),
		},
	}
	log := logger.Named("event")
	log.Debugf("start: %s, end: %s", *req.StartTime, *req.EndTime)

	return &eventPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       log,
	}
}

type eventPager struct {
	*pager.BasePager
	operator *Operator
	req      *actiontrail.LookupEventsRequest
	log      logger.Logger
}

func (p *eventPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryEvents(ctx, p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)
	p.log.Debugf("get %d event", len(resp.Items))

	set.Add(resp.ToAny()...)
	return nil
}

func (p *eventPager) nextReq() *actiontrail.LookupEventsRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.MaxResults = tea.String(fmt.Sprintf("%d", p.PageSize()))
	return p.req
}
