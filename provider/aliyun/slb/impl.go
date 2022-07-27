package slb

import (
	slb "github.com/alibabacloud-go/slb-20140515/v3/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewSLBOperator(client *slb.Client) *SLBOperator {
	return &SLBOperator{
		client: client,
		log:    zap.L().Named("ali.slb"),
	}
}

type SLBOperator struct {
	client *slb.Client
	log    logger.Logger
}
