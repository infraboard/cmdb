package impl

const (
	insertBillSQL = `INSERT INTO bill (
		vendor,year,month,owner_id,owner_name,product_type,product_code,product_detail,
		pay_mode,order_id,instance_id,instance_name,public_ip,private_ip,instance_config,
		region_code,region_name,sale_price,save_cost,real_cost,credit_pay,voucher_pay,
		cash_pay,storedcard_pay,outstanding_amount
	) VALUES (?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?);`

	queryBillSQL  = `SELECT * FROM bill`
	deleteBillSQL = `DELETE FROM bill WHERE task_id=?`
)
