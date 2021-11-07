package app

import (
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

// GRPCService GRPC服务的实例
type GRPCApp interface {
	Registry(*grpc.Server) error
	Config() error
	Name() string
}

// HTTPService Http服务的实例
type HTTPApp interface {
	Registry(*httprouter.Router) error
	Config() error
	Name() string
}
