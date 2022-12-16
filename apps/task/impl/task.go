package impl

import (
	"context"
	"time"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/cmdb/apps/task"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
)

// 任务回调
type SyncTaskCallback func(*task.Task)

// 通过回调更新任务状态
func (s *service) SyncTaskCallback(t *task.Task) {
	err := s.update(context.Background(), t)
	if err != nil {
		s.log.Error(err)
	}
}

func (s *service) CreatTask(ctx context.Context, req *task.CreateTaskRequst) (
	*task.Task, error) {
	t, err := task.NewTaskFromReq(req)
	if err != nil {
		return nil, err
	}

	secretIns, err := s.secret.DescribeSecret(ctx, secret.NewDescribeSecretRequest(req.CredentialId))
	if err != nil {
		return nil, err
	}

	secret := secretIns.Spec
	t.UpdateSecretDesc(secret.ShortDesc())

	// 如果不是vsphere 需要检查region
	if !(secret.Vendor.Equal(resource.VENDOR_VSPHERE) || req.ResourceType.IsIn(resource.TYPE_BILL)) {
		if req.Region == "" {
			return nil, exception.NewBadRequest("region required")
		}
		if !secret.IsAllowRegion(req.Region) {
			return nil, exception.NewBadRequest("this secret not allow sync region %s", req.Region)
		}
	}

	// 解密secret
	err = secret.DecryptAPISecret(conf.C().App.EncryptKey)
	if err != nil {
		s.log.Warnf("decrypt api secret error, %s", err)
	}

	// 资源同步
	syncCtx, _ := context.WithTimeout(context.Background(), time.Minute*30)
	switch req.ResourceType {
	case resource.TYPE_HOST:
		go s.syncHost(syncCtx, secretIns, t, s.SyncTaskCallback)
	case resource.TYPE_BILL:
		go s.syncBill(syncCtx, secretIns, t, s.SyncTaskCallback)
	case resource.TYPE_RDS:
		go s.syncRds(syncCtx, secretIns, t, s.SyncTaskCallback)
	}

	// 记录任务
	if err = s.insert(ctx, t); err != nil {
		return nil, err
	}

	return t, nil
}

func (s *service) QueryTask(ctx context.Context, req *task.QueryTaskRequest) (*task.TaskSet, error) {
	query := sqlbuilder.NewQuery(queryTaskSQL)

	querySQL, args := query.Order("start_at").Desc().Limit(req.OffSet(), uint(req.Page.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query task error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := task.NewTaskSet()
	for rows.Next() {
		ins := task.NewDefaultTask()
		err := rows.Scan(
			&ins.Id, &ins.Data.Region, &ins.Data.ResourceType, &ins.Data.CredentialId, &ins.CredentialDescription,
			&ins.Data.Timeout, &ins.Status.Stage, &ins.Status.Message, &ins.Status.StartAt, &ins.Status.EndAt,
			&ins.Status.TotalSucceed, &ins.Status.TotalFailed, &ins.Data.Domain, &ins.Data.Namespace,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query task error, %s", err.Error())
		}
		set.Add(ins)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	return set, nil
}

func (s *service) DescribeTask(ctx context.Context, req *task.DescribeTaskRequest) (
	*task.Task, error) {
	query := sqlbuilder.NewQuery(queryTaskSQL)
	query.Where("id = ?", req.Id)

	querySQL, args := query.BuildQuery()
	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, err
	}
	defer queryStmt.Close()

	ins := task.NewDefaultTask()
	err = queryStmt.QueryRow(args...).Scan(
		&ins.Id, &ins.Data.Region, &ins.Data.ResourceType, &ins.Data.CredentialId, &ins.CredentialDescription,
		&ins.Data.Timeout, &ins.Status.Stage, &ins.Status.Message, &ins.Status.StartAt, &ins.Status.EndAt,
		&ins.Status.TotalSucceed, &ins.Status.TotalFailed, &ins.Data.Domain, &ins.Data.Namespace,
	)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) QueryTaskRecord(ctx context.Context, req *task.QueryTaskRecordRequest) (
	*task.RecordSet, error) {
	query := sqlbuilder.NewQuery(queryTaskRecordSQL)
	query.Where("task_id = ?", req.TaskId)

	querySQL, args := query.BuildQuery()
	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, err
	}
	defer queryStmt.Close()

	set := task.NewRecordSet()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rc := task.NewDefaultTaskRecord()
		rows.Scan(
			&rc.InstanceId, &rc.Name, &rc.IsSuccess, &rc.Message,
			&rc.TaskId, &rc.CreateAt,
		)
		set.Add(rc)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return set, nil
}
