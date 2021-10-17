package cvm

import (
	"github.com/infraboard/cmdb/pkg/host"
)

func (o *CVMOperater) Query() host.Pager {
	return newPagger(20, o.client)
}
