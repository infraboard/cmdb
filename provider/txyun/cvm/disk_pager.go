package cvm

import (
	"context"

	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newDiskPager(operator *CVMOperator) pager.Pager {
	req := cbs.NewDescribeDisksRequest()

	return &diskPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("tx.disk"),
	}
}

type diskPager struct {
	*pager.BasePager
	operator *CVMOperator
	req      *cbs.DescribeDisksRequest
	log      logger.Logger
}

func (p *diskPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.QueryDisk(ctx, p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)
	p.log.Debugf("get %d hosts", len(resp.Items))

	set.Add(resp.ToAny()...)
	return nil
}

func (p *diskPager) nextReq() *cbs.DescribeDisksRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Uint64Ptr(uint64(p.Offset()))
	p.req.Limit = common.Uint64Ptr(uint64(p.PageSize()))
	return p.req
}
