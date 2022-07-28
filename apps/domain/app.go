package domain

import (
	"encoding/json"

	"github.com/infraboard/cmdb/apps/resource"
)

const (
	AppName = "domain"
)

func NewDefaultDomain() *Domain {
	return &Domain{
		Base: &resource.Base{
			ResourceType: resource.TYPE_DOMAIN,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func NewDomainSet() *DomainSet {
	return &DomainSet{
		Items: []*Domain{},
	}
}

func (s *DomainSet) GetLast() *Domain {
	l := s.Length()
	if l == 0 {
		return nil
	}

	return s.Items[l-1]
}

func (s *DomainSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Domain))
	}
}

func (s *DomainSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Base.Id)
	}
	return
}

func (s *DomainSet) Length() int64 {
	return int64(len(s.Items))
}

func (s *DomainSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *DomainSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
