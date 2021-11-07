package impl

import (
	"database/sql"

	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/cmdb/pkg"
	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/pkg/task"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	db     *sql.DB
	log    logger.Logger
	host   host.ServiceServer
	secret secret.Service
	task.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named("Task")
	s.db = db
	s.host = pkg.Host
	s.secret = pkg.Secret
	return nil
}
