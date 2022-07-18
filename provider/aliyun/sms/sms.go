package sms

import (
	sms "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/infraboard/mcube/logger"
)

// 参考: https://next.api.aliyun.com/api/Dysmsapi/2017-05-25/SendSms?lang=GO&params={}
type SmsOperator struct {
	client *sms.Client
	log    logger.Logger
}
