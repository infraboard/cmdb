package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/utils"
)

func newPager(operator *Ec2operator) pager.Pager {
	req := &ec2.DescribeInstancesInput{}

	return &ec2Pager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("Pagger"),
	}
}

type ec2Pager struct {
	*pager.BasePager
	operator *Ec2operator
	req      *ec2.DescribeInstancesInput
	log      logger.Logger
}

func (p *ec2Pager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(set)
	return nil
}

func (p *ec2Pager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *ec2Pager) nextReq() *ec2.DescribeInstancesInput {
	p.req.MaxResults = utils.Int32Ptr(int32(p.PageSize()))
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	return p.req
}
