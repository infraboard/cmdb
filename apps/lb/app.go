package lb

import (
	"encoding/json"

	resource "github.com/infraboard/cmdb/apps/resource"
)

const (
	AppName = "lb"
)

func NewDefaultLB() *LB {
	return &LB{
		Base: &resource.Base{
			ResourceType: resource.TYPE_LB,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func NewLBSet() *LBSet {
	return &LBSet{
		Items: []*LB{},
	}
}

func (s *LBSet) GetLast() *LB {
	l := s.Length()
	if l == 0 {
		return nil
	}

	return s.Items[l-1]
}

func (s *LBSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*LB))
	}
}

func (s *LBSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Base.Id)
	}
	return
}

func (s *LBSet) Length() int64 {
	return int64(len(s.Items))
}

func (s *LBSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *LBSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
