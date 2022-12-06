package oss

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/pager"
)

func newBucketPager(operator *OssOperator) pager.Pager {
	return &bucketPager{
		BasePager: pager.NewBasePager(),
		operator:  operator,
		req:       &listBucketRequest{},
		log:       operator.log,
	}
}

type bucketPager struct {
	*pager.BasePager
	operator *OssOperator
	log      logger.Logger
	req      *listBucketRequest
}

func (p *bucketPager) Scan(ctx context.Context, set pager.Set) error {
	resp, err := p.operator.query(p.nextReq())
	if err != nil {
		return err
	}
	p.CheckHasNext(resp)

	set.Add(resp.ToAny()...)
	return nil
}

func (p *bucketPager) nextReq() *listBucketRequest {
	p.log.Debugf("请求第%d页数据", p.PageNumber())
	p.req.pageSize = int(p.PageSize())
	return p.req
}
