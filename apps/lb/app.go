package lb

import (
	"encoding/json"

	resource "github.com/infraboard/cmdb/apps/resource"
)

const (
	AppName = "lb"
)

func NewDefaultLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		Resource: resource.NewDefaultResource(resource.TYPE_LB),
		Describe: &Describe{},
	}
}

func NewLoadBalancerSet() *LoadBalancerSet {
	return &LoadBalancerSet{
		Items: []*LoadBalancer{},
	}
}

func (s *LoadBalancerSet) GetLast() *LoadBalancer {
	l := s.Length()
	if l == 0 {
		return nil
	}

	return s.Items[l-1]
}

func (s *LoadBalancerSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*LoadBalancer))
	}
}

func (s *LoadBalancerSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Resource.Base.Id)
	}
	return
}

func (s *LoadBalancerSet) Length() int64 {
	return int64(len(s.Items))
}

func (s *LoadBalancerSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *LoadBalancerSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
