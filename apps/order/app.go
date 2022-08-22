package order

import "github.com/alibabacloud-go/tea/tea"

const (
	AppName = "order"
)

func NewDefaultOrder() *Order {
	return &Order{
		Cost: &Cost{},
	}
}

func NewOrderSet() *OrderSet {
	return &OrderSet{
		Items: []*Order{},
	}
}

func (s *OrderSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}

	return
}

func (s *OrderSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Order))
	}
}

func (s *OrderSet) GetOrderById(orderId string) *Order {
	for i := range s.Items {
		if s.Items[i].Id == orderId {
			return s.Items[i]
		}
	}

	return nil
}

func (s *OrderSet) Length() int64 {
	return int64(len(s.Items))
}

func (o *Order) ToJsonString() string {
	return tea.Prettify(o)
}
