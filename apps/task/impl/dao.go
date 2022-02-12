package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/task"
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

func (s *service) insertTaskDetail(ctx context.Context, detail *task.Record) error {
	stmt, err := s.db.Prepare(insertTaskRecordSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		detail.InstanceId, detail.Name, detail.IsSuccess, detail.Message,
		detail.TaskId, detail.CreateAt,
	)
	if err != nil {
		return fmt.Errorf("insert or update task %s detail info error, %s", detail.TaskId, err)
	}

	return nil
}
