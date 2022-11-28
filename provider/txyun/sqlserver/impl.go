package sqlserver

import (
	sqlserver "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver/v20180328"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewSQLServerOperator(client *sqlserver.Client) *SQLServerOperator {
	return &SQLServerOperator{
		client: client,
		log:    zap.L().Named("tx.cdb"),
	}
}

type SQLServerOperator struct {
	client *sqlserver.Client
	log    logger.Logger
}
