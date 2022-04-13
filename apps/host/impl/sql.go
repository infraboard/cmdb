package impl

const (
	insertHostSQL = `INSERT INTO resource_host  (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`

	updateHostSQL = `UPDATE resource_host  SET 
		cpu=?,memory=?,gpu_amount=?,gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?`

	queryHostSQL = `SELECT
	r.*,
	h.*
	FROM
	resource AS r
	LEFT JOIN resource_host  h ON r.id = h.resource_id
	LEFT JOIN resource_tag t ON r.id = t.resource_id`

	deleteHostSQL = `DELETE FROM resource_host  WHERE resource_id = ?;`
)
