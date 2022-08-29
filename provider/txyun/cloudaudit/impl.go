package cloudaudit

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	cloudaudit "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit/v20190319"
)

func NewOperator(client *cloudaudit.Client) *Operator {
	return &Operator{
		client: client,
		log:    zap.L().Named("tx.cloudaudit"),
	}
}

type Operator struct {
	client *cloudaudit.Client
	log    logger.Logger
}
