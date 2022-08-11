package bss

import (
	v2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"

	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBssOperator(client *v2.BssClient) *BssOperator {
	return &BssOperator{
		client: client,
		log:    zap.L().Named("hw.bss"),
		tb:     tokenbucket.NewBucketWithRate(10, 1),
	}
}

type BssOperator struct {
	client *v2.BssClient
	log    logger.Logger
	tb     *tokenbucket.Bucket
}
