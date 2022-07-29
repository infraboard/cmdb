package dns

import (
	dns "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewDnsOperator(client *dns.DnsClient) *DnsOperator {
	return &DnsOperator{
		client: client,
		log:    zap.L().Named("hw.dns"),
	}
}

type DnsOperator struct {
	client *dns.DnsClient
	log    logger.Logger
}
