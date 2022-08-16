package rds

import (
	hw_rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRdsOperator(client *hw_rds.RdsClient) *RdsOperator {
	return &RdsOperator{
		client: client,
		log:    zap.L().Named("hw.rds"),
	}
}

type RdsOperator struct {
	client *hw_rds.RdsClient
	log    logger.Logger
}
