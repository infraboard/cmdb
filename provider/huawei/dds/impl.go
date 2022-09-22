package dds

import (
	dds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dds/v3"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewDdsOperator(client *dds.DdsClient) *DdsOperator {
	return &DdsOperator{
		client: client,
		log:    zap.L().Named("huawei.dds"),
	}
}

type DdsOperator struct {
	client *dds.DdsClient
	log    logger.Logger
}
