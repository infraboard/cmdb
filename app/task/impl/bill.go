package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/app/bill"
	"github.com/infraboard/cmdb/app/resource"
	"github.com/infraboard/cmdb/app/secret"
	"github.com/infraboard/cmdb/app/task"
	"github.com/infraboard/cmdb/conf"

	bssOp "github.com/infraboard/cmdb/provider/aliyun/bss"
	aliConn "github.com/infraboard/cmdb/provider/aliyun/connectivity"
	hwBssOp "github.com/infraboard/cmdb/provider/huawei/bss"
	hwConn "github.com/infraboard/cmdb/provider/huawei/connectivity"
	billOp "github.com/infraboard/cmdb/provider/txyun/billing"
	txConn "github.com/infraboard/cmdb/provider/txyun/connectivity"
)

func (s *service) syncBill(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager bill.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		t.Completed()
		cb(t)
	}()

	// 解密secret
	err := secret.DecryptAPISecret(conf.C().App.EncryptKey)
	if err != nil {
		s.log.Warnf("decrypt api secret error, %s", err)
	}

	switch secret.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun bill ...")
		client := aliConn.NewAliCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		bc, err := client.BssClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}

		operater := bssOp.NewBssOperater(bc)
		req := bssOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun bill ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		operater := billOp.NewBillingOperater(client.BillingClient())
		req := billOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun bill ...")
		client := hwConn.NewHuaweiCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		bc, err := client.BssClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := hwBssOp.NewBssOperater(bc)
		req := hwBssOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		pager = operater.PageQuery(req)
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

			// 调用host服务保持数据
			for i := range p.Data.Items {
				target := p.Data.Items[i]
				b, err := s.bill.SaveBill(ctx, target)
				if err != nil {
					s.log.Warnf("save host error, %s", err)
					t.AddDetailFailed(target.InstanceId, err.Error())
				} else {
					s.log.Debugf("save host %s to db", b.ShortDesc())
					t.AddDetailSucceed(target.InstanceId, "")
				}
			}
		}
	}

}
