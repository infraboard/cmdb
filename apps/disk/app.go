package disk

import (
	"encoding/json"

	resource "github.com/infraboard/cmdb/apps/resource"
)

const (
	AppName = "disk"
)

func NewDefaultDisk() *Disk {
	return &Disk{
		Resource: &resource.Resource{
			Base: &resource.Base{
				ResourceType: resource.TYPE_DISK,
			},
			Information: &resource.Information{},
			Tags:        []*resource.Tag{},
		},
		Describe: &Describe{},
	}
}

func NewDiskSet() *DiskSet {
	return &DiskSet{
		Items: []*Disk{},
	}
}

func (s *DiskSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Disk))
	}
}

func (s *DiskSet) ResourceIds() (ids []string) {
	for i := range s.Items {
		ids = append(ids, s.Items[i].Resource.Base.Id)
	}
	return
}

func (s *DiskSet) Length() int64 {
	return int64(len(s.Items))
}

func (s *DiskSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *DiskSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
