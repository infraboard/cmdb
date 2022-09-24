package impl

import (
	"database/sql"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *sql.DB
	orm *gorm.DB
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

	db, err := orm.DB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.orm = orm
	s.db = db
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
}
