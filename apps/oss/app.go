package oss

import resource "github.com/infraboard/cmdb/apps/resource"

const (
	AppName = "oss"
)

func NewBucketSet() *BucketSet {
	return &BucketSet{
		Items: []*Bucket{},
	}
}

func (s *BucketSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *BucketSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Bucket))
	}
}

func (s *BucketSet) AddSet(set *BucketSet) {
	s.Items = append(s.Items, set.Items...)
}

func (s *BucketSet) Length() int64 {
	return int64(len(s.Items))
}

func NewDefaultBucket() *Bucket {
	return &Bucket{
		Base: &resource.Base{
			ResourceType: resource.Type_BUCKET,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}
