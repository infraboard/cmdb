package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/app/task"
)

const (
	insertTaskSQL = `INSERT INTO task (
		id,region,resource_type,secret_id,secret_desc,timeout,status,
		message,start_at,end_at,total_succeed,total_failed
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);`

	updateTaskSQL = `UPDATE task SET status=?,message=?,end_at=?,
	total_succeed=?,total_failed=? WHERE id = ?`

	queryTaskSQL = `SELECT * FROM task`

	updateOrInsertDetailSQL = `INSERT INTO task_detail (
		instance_id,instance_name,is_success,message,task_id) 
		VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE 
		instance_name=?,is_success=?,message=?,task_id=?`
)

func (s *service) insert(ctx context.Context, t *task.Task) error {
	stmt, err := s.db.Prepare(insertTaskSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		t.Id, t.Region, t.ResourceType, t.SecretId, t.SecretDescription, t.Timeout,
		t.Status, t.Message, t.StartAt, t.EndAt, t.TotalSucceed, t.TotalFailed,
	)
	if err != nil {
		return fmt.Errorf("save task info error, %s", err)
	}
	return nil
}

func (s *service) update(ctx context.Context, t *task.Task) error {
	stmt, err := s.db.Prepare(updateTaskSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		t.Status, t.Message, t.EndAt, t.TotalSucceed, t.TotalFailed, t.Id,
	)
	if err != nil {
		return fmt.Errorf("update task info error, %s", err)
	}

	return nil
}

func (s *service) insertOrUpdateDetail(ctx context.Context, taskId string, detail *task.Detail) error {
	stmt, err := s.db.Prepare(updateOrInsertDetailSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		detail.InstanceId, detail.Name, detail.IsSuccess, detail.Message, taskId,
		detail.Name, detail.IsSuccess, detail.Message, taskId,
	)
	if err != nil {
		return fmt.Errorf("insert or update task %s detail info error, %s", taskId, err)
	}

	return nil
}
