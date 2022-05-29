package bss

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *BssOperator, month string) pager.Pager {
	req := bssopenapi.CreateQueryInstanceBillRequest()
	req.BillingCycle = month

	return &bssPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.bss"),
	}
}

type bssPager struct {
	*pager.BasePager
	operator *BssOperator
	req      *bssopenapi.QueryInstanceBillRequest
	log      logger.Logger
}

func (p *bssPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(set)
	return nil
}

func (p *bssPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *bssPager) nextReq() *bssopenapi.QueryInstanceBillRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNum = requests.NewInteger(int(p.PageNumber()))
	p.req.PageSize = requests.NewInteger(int(p.PageSize()))
	return p.req
}
