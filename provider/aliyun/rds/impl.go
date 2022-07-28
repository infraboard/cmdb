package rds

import (
	rds "github.com/alibabacloud-go/rds-20140815/v2/client"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRdsOperator(client *rds.Client) *RdsOperator {
	return &RdsOperator{
		client: client,
		log:    zap.L().Named("ali.rds"),
	}
}

type RdsOperator struct {
	client *rds.Client
	log    logger.Logger
}
