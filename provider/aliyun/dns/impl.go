package dns

import (
	alidns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	domain "github.com/alibabacloud-go/domain-20180129/v3/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewDomainOperator(client *domain.Client, dns *alidns.Client) *DnsOperator {
	return &DnsOperator{
		client: client,
		dns:    dns,
		log:    zap.L().Named("ali.domain"),
	}
}

type DnsOperator struct {
	dns    *alidns.Client
	client *domain.Client
	log    logger.Logger
}
