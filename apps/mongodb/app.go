package mongodb

import resource "github.com/infraboard/cmdb/apps/resource"

const (
	AppName = "mongodb"
)

func NewDefaultMongoDB() *MongoDB {
	return &MongoDB{
		Base: &resource.Base{
			ResourceType: resource.TYPE_MONGODB,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func NewMongoDBSet() *MongoDBSet {
	return &MongoDBSet{
		Items: []*MongoDB{},
	}
}

func (s *MongoDBSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *MongoDBSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*MongoDB))
	}
}

func (s *MongoDBSet) AddSet(set *MongoDBSet) {
	s.Items = append(s.Items, set.Items...)
}

func (s *MongoDBSet) Length() int64 {
	return int64(len(s.Items))
}
