package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/cmdb/apps/credential"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/task"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/provider/aliyun"
	"github.com/infraboard/cmdb/provider/aws"
	"github.com/infraboard/cmdb/provider/huawei"
	"github.com/infraboard/cmdb/provider/txyun"
	vsConn "github.com/infraboard/cmdb/provider/vsphere/connectivity"
	vmOp "github.com/infraboard/cmdb/provider/vsphere/vm"
)

func (s *service) syncHost(ctx context.Context, credentialIns *credential.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager pager.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		if err := recover(); err != nil {
			t.Failed(fmt.Sprintf("pannic, %v", err))
		} else {
			t.Completed()
		}
		cb(t)
	}()

	credential := credentialIns.Data
	req := provider.NewQueryHostRequestWithRate(credential.RequestRate)

	switch credential.Vendor {
	case resource.Vendor_ALIYUN:
		s.log.Debugf("sync aliyun ecs ...")
		op, err := aliyun.NewOperator(credential.ApiKey, credential.ApiSecret, t.Data.Region)
		if err != nil {
			t.Failed(err.Error())
			return
		}
		pager = op.HostOperator().QueryHost(req)
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun cvm ...")
		op, err := txyun.NewOperator(credential.ApiKey, credential.ApiSecret, t.Data.Region)
		if err != nil {
			t.Failed(err.Error())
			return
		}
		pager = op.HostOperator().QueryHost(req)
	case resource.Vendor_HUAWEI:
		s.log.Debugf("sync hwyun ecs ...")
		op, err := huawei.NewOperator(credential.ApiKey, credential.ApiSecret, t.Data.Region)
		if err != nil {
			t.Failed(err.Error())
			return
		}
		pager = op.HostOperator().QueryHost(req)
	case resource.Vendor_AMAZON:
		s.log.Debugf("sync aws ec2 ...")
		op := aws.NewOperator(credential.ApiKey, credential.ApiSecret, t.Data.Region)
		pager = op.HostOperator().QueryHost(req)
	case resource.Vendor_VSPHERE:
		s.log.Debugf("sync vshpere vm ...")
		client := vsConn.NewVsphereClient(credential.Address, credential.ApiKey, credential.ApiSecret)
		ec, err := client.VimClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operator := vmOp.NewVMOperator(ec)
		// 通过回调直接保存
		err = operator.QueryHost(func(h *host.Host) {
			// 补充管理信息
			h.Base.CredentialId = credentialIns.Id
			s.doSyncHost(ctx, h, t)
		})
		if err != nil {
			t.Failed(err.Error())
			return
		}
	default:
		t.Failed(fmt.Sprintf("unsuport vendor %s", credential.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		for pager.Next() {
			set := host.NewHostSet()
			if err := pager.Scan(ctx, set); err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", err))
				return
			}

			// 调用host服务保持数据
			for i := range set.Items {
				target := set.Items[i]
				// 补充管理信息
				target.Base.CredentialId = credentialIns.Id
				s.doSyncHost(ctx, target, t)
			}
		}
	}
}

// Host主机数据入库
func (s *service) doSyncHost(ctx context.Context, ins *host.Host, t *task.Task) {
	// 添加Host
	ins.Base.SyncAt = time.Now().UnixMilli()
	h, err := s.host.SyncHost(ctx, ins)

	// 添加同步详情
	var detail *task.Record
	if err != nil {
		s.log.Warnf("save host error, %s", err)
		detail = task.NewSyncFailedRecord(t.Id, ins.Base.Id, ins.Information.Name, err.Error())
	} else {
		s.log.Debugf("save host %s to db", h.ShortDesc())
		detail = task.NewSyncSucceedRecord(t.Id, ins.Base.Id, ins.Information.Name)
	}

	t.AddDetail(detail)
	if err := s.insertTaskDetail(ctx, detail); err != nil {
		s.log.Errorf("update detail error, %s", err)
	}
}
