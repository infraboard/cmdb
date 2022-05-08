package cdb

import (
	"context"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func newPager(pageSize int, operater *CDBOperater) *pager {
	req := cdb.NewDescribeDBInstancesRequest()
	req.Limit = common.Uint64Ptr(uint64(pageSize))

	return &pager{
		size:     pageSize,
		number:   1,
		hasNext:  true,
		operater: operater,
		req:      req,
		log:      zap.L().Named("tx.cdb"),
	}
}

type pager struct {
	size     int
	number   int
	hasNext  bool
	operater *CDBOperater
	req      *cdb.DescribeDBInstancesRequest
	log      logger.Logger
}

func (p *pager) Scan(ctx context.Context, set *rds.Set) error {
	resp, err := p.operater.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.Items...)
	p.log.Debugf("get %d hosts", len(resp.Items))

	if set.Length() == 0 {
		p.log.Info("sync complete")
		p.hasNext = false
	}

	p.number++
	return nil
}

func (p *pager) nextReq() *cdb.DescribeDBInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.number)
	p.req.Offset = common.Uint64Ptr(uint64(p.offset()))
	return p.req
}

func (p *pager) Next() bool {
	return p.hasNext
}

func (p *pager) offset() int64 {
	return int64(p.size * (p.number - 1))
}
