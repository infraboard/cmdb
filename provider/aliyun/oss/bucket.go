package oss

import (
	"context"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	cmdbOss "github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *OssOperator) QueryBucket(ctx context.Context, req *provider.QueryRequest) pager.Pager {
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
	req.marker = resp.NextMarker

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

func (o *OssOperator) transferSet(items oss.ListBucketsResult) *cmdbOss.BucketSet {
	set := cmdbOss.NewBucketSet()
	for _, item := range items.Buckets {
		set.Add(o.transferBucket(item))
	}
	return set
}

func (o *OssOperator) transferBucket(ins oss.BucketProperties) *cmdbOss.Bucket {
	r := cmdbOss.NewDefaultBucket()
	region := strings.TrimPrefix(ins.Location, "oss-")

	b := r.Resource.Meta
	b.Id = resource.GenResourceHashId(resource.TYPE_BUCKET, o.account, region, ins.Name)
	b.CreateAt = ins.CreationDate.Unix()

	info := r.Resource.Spec
	info.Owner = o.account
	info.Name = ins.Name
	info.Vendor = resource.VENDOR_ALIYUN
	info.Region = region
	info.ResourceType = resource.TYPE_BUCKET

	desc := r.Describe
	desc.StorageClass = ins.StorageClass
	return r
}
