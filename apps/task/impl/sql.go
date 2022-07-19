package impl

const (
	insertTaskSQL = `
	INSERT INTO task (
		id,region,resource_type,credential_id,credential_desc,timeout,status,
		message,start_at,end_at,total_succeed,total_failed,
		domain,namespace
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`

	updateTaskSQL = `
	UPDATE task SET status=?,message=?,end_at=?,
	total_succeed=?,total_failed=? WHERE id = ?
	`

	queryTaskSQL = `SELECT * FROM task`

	insertTaskRecordSQL = `
	INSERT INTO task_record (
		instance_id,instance_name,is_success,message,task_id,create_at) 
		VALUES (?,?,?,?,?,?);
	`

	queryTaskRecordSQL = `SELECT * FROM task_record`
)
