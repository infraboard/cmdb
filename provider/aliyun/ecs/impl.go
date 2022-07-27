package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperator(client *ecs.Client) *EcsOperator {
	return &EcsOperator{
		client:        client,
		log:           zap.L().Named("ALI ECS"),
		AccountGetter: &resource.AccountGetter{},
	}
}

// https://next.api.aliyun.com/api/Ecs/2014-05-26/CreateInstance?lang=GO&params={}
type EcsOperator struct {
	client *ecs.Client
	log    logger.Logger
	*resource.AccountGetter
}
