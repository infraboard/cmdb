package alb

import (
	alb "github.com/alibabacloud-go/alb-20200616/v2/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewALBOperator(client *alb.Client) *ALBOperator {
	return &ALBOperator{
		client: client,
		log:    zap.L().Named("ali.slb"),
	}
}

type ALBOperator struct {
	client *alb.Client
	log    logger.Logger
}
