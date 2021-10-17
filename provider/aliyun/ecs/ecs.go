package ecs

import "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

func NewEcsOperater(client *ecs.Client) *EcsOperater {
	return &EcsOperater{
		client: client,
	}
}

type EcsOperater struct {
	client *ecs.Client
}
