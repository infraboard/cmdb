package clb

import (
	"github.com/infraboard/mcube/logger"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
)

type DomainOperator struct {
	client *clb.Client
	log    logger.Logger
}
