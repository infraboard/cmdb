package oss

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pager"
)

func newBucketPager(operator *OssOperator) pager.Pager {
	return &bucketPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       &listBucketRequest{},
		log:       zap.L().Named("ali.redis"),
	}
}

type bucketPager struct {
	*pager.BasePager
	operator *OssOperator
	log      logger.Logger
	req      *listBucketRequest
	marker   string
}

func (p *bucketPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.query(p.nextReq())
	if err != nil {
		return err
	}
	set.Add(resp.ToAny()...)

	p.CheckHasNext(resp)
	return nil
}

func (p *bucketPager) nextReq() *listBucketRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.pageSize = int(p.PageSize())
	p.req.marker = p.marker
	return p.req
}
