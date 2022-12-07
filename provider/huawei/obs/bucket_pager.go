package obs

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *ObsOperator) pager.Pager {
	return &obsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       &obs.ListBucketsInput{},
		log:       operator.log,
	}
}

type obsPager struct {
	*pager.BasePager
	operator *ObsOperator
	req      *obs.ListBucketsInput
	log      logger.Logger
}

func (p *obsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *obsPager) nextReq() *obs.ListBucketsInput {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	return p.req
}
