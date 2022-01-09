package impl

import (
	"database/sql"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例

	svr = &service{}
)

type service struct {
	db  *sql.DB
	log logger.Logger
	bill.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db
	return nil
}

func (s *service) Name() string {
	return bill.AppName
}

func (s *service) Registry(server *grpc.Server) {
	bill.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
