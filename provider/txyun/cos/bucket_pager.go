package cos

import (
	"context"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *CosOperator) pager.Pager {
	req := cvm.NewDescribeInstancesRequest()

	return &cosPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("tx.cos"),
	}
}

type cosPager struct {
	*pager.BasePager
	operator *CosOperator
	req      *cvm.DescribeInstancesRequest
	log      logger.Logger
}

func (p *cosPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryBucket(ctx)
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d buckets", len(resp.Items))

	// 无分页
	resp.Items = []*oss.Bucket{}
	p.CheckHasNext(resp)
	return nil
}
