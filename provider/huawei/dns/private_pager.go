package dns

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"

	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPrivateZonePager(operator *DnsOperator) pager.Pager {
	req := &model.ListPublicZonesRequest{}

	return &privateZonePager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("hw.zone.private"),
	}
}

type privateZonePager struct {
	*pager.BasePager
	operator *DnsOperator
	req      *model.ListPublicZonesRequest
	log      logger.Logger
}

func (p *privateZonePager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryPublicDomain(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *privateZonePager) nextReq() *model.ListPublicZonesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.Offset()))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
