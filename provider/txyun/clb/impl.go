package clb

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
)

func NewCLBOperator(client *clb.Client) *CLBOperator {
	return &CLBOperator{
		client: client,
		log:    zap.L().Named("tx.clb"),
	}
}

type CLBOperator struct {
	client *clb.Client
	log    logger.Logger
}
