package ecs

import (
	"github.com/infraboard/cmdb/pkg/host"
)

func (o *EcsOperater) Query() host.Pager {
	return newPagger(20, o.client)
}
