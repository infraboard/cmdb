package obs

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *ObsOperator) QueryBucket(ctx context.Context, req *provider.QueryRequest) pager.Pager {
	return newPager(o)
}

// 获取桶列表
// 参考: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=OBS&api=ListBuckets
func (o *ObsOperator) Query(req *obs.ListBucketsInput) (*oss.BucketSet, error) {
	set := oss.NewBucketSet()

	resp, err := o.client.ListBuckets(req)
	if err != nil {
		return nil, err
	}

	set.Items = o.transferSet(resp.Buckets).Items

	return set, nil
}

func (o *ObsOperator) transferSet(items []obs.Bucket) *oss.BucketSet {
	set := oss.NewBucketSet()
	for _, b := range items {
		set.Add(o.transferOne(b))
	}
	return set
}

func (o *ObsOperator) transferOne(ins obs.Bucket) *oss.Bucket {
	r := oss.NewDefaultBucket()
	b := r.Resource.Meta
	b.Id = ins.Name

	info := r.Resource.Spec
	info.Name = ins.Name
	info.ResourceType = resource.TYPE_BUCKET
	info.Vendor = resource.VENDOR_TENCENT
	return r
}
