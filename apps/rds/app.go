package rds

import (
	"encoding/json"
	"fmt"
	"strings"

	resource "github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger/zap"
)

const (
	AppName = "rds"
)

func NewSet() *RdsSet {
	return &RdsSet{
		Items: []*Rds{},
	}
}

func (s *RdsSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *RdsSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Rds))
	}
}

func (s *RdsSet) AddSet(set *RdsSet) {
	s.Items = append(s.Items, set.Items...)
}

func (s *RdsSet) Length() int64 {
	return int64(len(s.Items))
}

func NewDefaultRDS() *Rds {
	return &Rds{
		Resource: resource.NewDefaultResource(resource.TYPE_RDS),
		Describe: &Describe{},
	}
}

func (r *Rds) ShortDesc() string {
	return fmt.Sprintf("%s [%s]", r.Resource.Spec.Name, r.Resource.Meta.Id)
}

func (r *Rds) GenHash() error {
	r.Resource.ContentHash.SpecHash = r.Resource.Spec.Hash()
	r.Resource.ContentHash.CustomHash = utils.Hash(r.Describe)
	return nil
}

func (d *Describe) ExtraToJson() string {
	if d != nil && len(d.Extra) > 0 {
		b, err := json.Marshal(d.Extra)
		if err != nil {
			zap.L().Error("marshal rds extra error, %s")
			return ""
		}
		return string(b)
	}
	return ""
}

func (d *Describe) SecurityIpListToString() string {
	return strings.Join(d.SecurityIpList, ",")
}
