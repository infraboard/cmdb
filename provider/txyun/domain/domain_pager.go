package domain

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newDomainPager(operator *DnsOperator) pager.Pager {
	return &domainPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       dnspod.NewDescribeDomainListRequest(),
		log:       zap.L().Named("tx.domain"),
	}
}

type domainPager struct {
	*pager.BasePager
	operator *DnsOperator
	req      *dnspod.DescribeDomainListRequest
	log      logger.Logger
}

func (p *domainPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryDomain(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d mysql", len(resp.Items))

	p.CheckHasNext(resp)
	return nil
}

func (p *domainPager) nextReq() *dnspod.DescribeDomainListRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Int64Ptr(int64(p.Offset()))
	p.req.Limit = common.Int64Ptr(int64(p.PageSize()))
	return p.req
}
