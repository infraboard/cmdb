package impl

import "github.com/infraboard/cmdb/apps/resource"

func BuildResourceBatch(set *resource.ResourceSet) *ResourceBatch {
	batch := NewResourceBatch()
	for i := range set.Items {
		item := set.Items[i]
		batch.AddMeta(item.Meta)
		batch.AddSpec(item.Meta.Id, item.Spec)
		batch.AddCost(item.Meta.Id, item.Cost)
		batch.AddStatus(item.Meta.Id, item.Status)
		batch.AddTags(item.Meta.Id, item.Tags)
		batch.AddRelations(item.Meta.Id, item.RelatedResources)
	}
	return batch
}

func NewResourceBatch() *ResourceBatch {
	return &ResourceBatch{
		Metas:     []*resource.Meta{},
		Specs:     []*ResourceSpec{},
		Costs:     []*ResourceCost{},
		Status:    []*ResourceStatus{},
		Tags:      []*ResourceTag{},
		Relations: []*ResourceRelation{},
	}
}

type ResourceBatch struct {
	Metas     []*resource.Meta
	Specs     []*ResourceSpec
	Costs     []*ResourceCost
	Status    []*ResourceStatus
	Tags      []*ResourceTag
	Relations []*ResourceRelation
}

func (b *ResourceBatch) Records() (items []any) {
	items = append(items, b.Metas, b.Specs, b.Costs, b.Status, b.Tags, b.Relations)
	return
}

func (b *ResourceBatch) AddMeta(m *resource.Meta) {
	b.Metas = append(b.Metas, m)
}

func (b *ResourceBatch) AddSpec(id string, spec *resource.Spec) {
	b.Specs = append(b.Specs, &ResourceSpec{
		ResourceId: id,
		Spec:       spec,
	})
}

func (b *ResourceBatch) AddCost(id string, cost *resource.Cost) {
	b.Costs = append(b.Costs, &ResourceCost{
		ResourceId: id,
		Cost:       cost,
	})
}

func (b *ResourceBatch) AddStatus(id string, status *resource.Status) {
	b.Status = append(b.Status, &ResourceStatus{
		ResourceId: id,
		Status:     status,
	})
}

func (b *ResourceBatch) AddTags(id string, tags []*resource.Tag) {
	for i := range tags {
		b.Tags = append(b.Tags, &ResourceTag{
			ResourceId: id,
			Tag:        tags[i],
		})
	}
}

func (b *ResourceBatch) AddRelations(id string, resources []*resource.Resource) {
	for i := range resources {
		b.Relations = append(b.Relations, &ResourceRelation{
			SourceId: id,
			TargetId: resources[i].Meta.Id,
		})
	}
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
