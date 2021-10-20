package ecs

import (
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperater(client *ecs.EcsClient) *EcsOperater {
	return &EcsOperater{
		client: client,
		log:    zap.L().Named("Huawei ECS"),
	}
}

type EcsOperater struct {
	client *ecs.EcsClient
	log    logger.Logger
}

func (o *EcsOperater) transferSet(list *[]model.ServerDetail) *host.HostSet {
	set := host.NewHostSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *EcsOperater) transferOne(ins model.ServerDetail) *host.Host {
	h := host.NewDefaultHost()
	return h
}
