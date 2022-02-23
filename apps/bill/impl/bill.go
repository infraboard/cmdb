package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"

	"github.com/infraboard/cmdb/apps/bill"
)

func (s *service) SyncBill(ctx context.Context, req *bill.Bill) (
	*bill.Bill, error) {
	stmt, err := s.db.Prepare(insertBillSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	y, m := req.YearMonth()

	_, err = stmt.Exec(
		req.Vendor, y, m, req.OwnerId, req.OwnerName, req.ProductType, req.ProductCode, req.ProductDetail,
		req.PayMode, req.OrderId, req.InstanceId, req.InstanceName, req.PublicIp, req.PrivateIp, req.InstanceConfig,
		req.RegionCode, req.RegionName, req.SalePrice, req.SaveCost, req.RealCost, req.CreditPay, req.VoucherPay,
		req.CashPay, req.StoredcardPay, req.OutstandingAmount,
	)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *service) QueryBill(ctx context.Context, req *bill.QueryBillRequest) (
	*bill.BillSet, error) {
	return nil, nil
}

func (s *service) DeleteBill(ctx context.Context, req *bill.DeleteBillRequest) (
	*bill.BillSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate delete bill error, %s", err)
	}

	set := bill.NewBillSet()
	builder := sqlbuilder.NewQuery(deleteBillSQL).Where("account_id", req.AccountId)
	if req.Month != "" {
		builder.Where("month", req.Month)
	}
	if req.TaskId != "" {
		builder.Where("task_id", req.TaskId)
	}

	deleteSQL, args := builder.BuildQuery()
	s.log.Debugf("sql: %s, args: %v", deleteSQL, args)

	stmt, err := s.db.Prepare(deleteSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	set.Total, _ = ret.RowsAffected()

	return set, nil
}
