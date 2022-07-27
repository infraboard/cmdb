package elb

import (
	elb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewELBOperator(client *elb.ElbClient) *ELBOperator {
	return &ELBOperator{
		client: client,
		log:    zap.L().Named("hw.elb"),
	}
}

type ELBOperator struct {
	client *elb.ElbClient
	log    logger.Logger
}
