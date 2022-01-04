package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/app/rds"
	"github.com/infraboard/cmdb/app/resource"
	"github.com/infraboard/cmdb/app/secret"
	"github.com/infraboard/cmdb/app/task"

	aliConn "github.com/infraboard/cmdb/provider/aliyun/connectivity"
	rdsOp "github.com/infraboard/cmdb/provider/aliyun/rds"
	hwConn "github.com/infraboard/cmdb/provider/huawei/connectivity"
	hwRdsOp "github.com/infraboard/cmdb/provider/huawei/rds"
	cdbOp "github.com/infraboard/cmdb/provider/txyun/cdb"
	txConn "github.com/infraboard/cmdb/provider/txyun/connectivity"
)

func (s *service) syncRds(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager rds.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		t.Completed()
		cb(t)
	}()

	switch secret.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun rds ...")
		client := aliConn.NewAliCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		bc, err := client.RdsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}

		operater := rdsOp.NewRdsOperater(bc)
		req := rdsOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun rds ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		operater := cdbOp.NewCDBOperater(client.CDBClient())
		pager = operater.PageQuery()
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun rds ...")
		client := hwConn.NewHuaweiCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		ec, err := client.RdsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := hwRdsOp.NewEcsOperater(ec)
		pager = operater.PageQuery()
	default:
		t.Failed(fmt.Sprintf("unsuport bill syncing vendor %s", secret.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		hasNext := true
		for hasNext {
			p := pager.Next()
			hasNext = p.HasNext

			if p.Err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", p.Err))
				return
			}

			// 调用rds服务保持数据
			for i := range p.Data.Items {
				target := p.Data.Items[i]
				target.Base.SecretId = secret.Id
				s.SaveOrUpdateRds(ctx, target, t)
			}
		}
	}
}

// Rds数据入库
func (s *service) SaveOrUpdateRds(ctx context.Context, ins *rds.RDS, t *task.Task) {
	b, err := s.rds.SaveRDS(ctx, ins)

	var detail *task.Record
	if err != nil {
		s.log.Warnf("save rds error, %s", err)
		detail = task.NewSyncFailedRecord(t.Id, ins.Base.InstanceId, ins.Information.Name, err.Error())
	} else {
		s.log.Debugf("save host %s to db", b.ShortDesc())
		detail = task.NewSyncSucceedRecord(t.Id, ins.Base.InstanceId, ins.Information.Name)
	}

	t.AddDetail(detail)
	if err := s.insertTaskDetail(ctx, detail); err != nil {
		s.log.Errorf("update detail error, %s", err)
	}
}
