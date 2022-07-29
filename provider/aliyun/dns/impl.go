package dns

import (
	domain "github.com/alibabacloud-go/domain-20180129/v3/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewDomainOperator(client *domain.Client) *DnsOperator {
	return &DnsOperator{
		client: client,
		log:    zap.L().Named("ali.domain"),
	}
}

type DnsOperator struct {
	client *domain.Client
	log    logger.Logger
}
