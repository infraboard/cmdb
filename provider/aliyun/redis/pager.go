package redis

import (
	"context"

	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newPager(operator *RedisOperator) pager.Pager {
	req := &redis.DescribeInstancesRequest{
		RegionId: operator.client.RegionId,
	}

	return &redisPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       req,
		log:       zap.L().Named("ali.redis"),
	}
}

type redisPager struct {
	*pager.BasePager
	operator *RedisOperator
	req      *redis.DescribeInstancesRequest
	log      logger.Logger
}

func (p *redisPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *redisPager) nextReq() *redis.DescribeInstancesRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.PageNumber = tea.Int32(int32(p.PageNumber()))
	p.req.PageSize = tea.Int32(int32(p.PageSize()))
	return p.req
}
