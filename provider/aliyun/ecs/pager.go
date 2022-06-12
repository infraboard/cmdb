package ecs

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *EcsOperator) pager.Pager {
	req := ecs.CreateDescribeInstancesRequest()

	return &ecsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.ecs"),
	}
}

type ecsPager struct {
	*pager.BasePager
	operator *EcsOperator
	req      *ecs.DescribeInstancesRequest
	log      logger.Logger
}

func (p *ecsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(set)
	return nil
}

func (p *ecsPager) nextReq() *ecs.DescribeInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNumber = requests.NewInteger(int(p.PageNumber()))
	p.req.PageSize = requests.NewInteger(int(p.PageSize()))
	return p.req
}
