package bill

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "Bill"
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
		Items: []*Bill{},
	}
}

func (s *BillSet) Add(items ...*Bill) {
	s.Items = append(s.Items, items...)
}

// 分页迭代器
type Pager interface {
	Next() bool
	Scan(context.Context, *BillSet) error
}

func (b *Bill) YearMonth() (int, int) {
	if b.Month == "" {
		return 0, 0
	}

	ym := strings.Split(b.Month, "-")
	if len(ym) > 2 {
		y, _ := strconv.Atoi(ym[0])
		m, _ := strconv.Atoi(ym[1])
		return y, m
	}

	return 0, 0
}

func (b *Bill) ShortDesc() string {
	return fmt.Sprintf("%s %s", b.InstanceId, b.InstanceName)
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
