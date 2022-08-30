package cts

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	cts "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cts/v3"
)

func NewOperator(client *cts.CtsClient) *Operator {
	return &Operator{
		client: client,
		log:    zap.L().Named("huawei.cts"),
	}
}

type Operator struct {
	client *cts.CtsClient
	log    logger.Logger
}
