package impl

const (
	insertHostSQL = `
	INSERT INTO resource_host  (
		resource_id,gpu_spec,os_type,os_name,
		image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?);
	`

	updateHostSQL = `
	UPDATE resource_host  SET 
		gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?
	`

	queryHostSQL = `
	SELECT
	r.*,
	h.*
	FROM
	resource AS r
	LEFT JOIN resource_host  h ON r.id = h.resource_id
	LEFT JOIN resource_tag t ON r.id = t.resource_id`
	countHostSQL = `SELECT
	COUNT(DISTINCT r.id)
	FROM
	resource AS r
	LEFT JOIN resource_host  h ON r.id = h.resource_id
	LEFT JOIN resource_tag t ON r.id = t.resource_id
	`

	deleteHostSQL = `DELETE FROM resource_host  WHERE resource_id = ?;`
)
