package impl

const (
	insertSecretSQL = `INSERT INTO credential (
		id,create_at,description,vendor,address,allow_regions,
		crendential_type,api_key,api_credential,request_rate,
		domain,namespace
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);`

	querySecretSQL = `SELECT * FROM credential`

	deleteSecretSQL = `DELETE FROM credential WHERE id = ?;`
)
