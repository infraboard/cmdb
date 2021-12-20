package bill

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (b *Bill) YearMonth() (int, int) {
	if b.Month == "" {
		return 0, 0
	}

	ym := strings.Split(b.Month, "-")
	if len(ym) == 2 {
		y, _ := strconv.Atoi(ym[0])
		m, _ := strconv.Atoi(ym[1])
		return y, m
	}

	return 0, 0
}

func (b *Bill) ShortDesc() string {
	return fmt.Sprintf("%s %s", b.InstanceId, b.InstanceName)
}
