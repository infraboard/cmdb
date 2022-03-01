package impl

const (
	SQLInsertResource = `INSERT INTO resource (
		id,resource_type,vendor,region,zone,create_at,expire_at,category,type,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash,secret_id,domain,
		namespace,env,usage_mode
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	SQLUpdateResource = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?,
		secret_id=?,namespace=?,env=?,usage_mode=?
	WHERE id = ?`

	SQLDeleteResource = `DELETE FROM resource WHERE id = ?;`
	SQLQueryResource  = `
	SELECT
		r.*,
		IFNULL( GROUP_CONCAT( t.t_key ), '' ) tag_keys,
		IFNULL( GROUP_CONCAT( t.t_value ), '' ) tag_values,
		IFNULL( GROUP_CONCAT( t.description ), '' ) tag_describes,
		IFNULL( GROUP_CONCAT( t.weight ), '' ) tag_weights,
		IFNULL( GROUP_CONCAT( t.type ), '' ) tag_types 
	FROM
		resource r
		LEFT JOIN tag t ON r.id = t.resource_id`

	SQLDeleteResourceTag = `
		DELETE 
		FROM
			tag 
		WHERE
			resource_id =? 
			AND t_key =? 
			AND t_value =?;
	`

	SQLInsertOrUpdateResourceTag = `
		INSERT INTO tag ( type, t_key, t_value, description, resource_id, weight )
		VALUES
			( ?,?,?,?,?,? ) 
			ON DUPLICATE KEY UPDATE description =
		IF
			( type != 1,?, description ),
			weight =
		IF
			( type != 1,?, weight );
	`
)
