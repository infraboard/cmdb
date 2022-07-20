package redis

import (
	"context"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(pageSize int, operator *RedisOperator) pager.Pager {
	return &cdbPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       redis.NewDescribeInstancesRequest(),
		log:       zap.L().Named("tx.redis"),
	}
}

type cdbPager struct {
	*pager.BasePager
	operator *RedisOperator
	req      *redis.DescribeInstancesRequest
	log      logger.Logger
}

func (p *cdbPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)
	p.log.Debugf("get %d mysql", len(resp.Items))

	p.CheckHasNext(set)
	return nil
}

func (p *cdbPager) nextReq() *redis.DescribeInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.Offset = common.Uint64Ptr(uint64(p.Offset()))
	p.req.Limit = common.Uint64Ptr(uint64(p.PageSize()))
	return p.req
}
