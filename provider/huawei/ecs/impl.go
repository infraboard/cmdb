package ecs

import (
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	eip "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2"
	evs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperator(client *ecs.EcsClient, evs *evs.EvsClient, eip *eip.EipClient) *EcsOperator {
	return &EcsOperator{
		client:        client,
		evs:           evs,
		eip:           eip,
		log:           zap.L().Named("hw.ecs"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type EcsOperator struct {
	client *ecs.EcsClient
	evs    *evs.EvsClient
	eip    *eip.EipClient
	log    logger.Logger
	*resource.AccountGetter
}
