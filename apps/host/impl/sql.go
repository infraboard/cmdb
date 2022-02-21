package impl

const (
	insertHostSQL = `INSERT INTO host (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`

	updateHostSQL = `UPDATE host SET 
		cpu=?,memory=?,gpu_amount=?,gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?`

	queryHostSQL = `SELECT
	r.*,
	h.*,
	IFNULL(GROUP_CONCAT(t.key),'') tag_keys,
	IFNULL(GROUP_CONCAT(t.value), '') tag_values,
	IFNULL(GROUP_CONCAT(t.describe), '') tag_describe 
	FROM
	resource AS r
	LEFT JOIN host h ON r.id = h.resource_id
	LEFT JOIN tag t ON r.id = t.resource_id`

	deleteHostSQL = `DELETE FROM host WHERE resource_id = ?;`
)
