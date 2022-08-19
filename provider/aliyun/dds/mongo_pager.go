package dds

import (
	"context"

	dds "github.com/alibabacloud-go/dds-20151201/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newMongoPager(operator *Operator) pager.Pager {
	bp := pager.NewBasePager()
	//每页记录数，阿里仅支持取值：30, 50, 100
	bp.SetPageSize(50)
	return &mongoPager{
		BasePager: bp,
		operator:  operator,
		req: &dds.DescribeDBInstancesRequest{
			RegionId: operator.client.RegionId,
		},
		log: zap.L().Named("ali.mongo"),
	}
}

type mongoPager struct {
	*pager.BasePager
	operator *Operator
	req      *dds.DescribeDBInstancesRequest
	log      logger.Logger
}

func (p *mongoPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *mongoPager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *mongoPager) nextReq() *dds.DescribeDBInstancesRequest {
	p.req.PageNumber = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	p.log.Debugf("请求第%d页数据, 页面大小:%d", *p.req.PageNumber, *p.req.PageSize)
	return p.req
}
