package provider

import (
	"context"

	"github.com/infraboard/mcube/pager"
)

type OssOperator interface {
	QueryBucket(ctx context.Context, req *QueryRequest) pager.Pager
}
