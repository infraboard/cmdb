package sqlserver

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sqlserver "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver/v20180328"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(pageSize int, operator *SQLServerOperator) pager.Pager {
	return &sqlserverPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       sqlserver.NewDescribeDBInstancesRequest(),
		log:       zap.L().Named("tx.cdb"),
	}
}

type sqlserverPager struct {
	*pager.BasePager
	operator *SQLServerOperator
	req      *sqlserver.DescribeDBInstancesRequest
	log      logger.Logger
}

func (p *sqlserverPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)
	p.log.Debugf("get %d sqlserver", len(resp.Items))

	set.Add(resp.ToAny()...)
	return nil
}

func (p *sqlserverPager) nextReq() *sqlserver.DescribeDBInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Int64Ptr(int64(p.Offset()))
	p.req.Limit = common.Int64Ptr(int64(p.PageSize()))
	return p.req
}
