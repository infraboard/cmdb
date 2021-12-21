package rds

import (
	"fmt"

	"github.com/infraboard/cmdb/app/resource"
)

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewSet(),
	}
}

func NewSet() *Set {
	return &Set{
		Items: []*RDS{},
	}
}

func (s *Set) Add(item *RDS) {
	s.Items = append(s.Items, item)
}

func (s *Set) AddSet(set *Set) {
	s.Items = append(s.Items, set.Items...)
}

func NewDefaultRDS() *RDS {
	return &RDS{
		Base:        &resource.Base{},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

type PagerResult struct {
	Data    *Set
	Err     error
	HasNext bool
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}

func (r *RDS) ShortDesc() string {
	return fmt.Sprintf("%s %s", r.Information.Name, r.Information.PrivateIp)
}
