package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewOssOperator(client *oss.Client) *OssOperator {
	return &OssOperator{
		client: client,
		log:    zap.L().Named("ALI Oss"),
	}
}

type OssOperator struct {
	client *oss.Client
	log    logger.Logger
}
