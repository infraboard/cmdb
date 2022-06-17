package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/cmdb/apps/task"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/cmdb/provider/huawei"
	"github.com/infraboard/cmdb/provider/txyun"
)

func (s *service) syncRds(ctx context.Context, secretIns *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager pager.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		t.Completed()
		cb(t)
	}()

	secret := secretIns.Data
	req := provider.NewQueryRdsRequestWithRate(secret.RequestRate)
	switch secret.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun rds ...")
		op, err := aliyun.NewOperator(secret.ApiKey, secret.ApiSecret, t.Data.Region)
		if err != nil {
			t.Failed(err.Error())
			return
		}
		pager = op.RdsOperator().QueryRds(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun rds ...")
		op, err := txyun.NewOperator(secret.ApiKey, secret.ApiSecret, t.Data.Region)
		if err != nil {
			t.Failed(err.Error())
			return
		}
		pager = op.RdsOperator().QueryRds(req)
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun rds ...")
		op, err := huawei.NewOperator(secret.ApiKey, secret.ApiSecret, t.Data.Region)
		if err != nil {
			t.Failed(err.Error())
			return
		}
		pager = op.RdsOperator().QueryRds(req)
	default:
		t.Failed(fmt.Sprintf("unsuport bill syncing vendor %s", secret.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		for pager.Next() {
			set := rds.NewSet()
			if err := pager.Scan(ctx, set); err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", err))
				return
			}
			// 调用rds服务保持数据
			for i := range set.Items {
				target := set.Items[i]
				target.Base.SecretId = secretIns.Id
				s.SaveOrUpdateRds(ctx, target, t)
			}
		}
	}
}

// Rds数据入库
func (s *service) SaveOrUpdateRds(ctx context.Context, ins *rds.Rds, t *task.Task) {
	b, err := s.rds.SyncRDS(ctx, ins)

	var detail *task.Record
	if err != nil {
		s.log.Warnf("save rds error, %s", err)
		detail = task.NewSyncFailedRecord(t.Id, ins.Base.Id, ins.Information.Name, err.Error())
	} else {
		s.log.Debugf("save rds %s to db", b.ShortDesc())
		detail = task.NewSyncSucceedRecord(t.Id, ins.Base.Id, ins.Information.Name)
	}

	t.AddDetail(detail)
	if err := s.insertTaskDetail(ctx, detail); err != nil {
		s.log.Errorf("update detail error, %s", err)
	}
}
