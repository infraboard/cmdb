package dns

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func NewDnsOperator(client *dnspod.Client) *DnsOperator {
	return &DnsOperator{
		client: client,
		log:    zap.L().Named("tx.dns"),
	}
}

type DnsOperator struct {
	client *dnspod.Client
	log    logger.Logger
}
