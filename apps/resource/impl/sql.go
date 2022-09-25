package impl

const (
	sqlInsertResource = `INSERT INTO resource (
		id,resource_type,vendor,region,zone,create_at,expire_at,category,type,
		name,description,status,update_at,sync_at,owner,public_ip,
		private_ip,pay_type,describe_hash,resource_hash,credential_id,domain,
		namespace,env,usage_mode
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	sqlUpdateResource = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,owner=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?,
		credential_id=?,namespace=?,env=?,usage_mode=?
	WHERE id = ?`
	sqlDeleteResource = `DELETE FROM resource WHERE id = ?;`

	sqlQueryResource = `SELECT r.* FROM resource r %s JOIN resource_tag t ON r.id = t.resource_id`
	sqlCountResource = `SELECT COUNT(DISTINCT r.id) FROM resource r %s JOIN resource_tag t ON r.id = t.resource_id`

	sqlQueryResourceTag  = `SELECT t_key,t_value,description,resource_id,weight,type FROM resource_tag`
	sqlDeleteResourceTag = `
		DELETE 
		FROM
			resource_tag 
		WHERE
			resource_id =? 
			AND t_key =? 
			AND t_value =?;
	`
	sqlInsertOrUpdateResourceTag = `
		INSERT INTO resource_tag ( type, t_key, t_value, description, resource_id, weight, create_at)
		VALUES
			( ?,?,?,?,?,?,? ) 
			ON DUPLICATE KEY UPDATE description =
		IF
			( type != 1,?, description ),
			weight =
		IF
			( type != 1,?, weight );
	`
)
