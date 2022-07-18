package voice

import (
	vms "github.com/alibabacloud-go/dyvmsapi-20170525/v2/client"
	"github.com/infraboard/mcube/logger"
)

// 参考: https://next.api.aliyun.com/api/Dyvmsapi/2017-05-25/SingleCallByVoice?params={}
type VmsOperator struct {
	client *vms.Client
	log    logger.Logger
}
