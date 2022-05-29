package cdb

import (
	"context"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(pageSize int, operator *CDBOperator) pager.Pager {
	return &cdbPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       cdb.NewDescribeDBInstancesRequest(),
		log:       zap.L().Named("tx.cdb"),
	}
}

type cdbPager struct {
	*pager.BasePager
	operator *CDBOperator
	req      *cdb.DescribeDBInstancesRequest
	log      logger.Logger
}

func (p *cdbPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d hosts", len(resp.Items))

	p.CheckHasNext(set)
	return nil
}

func (p *cdbPager) nextReq() *cdb.DescribeDBInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Uint64Ptr(uint64(p.Offset()))
	p.req.Limit = common.Uint64Ptr(uint64(p.PageSize()))
	return p.req
}
