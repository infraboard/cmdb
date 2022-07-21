package domain

import (
	"github.com/infraboard/mcube/logger"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

type DomainOperator struct {
	client *dnspod.Client
	log    logger.Logger
}
