package obs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewObsOperator(client *obs.ObsClient) *ObsOperator {
	return &ObsOperator{
		client: client,
		log:    zap.L().Named("huawei.obs"),
	}
}

type ObsOperator struct {
	client *obs.ObsClient
	log    logger.Logger
}
