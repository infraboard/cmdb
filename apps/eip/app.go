package eip

import (
	"encoding/json"

	resource "github.com/infraboard/cmdb/apps/resource"
)

const (
	AppName = "eip"
)

func NewDefaultEip() *EIP {
	return &EIP{
		Resource: resource.NewDefaultResource(resource.TYPE_EIP),
		Describe: &Describe{},
	}
}

func NewEIPSet() *EIPSet {
	return &EIPSet{
		Items: []*EIP{},
	}
}

func (s *EIPSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*EIP))
	}
}

func (s *EIPSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Resource.Meta.Id)
	}
	return
}

func (s *EIPSet) Length() int64 {
	return int64(len(s.Items))
}

func (s *EIPSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *EIPSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
