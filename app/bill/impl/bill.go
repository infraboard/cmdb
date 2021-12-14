package impl

import (
	"context"

	"github.com/infraboard/cmdb/app/bill"
)

func (s *service) SaveBill(ctx context.Context, req *bill.Bill) (
	*bill.Bill, error) {
	return nil, nil
}

func (s *service) QueryBill(ctx context.Context, req *bill.QueryBillRequest) (
	*bill.BillSet, error) {
	return nil, nil
}
