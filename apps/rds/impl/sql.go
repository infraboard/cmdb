package impl

const (
	insertRdsSQL = `INSERT INTO rds (
		resource_id,engine_type,engine_version,instance_class,class_type,export_type,
		network_type,type,cpu,memory,db_max_quantity,account_max_quantity,max_connections,
		max_iops,collation,time_zone,storage_capacity,storage_type,security_ip_mode,
		security_ip_list,connection_mode,ip_type,lock_mode,lock_reason,deploy_mode,
		port,extra
	) VALUES (,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	updateRdsSQL = `UPDATE rds SET 
		cpu=?,memory=?,collation=?,time_zone=?,storage_type=?,storage_capacity=?,
		security_ip_mode=?,security_ip_list=?,connection_mode=?,ip_type=?,
		lock_mode=?,lock_reason=?,port=?,extra=?
	WHERE resource_id = ?`

	queryRdsSQL  = `SELECT * FROM resource as r LEFT JOIN rds i ON i.id=h.resource_id`
	deleteRdsSQL = `DELETE FROM rds WHERE resource_id = ?;`
)
