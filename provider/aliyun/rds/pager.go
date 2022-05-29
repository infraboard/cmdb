package rds

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *RdsOperator) pager.Pager {
	return &rdsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       rds.CreateDescribeDBInstancesRequest(),
		log:       zap.L().Named("ali.rds"),
	}
}

type rdsPager struct {
	*pager.BasePager
	operator *RdsOperator
	req      *rds.DescribeDBInstancesRequest
	log      logger.Logger
}

func (p *rdsPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(set)
	return nil
}

func (p *rdsPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *rdsPager) nextReq() *rds.DescribeDBInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNumber = requests.NewInteger(int(p.PageNumber()))
	p.req.PageSize = requests.NewInteger(int(p.PageSize()))
	return p.req
}
