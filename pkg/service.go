package pkg

import (
	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/syncer"
)

var (
	Host     host.Service
	Syncer   syncer.Service
	Resource resource.Service
)
