package rds

import "github.com/infraboard/cmdb/app/resource"

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
