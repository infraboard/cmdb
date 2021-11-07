package pkg

import (
	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/pkg/task"
	"google.golang.org/grpc"
)

var (
	Host     host.ServiceServer
	Secret   secret.ServiceServer
	Resource resource.ServiceServer
	Task     task.ServiceServer
)

// InitV1GRPCAPI 初始化API服务
func InitV1GRPCAPI(server *grpc.Server) {
	resource.RegisterServiceServer(server, Resource)
	host.RegisterServiceServer(server, Host)
	secret.RegisterServiceServer(server, Secret)
	task.RegisterServiceServer(server, Task)
}
