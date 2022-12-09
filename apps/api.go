package apps

import (
	// 内置健康检查
	_ "github.com/infraboard/mcube/app/health/api"

	_ "github.com/infraboard/cmdb/apps/bill/api"
	_ "github.com/infraboard/cmdb/apps/dict/api"
	_ "github.com/infraboard/cmdb/apps/host/api"
	_ "github.com/infraboard/cmdb/apps/resource/api"
	_ "github.com/infraboard/cmdb/apps/secret/api"
	_ "github.com/infraboard/cmdb/apps/task/api"
)
