package lb

import (
	slb "github.com/alibabacloud-go/slb-20140515/v3/client"
	"github.com/infraboard/mcube/logger"
)

type LBOperator struct {
	client *slb.Client
	log    logger.Logger
}
