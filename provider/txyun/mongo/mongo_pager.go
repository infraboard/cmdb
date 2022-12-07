package mongo

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	mongo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(pageSize int, operator *MongoOperator) pager.Pager {
	return &mongoPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       mongo.NewDescribeDBInstancesRequest(),
		log:       zap.L().Named("tx.mongodb"),
	}
}

type mongoPager struct {
	*pager.BasePager
	operator *MongoOperator
	req      *mongo.DescribeDBInstancesRequest
	log      logger.Logger
}

func (p *mongoPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)
	p.log.Debugf("get %d mongodb instances", len(resp.Items))

	set.Add(resp.ToAny()...)
	return nil
}

func (p *mongoPager) nextReq() *mongo.DescribeDBInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Uint64Ptr(uint64(p.Offset()))
	p.req.Limit = common.Uint64Ptr(uint64(p.PageSize()))
	return p.req
}
