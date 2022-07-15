package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	cmdbOss "github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *OssOperator) QueryBucket(req *provider.QueryBucketRequest) pager.Pager {
	p := newBucketPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 列举请求者拥有的所有存储空间（Bucket）
// 参考文档: https://next.api.aliyun.com/api/Oss/2019-05-17/ListBuckets?params={}&sdkStyle=dara
func (o *OssOperator) query(req *listBucketRequest) (*cmdbOss.BucketSet, error) {
	resp, err := o.client.ListBuckets(req.Options()...)
	if err != nil {
		return nil, err
	}
	req.marker = resp.Marker

	set := cmdbOss.NewBucketSet()
	set.Items = o.transferSet(resp).Items
	return set, nil
}

type listBucketRequest struct {
	marker   string
	pageSize int
}

func (req *listBucketRequest) Options() (options []oss.Option) {
	options = append(options,
		oss.Marker(req.marker),
		oss.MaxKeys(req.pageSize),
	)
	return
}
