package impl

import (
	"context"

	"github.com/infraboard/cmdb/apps/bill"
)

const (
	insertBillSQL = `INSERT INTO bill (
		vendor,year,month,owner_id,owner_name,product_type,product_code,product_detail,
		pay_mode,order_id,instance_id,instance_name,public_ip,private_ip,instance_config,
		region_code,region_name,sale_price,save_cost,real_cost,credit_pay,voucher_pay,
		cash_pay,storedcard_pay,outstanding_amount
	) VALUES (?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?);`

	queryBillSQL = `SELECT * FROM bill`
)

func (s *service) SaveBill(ctx context.Context, req *bill.Bill) (
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
