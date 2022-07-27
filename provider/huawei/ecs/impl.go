package ecs

import (
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	evs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperator(client *ecs.EcsClient, evs *evs.EvsClient) *EcsOperator {
	return &EcsOperator{
		client:        client,
		evs:           evs,
		log:           zap.L().Named("hw.ecs"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type EcsOperator struct {
	client *ecs.EcsClient
	evs    *evs.EvsClient
	log    logger.Logger
	*resource.AccountGetter
}
