package cos

import (
	"context"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *CosOperator) pager.Pager {
	req := cvm.NewDescribeInstancesRequest()

	return &cvmPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("tx.cos"),
	}
}

type cvmPager struct {
	*pager.BasePager
	operator *CosOperator
	req      *cvm.DescribeInstancesRequest
	log      logger.Logger
}

func (p *cvmPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.queryBucket(ctx)
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d buckets", len(resp.Items))

	p.CheckHasNext(resp)
	return nil
}
