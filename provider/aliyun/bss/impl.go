package bss

import (
	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBssOperator(client *bssopenapi.Client) *BssOperator {
	return &BssOperator{
		client: client,
		log:    zap.L().Named("ali.bss"),
	}
}

type BssOperator struct {
	client *bssopenapi.Client
	log    logger.Logger
}
