package domain

import (
	domain "github.com/alibabacloud-go/domain-20180129/client"
	"github.com/infraboard/mcube/logger"
)

type DomainOperator struct {
	client *domain.Client
	log    logger.Logger
}
