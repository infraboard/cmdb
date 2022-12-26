package nlb

import (
	nlb "github.com/alibabacloud-go/nlb-20220430/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewNLBOperator(client *nlb.Client) *NLBOperator {
	return &NLBOperator{
		client: client,
		log:    zap.L().Named("ali.nlb"),
	}
}

type NLBOperator struct {
	client *nlb.Client
	log    logger.Logger
}
