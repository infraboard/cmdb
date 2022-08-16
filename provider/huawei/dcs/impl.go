package dcs

import (
	dcs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewDcsOperator(client *dcs.DcsClient) *DcsOperator {
	return &DcsOperator{
		client: client,
		log:    zap.L().Named("huawei.redis"),
	}
}

type DcsOperator struct {
	client *dcs.DcsClient
	log    logger.Logger
}
