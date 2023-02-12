package impl

import "github.com/infraboard/cmdb/apps/resource"

func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Items: []*Resource{},
	}
}

func (s *ResourceSet) ResourceSet() *resource.ResourceSet {
	set := resource.NewResourceSet()
	set.Total = s.Total
	for i := range s.Items {
		set.Add(s.Items[i].Resource())
	}
	return set
}

type ResourceSet struct {
	Total int64
	Items []*Resource
}

func NewResource(res *resource.Resource) *Resource {
	rid := &ResourceId{ResourceId: res.Meta.Id}
	temp := &Resource{
		ResourceMeta: &ResourceMeta{
			Meta:        res.Meta,
			ContentHash: res.ContentHash,
		},
		ResourceSpec: &ResourceSpec{
			ResourceId: rid,
			Spec:       res.Spec,
		},
		ResourceCost: &ResourceCost{
			ResourceId: rid,
			Cost:       res.Cost,
		},
		ResourceStatus: &ResourceStatus{
			ResourceId: rid,
			Status:     res.Status,
		},
		Tags: []*ResourceTag{},
	}
	for i := range res.Spec.Tags {
		temp.Tags = append(temp.Tags, &ResourceTag{
			ResourceId: rid,
			Tag:        res.Spec.Tags[i],
		})
	}
	return temp
}

type Resource struct {
	*ResourceMeta
	*ResourceSpec
	*ResourceCost
	*ResourceStatus
	Tags []*ResourceTag
}

func (r *Resource) Resource() *resource.Resource {
	ins := resource.NewDefaultResource(r.ResourceType)
	ins.Meta = r.Meta
	ins.Spec = r.Spec
	ins.Cost = r.Cost
	ins.Status = r.Status
	ins.ContentHash = r.ContentHash

	for i := range r.Tags {
		item := r.Tags[i]
		ins.Spec.Tags = append(ins.Spec.Tags, item.Tag)
	}
	return nil
}

type ResourceMeta struct {
	*resource.Meta
	*resource.ContentHash
}

const (
	RESOURCE_META_TABLE_NAME = "resource_meta"
)

func (s *ResourceMeta) TableName() string {
	return RESOURCE_META_TABLE_NAME
}

type ResourceId struct {
	ResourceId string `json:"resource_id"`
}

type ResourceSpec struct {
	*ResourceId
	*resource.Spec
}

func (s *ResourceSpec) TableName() string {
	return "resource_spec"
}

type ResourceCost struct {
	*ResourceId
	*resource.Cost
}

func (s *ResourceCost) TableName() string {
	return "resource_cost"
}

type ResourceStatus struct {
	*ResourceId
	*resource.Status
}

func (s *ResourceStatus) TableName() string {
	return "resource_status"
}

type ResourceTag struct {
	*ResourceId
	*resource.Tag

	UpdateAt int64 `json:"update_at"`
}

func (s *ResourceTag) TableName() string {
	return "resource_tag"
}

type ResourceRelation struct {
	*ResourceId
	TargetId string
}

func (s *ResourceRelation) TableName() string {
	return "resource_relation"
}
