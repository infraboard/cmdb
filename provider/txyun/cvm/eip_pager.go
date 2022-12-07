package cvm

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newEipPager(operator *CVMOperator) pager.Pager {
	req := vpc.NewDescribeAddressesRequest()

	return &eipPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("tx.eip"),
	}
}

type eipPager struct {
	*pager.BasePager
	operator *CVMOperator
	req      *vpc.DescribeAddressesRequest
	log      logger.Logger
}

func (p *eipPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryEip(ctx, p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)
	p.log.Debugf("get %d eips", len(resp.Items))

	set.Add(resp.ToAny()...)
	return nil
}

func (p *eipPager) nextReq() *vpc.DescribeAddressesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Int64Ptr(int64(p.Offset()))
	p.req.Limit = common.Int64Ptr(int64(p.PageSize()))
	return p.req
}
