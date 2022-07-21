package cos

import (
	"fmt"

	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/infraboard/cmdb/apps/oss"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCosOperator(client *cos.Client) *CosOperator {
	return &CosOperator{
		client:        client,
		log:           zap.L().Named("tx.cos"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type CosOperator struct {
	client *cos.Client
	log    logger.Logger
	*resource.AccountGetter
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
	b := r.Base
	b.Vendor = resource.Vendor_TENCENT
	b.Region = ins.Region
	b.Id = fmt.Sprintf("%s.%s", ins.Region, ins.Name)

	info := r.Information
	info.Name = ins.Name
	return r
}
