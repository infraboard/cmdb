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
		req.RegionCode, req.RegionName, req.Cost.SalePrice, req.Cost.SaveCost, req.Cost.RealCost, req.Cost.CreditPay,
		req.Cost.VoucherPay, req.Cost.CashPay, req.Cost.StoredcardPay, req.Cost.OutstandingAmount, req.TaskId,
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

// 确认账单, 确认当前月的账单以TaskId同步的准, 之前其他Task同步的账单作废需要删除
func (s *service) ConfirmBill(ctx context.Context, req *bill.ConfirmBillRequest) (
	*bill.BillSet, error) {
	b := sqlbuilder.NewQuery(deleteBillSQL)
	b.Where("year = ?")

	// 清理该月之前同步的账单数据

	// 生产实例按照月的聚合的数据

	return nil, nil
}

func (s *service) DeleteBill(ctx context.Context, req *bill.DeleteBillRequest) (
	*bill.BillSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate delete bill error, %s", err)
	}

	set := bill.NewBillSet()
	stmt, err := s.db.Prepare(deleteBillSQL)
	if err != nil {
		return set, err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(req.TaskId)
	if err != nil {
		return set, err
	}
	set.Total, _ = ret.RowsAffected()

	return set, nil
}
