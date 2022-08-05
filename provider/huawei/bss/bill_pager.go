package bss

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
	"github.com/infraboard/cmdb/utils"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *BssOperator, month string) pager.Pager {
	req := &model.ListCustomerselfResourceRecordsRequest{}
	req.Cycle = month

	return &bssPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("huawei.bss"),
	}
}

type bssPager struct {
	*pager.BasePager
	operator *BssOperator
	req      *model.ListCustomerselfResourceRecordsRequest
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

func (p *bssPager) nextReq() *model.ListCustomerselfResourceRecordsRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())

	// 注意: 华为云的Offse表示的是页码
	p.req.Offset = utils.Int32Ptr(int32(p.PageNumber()))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
