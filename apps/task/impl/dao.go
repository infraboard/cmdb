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
		t.Id, t.Data.Region, t.Data.ResourceType, t.Data.SecretId, t.SecretDescription, t.Data.Timeout,
		t.Status.Stage, t.Status.Message, t.Status.StartAt, t.Status.EndAt, t.Status.TotalSucceed, t.Status.TotalFailed,
		t.Data.Domain, t.Data.Namespace,
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

	status := t.Status

	_, err = stmt.Exec(
		status.Stage, status.Message, status.EndAt, status.TotalSucceed, status.TotalFailed, t.Id,
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
