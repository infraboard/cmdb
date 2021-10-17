package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/infraboard/cmdb/pkg/host"
)

func NewEcsOperater(client *ecs.Client) *EcsOperater {
	return &EcsOperater{
		client: client,
	}
}

type EcsOperater struct {
	client *ecs.Client
}

func (o *EcsOperater) transferSet(items []ecs.Instance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *EcsOperater) transferOne(ins ecs.Instance) *host.Host {
	return nil
}
