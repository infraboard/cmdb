# 服务健康检查

在 加载由mcube提供的内置健康检查API, 内置检查API依赖业务自己实现一个健康检查服务实例: health.AppName, 然后托管给ioc
```go
package apps

import (
	// 内置健康检查
	_ "github.com/infraboard/mcube/app/health/api"

    ...
)
```