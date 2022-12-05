package cos

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func (o *CosOperator) QueryBucket(ctx context.Context, req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询请求者名下的所有存储桶列表或特定地域下的存储桶列表
// 参考: https://console.cloud.tencent.com/api/explorer?Product=cos&Version=2018-11-26&Action=GetService&SignVersion=
func (o *CosOperator) queryBucket(ctx context.Context) (*oss.BucketSet, error) {
	resp, _, err := o.client.Service.Get(ctx)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Buckets)
	return set, nil
}

func (o *CosOperator) transferSet(items []cos.Bucket) *oss.BucketSet {
	set := oss.NewBucketSet()
	for _, b := range items {
		set.Add(o.transferOne(b))
	}
	return set
}

func (o *CosOperator) transferOne(ins cos.Bucket) *oss.Bucket {
	r := oss.NewDefaultBucket()
	b := r.Resource.Meta
	b.Id = fmt.Sprintf("%s.%s", ins.Region, ins.Name)

	info := r.Resource.Spec
	info.Name = ins.Name
	info.ResourceType = resource.TYPE_BUCKET
	info.Vendor = resource.VENDOR_TENCENT
	info.Region = ins.Region
	return r
}
