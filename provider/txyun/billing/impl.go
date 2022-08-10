package billing

import (
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBillingoperator(client *billing.Client) *BillOperator {
	return &BillOperator{
		client: client,
		log:    zap.L().Named("tx.billing"),
	}
}

type BillOperator struct {
	client *billing.Client
	log    logger.Logger
}
