package bill

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "bill"
)

var (
	validate = validator.New()
)

func NewDefaultBill() *Bill {
	return &Bill{
		Cost: &Cost{},
	}
}

func NewBillSet() *BillSet {
	return &BillSet{
		Sum:   &Cost{},
		Items: []*Bill{},
	}
}

func (s *BillSet) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}

	return
}

func (s *BillSet) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Bill))
	}
}

func (s *BillSet) Length() int64 {
	return int64(len(s.Items))
}

func (b *Bill) YearMonth() (int, int) {
	if b.Month == "" {
		return 0, 0
	}

	ym := strings.Split(b.Month, "-")
	if len(ym) > 1 {
		y, _ := strconv.Atoi(ym[0])
		m, _ := strconv.Atoi(ym[1])
		return y, m
	}

	return 0, 0
}

func (b *Bill) ShortDesc() string {
	return fmt.Sprintf("%s %s", b.InstanceId, b.InstanceName)
}

func (b *Bill) ToJsonString() string {
	return tea.Prettify(b)
}

func NewDeleteBillRequest(taskId string) *DeleteBillRequest {
	return &DeleteBillRequest{
		TaskId: taskId,
	}
}

func (req *DeleteBillRequest) Validate() error {
	return validate.Struct(req)
}

func NewConfirmBillRequest(taskId string) *ConfirmBillRequest {
	return &ConfirmBillRequest{
		TaskId: taskId,
	}
}
