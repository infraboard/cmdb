package dns

import (
	"context"

	dom "github.com/alibabacloud-go/domain-20180129/v3/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newDomainPager(operator *DnsOperator) pager.Pager {
	req := &dom.QueryDomainListRequest{}

	return &domainPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.dns.domain"),
	}
}

type domainPager struct {
	*pager.BasePager
	operator *DnsOperator
	req      *dom.QueryDomainListRequest
	log      logger.Logger
}

func (p *domainPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryDomain(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *domainPager) nextReq() *dom.QueryDomainListRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNum = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	return p.req
}
