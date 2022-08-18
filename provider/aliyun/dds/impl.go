package dds

import (
	dds "github.com/alibabacloud-go/dds-20151201/v3/client"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewOperator(client *dds.Client) *Operator {
	return &Operator{
		client: client,
		log:    zap.L().Named("ali.mongodb"),
	}
}

type Operator struct {
	client *dds.Client
	log    logger.Logger
}
