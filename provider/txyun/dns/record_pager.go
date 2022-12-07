package dns

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newRecordPager(operator *DnsOperator) *recordPager {
	return &recordPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       dnspod.NewDescribeRecordListRequest(),
		log:       zap.L().Named("tx.dns.record"),
	}
}

type recordPager struct {
	*pager.BasePager
	operator *DnsOperator
	req      *dnspod.DescribeRecordListRequest
	log      logger.Logger
}

func (p *recordPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryRecord(ctx, p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)
	p.log.Debugf("get domain %s %d dns record", *p.req.Domain, len(resp.Items))

	set.Add(resp.ToAny()...)
	return nil
}

func (p *recordPager) nextReq() *dnspod.DescribeRecordListRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Uint64Ptr(uint64(p.Offset()))
	p.req.Limit = common.Uint64Ptr(uint64(p.PageSize()))
	return p.req
}
