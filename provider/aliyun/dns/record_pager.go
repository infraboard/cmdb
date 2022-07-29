package dns

import (
	"context"

	alidns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newRecordPager(operator *DnsOperator) *recordPager {
	req := &alidns.DescribeDomainRecordsRequest{}

	return &recordPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.dns.record"),
	}
}

type recordPager struct {
	*pager.BasePager
	operator *DnsOperator
	req      *alidns.DescribeDomainRecordsRequest
	log      logger.Logger
}

func (p *recordPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryRecord(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *recordPager) nextReq() *alidns.DescribeDomainRecordsRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNumber = tea.Int64(p.PageNumber())
	p.req.PageSize = tea.Int64(p.PageSize())
	return p.req
}
