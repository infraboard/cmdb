package provider

import "github.com/infraboard/mcube/pager"

type OssOperator interface {
	QueryBucket(req *QueryBucketRequest) pager.Pager
}

func NewQueryBucketRate(rate int32) *QueryBucketRequest {
	return &QueryBucketRequest{
		Rate: float64(rate),
	}
}

type QueryBucketRequest struct {
	Rate float64 `json:"rate"`
}
