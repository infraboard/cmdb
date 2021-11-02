package pkg

import (
	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/pkg/task"
)

var (
	Host     host.Service
	Secret   secret.Service
	Resource resource.Service
	Task     task.Service
)
