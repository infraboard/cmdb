package pkg

import (
	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/pkg/task"
	"google.golang.org/grpc"
)

var (
	Host     host.Service
	Secret   secret.Service
	Resource resource.ServiceServer
	Task     task.Service
)

// InitV1GRPCAPI 初始化API服务
func InitV1GRPCAPI(server *grpc.Server) {
	resource.RegisterServiceServer(server, Resource)
}
