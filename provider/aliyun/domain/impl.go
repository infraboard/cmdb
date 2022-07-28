package domain

import (
	domain "github.com/alibabacloud-go/domain-20180129/v3/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewDomainOperator(client *domain.Client) *DomainOperator {
	return &DomainOperator{
		client: client,
		log:    zap.L().Named("ali.domain"),
	}
}

type DomainOperator struct {
	client *domain.Client
	log    logger.Logger
}
