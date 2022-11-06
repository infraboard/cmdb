package impl

import "github.com/infraboard/cmdb/apps/resource"

type Resource struct {
	*resource.Meta
	*ResourceSpec
	*ResourceCost
	*ResourceStatus
}

type ResourceSpec struct {
	ResourceId string
	*resource.Spec
}

func (s *ResourceSpec) TableName() string {
	return "resource_spec"
}

type ResourceCost struct {
	ResourceId string
	*resource.Cost
}

func (s *ResourceCost) TableName() string {
	return "resource_cost"
}

type ResourceStatus struct {
	ResourceId string
	*resource.Status
}

func (s *ResourceStatus) TableName() string {
	return "resource_status"
}

type ResourceTag struct {
	ResourceId string
	*resource.Tag
}

func (s *ResourceTag) TableName() string {
	return "resource_tag"
}

type ResourceRelation struct {
	SourceId string
	TargetId string
}

func (s *ResourceRelation) TableName() string {
	return "resource_relation"
}
