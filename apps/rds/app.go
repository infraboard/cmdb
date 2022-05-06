package rds

import (
	"context"
	"fmt"

	resource "github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
)

const (
	AppName = "rds"
)

func NewSet() *Set {
	return &Set{
		Items: []*RDS{},
	}
}

func (s *Set) Add(items ...*RDS) {
	s.Items = append(s.Items, items...)
}

func (s *Set) AddSet(set *Set) {
	s.Items = append(s.Items, set.Items...)
}

func NewDefaultRDS() *RDS {
	return &RDS{
		Base: &resource.Base{
			ResourceType: resource.Type_RDS,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

// 分页迭代器
type Pager interface {
	Next() bool
	Scan(context.Context, *Set) error
}

func (r *RDS) ShortDesc() string {
	return fmt.Sprintf("%s %s", r.Information.Name, r.Information.PrivateIp)
}

func (d *Describe) Hash() string {
	return utils.Hash(d)
}
