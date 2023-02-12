package resource

import (
	"encoding/json"
	"sort"

	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger/zap"
)

func (c *Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*Meta
		*Spec
		*Cost
		*Status
		*ContentHash
		RelatedResources []*Resource
	}{c.Meta, c.Spec, c.Cost, c.Status, c.ContentHash, c.RelatedResources})
}

func (i *Resource) SortTag() {
	sort.Slice(i.Spec.Tags, func(m, n int) bool {
		return i.Spec.Tags[m].Weight < i.Spec.Tags[n].Weight
	})
}

func (i *Spec) Hash() string {
	return utils.Hash(i)
}

func (r *Resource) GetTagValueOne(key string) string {
	tags := r.Spec.Tags
	for i := range tags {
		if tags[i].Key == key {
			return tags[i].Value
		}
	}

	return ""
}

func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Items: []*Resource{},
	}
}

func (s *ResourceSet) Metas() (metas []*Meta) {
	for i := range s.Items {
		metas = append(metas, s.Items[i].Meta)
	}
	return
}

func (s *ResourceSet) Add(item *Resource) {
	s.Items = append(s.Items, item)
}

func (s *ResourceSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Meta.Id)
	}

	return
}

func (s *ResourceSet) PrometheusFormat() (targets []*PrometheusTarget) {
	for i := range s.Items {
		item := s.Items[i]
		if item.GetTagValueOne(PROMETHEUS_SCRAPE) == "true" {
			t, err := item.PrometheusTarget()
			if err != nil {
				zap.L().Errorf("new Prometheus Target errror, %s", err)
				continue
			}
			targets = append(targets, t)
		}
	}
	return
}
