package impl

import (
	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *gorm.DB
	log logger.Logger

	secret.UnimplementedRPCServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.ORM()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db
	return nil
}

func (s *service) Name() string {
	return secret.AppName
}

func (s *service) Registry(server *grpc.Server) {
	secret.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
