package impl

const (
	SQLInsertResource = `INSERT INTO resource (
		id,resource_type,vendor,region,zone,create_at,expire_at,category,type,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash,secret_id,domain,
		namespace,env,usage_mode
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	SQLUpdateResource = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?,
		secret_id=?,namespace=?,env=?,usage_mode=?
	WHERE id = ?`

	SQLQueryResource  = `SELECT * FROM resource`
	SQLDeleteResource = `DELETE FROM resource WHERE id = ?;`

	SQLDeleteResourceTag         = `DELETE FROM tag WHERE resource_id = ?;`
	SQLInsertOrUpdateResourceTag = `INSERT INTO tag (
		type,t_key,t_value,description,resource_id,weight
	) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE 
		description=IF(type!=1,?,description),
		weight=IF(type!=1,?,weight);`
	SQLUpdateResourceTag = `UPDATE tag SET
		description=?,weight=? 
	WHERE id = ?
	`
)
