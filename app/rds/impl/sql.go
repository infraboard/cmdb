package impl

const (
	insertRdsSQL = `INSERT INTO rds (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	updateRdsSQL = `UPDATE rds SET 
		cpu=?,memory=?,gpu_amount=?,gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?`

	queryRdsSQL  = `SELECT * FROM resource as r LEFT JOIN rds i ON i.id=h.resource_id`
	deleteRdsSQL = `DELETE FROM rds WHERE resource_id = ?;`
)
