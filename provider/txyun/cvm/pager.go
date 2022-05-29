package cvm

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *CVMOperator) pager.Pager {
	req := cvm.NewDescribeInstancesRequest()

	return &cvmPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("tx.cvm"),
	}
}

type cvmPager struct {
	*pager.BasePager
	operator *CVMOperator
	req      *cvm.DescribeInstancesRequest
	log      logger.Logger
}

func (p *cvmPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d hosts", len(resp.Items))

	p.CheckHasNext(set)
	return nil
}

func (p *cvmPager) nextReq() *cvm.DescribeInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Int64Ptr(p.Offset())
	p.req.Limit = common.Int64Ptr(int64(p.PageSize()))
	return p.req
}
