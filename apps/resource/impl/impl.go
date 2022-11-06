package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *gorm.DB
	log logger.Logger

	resource.UnimplementedRPCServer
}

func (s *service) Config() error {
	orm, err := conf.C().MySQL.ORM()
	if err != nil {
		return err
	}
	// 在冲突时，更新除主键以外的所有列到新值。
	orm = orm.Clauses(clause.OnConflict{
		UpdateAll: true,
	})
	// 是否开启debug
	if conf.C().Log.Level == "debug" {
		orm.Debug()
	}

	s.log = zap.L().Named(s.Name())
	s.db = orm
	return nil
}

func (s *service) Name() string {
	return resource.AppName
}

func (s *service) Registry(server *grpc.Server) {
	resource.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
	app.RegistryInternalApp(svr)
}
