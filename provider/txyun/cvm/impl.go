package cvm

import (
	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCVMOperator(client *cvm.Client, cbs *cbs.Client) *CVMOperator {
	return &CVMOperator{
		cbs:           cbs,
		client:        client,
		log:           zap.L().Named("tx.cvm"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type CVMOperator struct {
	client *cvm.Client
	cbs    *cbs.Client
	log    logger.Logger
	*resource.AccountGetter
}
