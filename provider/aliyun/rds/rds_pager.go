package rds

import (
	"context"

	rds "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *RdsOperator) pager.Pager {
	return &rdsPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req: &rds.DescribeDBInstancesRequest{
			RegionId: operator.client.RegionId,
		},
		log: zap.L().Named("ali.rds"),
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

	p.CheckHasNext(resp)
	return nil
}

func (p *rdsPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *rdsPager) nextReq() *rds.DescribeDBInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNumber = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	return p.req
}
