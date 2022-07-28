package cms

import (
	cms "github.com/alibabacloud-go/cms-20190101/v7/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCmsOperator(client *cms.Client) *CmsOperator {
	return &CmsOperator{
		client: client,
		log:    zap.L().Named("ALI CMS"),
	}
}

type CmsOperator struct {
	client *cms.Client
	log    logger.Logger
}
