package cos

import (
	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCosOperator(client *cos.Client) *CosOperator {
	return &CosOperator{
		client:        client,
		log:           zap.L().Named("tx.cos"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type CosOperator struct {
	client *cos.Client
	log    logger.Logger
	*resource.AccountGetter
}
