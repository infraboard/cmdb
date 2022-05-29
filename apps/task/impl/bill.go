package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/cmdb/apps/task"

	bssOp "github.com/infraboard/cmdb/provider/aliyun/bss"
	aliConn "github.com/infraboard/cmdb/provider/aliyun/connectivity"
	hwBssOp "github.com/infraboard/cmdb/provider/huawei/bss"
	hwConn "github.com/infraboard/cmdb/provider/huawei/connectivity"
	billOp "github.com/infraboard/cmdb/provider/txyun/billing"
	txConn "github.com/infraboard/cmdb/provider/txyun/connectivity"
)

func (s *service) syncBill(ctx context.Context, secretIns *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager bill.Pager
	)

	// 处理任务状态
	t.Run()
	defer s.syncBillDown(ctx, t, cb)

	secret := secretIns.Data
	switch secret.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun bill ...")
		client := aliConn.NewAliCloudClient(secret.ApiKey, secret.ApiSecret, t.Data.Region)
		bc, err := client.BssClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}

		operator := bssOp.NewBssOperator(bc)
		req := bssOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		req.Month = t.Data.Params["month"]
		pager = operator.PageQuery(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun bill ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Data.Region)
		operator := billOp.NewBillingoperator(client.BillingClient())
		req := billOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		req.Month = t.Data.Params["month"]
		pager = operator.PageQuery(req)
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun bill ...")
		client := hwConn.NewHuaweiCloudClient(secret.ApiKey, secret.ApiSecret, t.Data.Region)
		bc, err := client.BssClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operator := hwBssOp.NewBssOperator(bc)
		req := hwBssOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		req.Month = t.Data.Params["month"]
		pager = operator.PageQuery(req)
	default:
		t.Failed(fmt.Sprintf("unsuport bill syncing vendor %s", secret.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		for pager.Next() {
			set := bill.NewBillSet()
			if err := pager.Scan(ctx, set); err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", err))
				return
			}
			for i := range set.Items {
				target := set.Items[i]
				target.TaskId = t.Id
				s.doSyncBill(ctx, target, t)
			}
		}
	}

}

// 月底账单数据入库
func (s *service) doSyncBill(ctx context.Context, ins *bill.Bill, t *task.Task) {
	h, err := s.bill.SyncBill(ctx, ins)

	var detail *task.Record
	if err != nil {
		s.log.Warnf("save bill error, %s", err)
		detail = task.NewSyncFailedRecord(t.Id, ins.InstanceId, ins.InstanceName, err.Error())
	} else {
		s.log.Debugf("save bill %s to db", h.ShortDesc())
		detail = task.NewSyncSucceedRecord(t.Id, ins.InstanceId, ins.InstanceName)
	}

	t.AddDetail(detail)
	if err := s.insertTaskDetail(ctx, detail); err != nil {
		s.log.Errorf("update detail error, %s", err)
	}
}

func (s *service) syncBillDown(ctx context.Context, t *task.Task, cb SyncTaskCallback) {
	t.Completed()
	cb(t)

	s.log.Debugf("task status: %s", t.Status)
	// 调用bill服务保存数据, 由于账单对象没有更新逻辑
	// 任务同步成功, 确认当前同步版本为正确版本, 删除之前的成本
	// 任务同步失败, 删除当前同步的版本
	if t.Status.Stage.Equal(task.Stage_SUCCESS) {
		resp, err := s.bill.ConfirmBill(ctx, bill.NewConfirmBillRequest(t.Id))
		if err != nil {
			s.log.Errorf("confirm bill error, %s", err)
		} else {
			s.log.Debugf("confirm bill success, total: %d bill", resp.Total)
		}
	} else {
		resp, err := s.bill.DeleteBill(ctx, bill.NewDeleteBillRequest(t.Id))
		if err != nil {
			s.log.Errorf("delete bill error, %s", err)
		} else {
			s.log.Debugf("delete bill success, total: %d bill", resp.Total)
		}
	}
}
