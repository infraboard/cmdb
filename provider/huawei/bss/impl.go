package bss

import (
	v2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBssOperator(client *v2.BssClient) *BssOperator {
	return &BssOperator{
		client: client,
		log:    zap.L().Named("Huawei BSS"),
	}
}

type BssOperator struct {
	client *v2.BssClient
	log    logger.Logger
}
