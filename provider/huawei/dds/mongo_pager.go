package dds

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dds/v3/model"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/utils"
)

func newPager(operator *DdsOperator) pager.Pager {
	return &ddsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       &model.ListInstancesRequest{},
		log:       zap.L().Named("huawei.dds"),
	}
}

type ddsPager struct {
	*pager.BasePager
	operator *DdsOperator
	req      *model.ListInstancesRequest
	log      logger.Logger
}

func (p *ddsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *ddsPager) nextReq() *model.ListInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = utils.Int32Ptr(int32(p.Offset()))
	p.req.Limit = utils.Int32Ptr(int32(p.PageSize()))
	return p.req
}
