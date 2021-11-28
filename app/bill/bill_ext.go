package bill

func NewDefaultBill() *Bill {
	return &Bill{}
}

func NewBillSet() *BillSet {
	return &BillSet{
		Items: []*Bill{},
	}
}

func (s *BillSet) Add(item *Bill) {
	s.Items = append(s.Items, item)
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewBillSet(),
	}
}

type PagerResult struct {
	Data    *BillSet
	Err     error
	HasNext bool
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}
