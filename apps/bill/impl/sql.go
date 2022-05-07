package impl

const (
	insertBillSQL = `INSERT INTO bill_raw (
		vendor,year,month,owner_id,owner_name,product_type,product_code,product_detail,
		pay_mode,order_id,resource_id,resource_name,public_ip,private_ip,instance_config,
		region_code,region_name,sale_price,save_cost,real_cost,credit_pay,voucher_pay,
		cash_pay,storedcard_pay,outstanding_amount,is_merged,task_id
	) VALUES (?,?,?,?,?,?,?,?,?,?,?, ?,?,?,?,?,?,?,?,?,?,?,?,?,?,0,?);`

	queryBillSQL  = `SELECT * FROM bill_raw`
	deleteBillSQL = `DELETE FROM bill_raw WHERE task_id=?`
)

const (
	mergeBillSQL = `INSERT INTO bill (
		vendor,year,month,owner_id,owner_name,product_type,product_code,product_detail,
		pay_mode,order_id,resource_id,resource_name,public_ip,private_ip,instance_config,
		region_code,region_name,sale_price,save_cost,real_cost,credit_pay,voucher_pay,
		cash_pay,storedcard_pay,outstanding_amount,task_id,is_merged
	) SELECT 
		vendor,year,month,owner_id,owner_name,product_type,product_code,product_detail,
		pay_mode,order_id,resource_id,resource_name,public_ip,private_ip,instance_config,
		region_code,region_name,sum(sale_price),sum(save_cost),sum(real_cost),sum(credit_pay),sum(voucher_pay),
		sum(cash_pay),sum(storedcard_pay),sum(outstanding_amount),task_id,1
	 FROM bill WHERE is_merged=0 AND task_id=? GROUP BY resource_id`
)
