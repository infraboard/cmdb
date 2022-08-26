package dns

import (
	"encoding/json"

	"github.com/infraboard/cmdb/apps/resource"
)

const (
	AppName = "domain"
)

func NewDefaultDomain() *Domain {
	return &Domain{
		Resource: resource.NewDefaultResource(resource.TYPE_DOMAIN),
		Records:  NewRecordSet(),
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
		ids = append(ids, s.Items[i].Resource.Base.Id)
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

func NewDefaultRecord() *Record {
	return &Record{}
}

func NewRecordSet() *RecordSet {
	return &RecordSet{
		Items: []*Record{},
	}
}

func (s *RecordSet) GetLast() *Record {
	l := s.Length()
	if l == 0 {
		return nil
	}

	return s.Items[l-1]
}

func (s *RecordSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Record))
	}
}

func (s *RecordSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Id)
	}
	return
}

func (s *RecordSet) Length() int64 {
	return int64(len(s.Items))
}

func (s *RecordSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *RecordSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
