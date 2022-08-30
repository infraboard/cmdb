package actiontrail

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	actiontrail "github.com/alibabacloud-go/actiontrail-20200706/v2/client"
)

func NewOperator(client *actiontrail.Client) *Operator {
	return &Operator{
		client: client,
		log:    zap.L().Named("ali.actiontrail"),
	}
}

type Operator struct {
	client *actiontrail.Client
	log    logger.Logger
}
