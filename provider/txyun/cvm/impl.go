package cvm

import (
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCVMOperator(client *cvm.Client) *CVMOperator {
	return &CVMOperator{
		client:        client,
		log:           zap.L().Named("tx.cvm"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type CVMOperator struct {
	client *cvm.Client
	log    logger.Logger
	*resource.AccountGetter
}
