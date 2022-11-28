package cdb

import (
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCDBOperator(client *cdb.Client) *CDBOperator {
	return &CDBOperator{
		client: client,
		log:    zap.L().Named("tx.cdb"),
	}
}

type CDBOperator struct {
	client *cdb.Client
	log    logger.Logger
}
