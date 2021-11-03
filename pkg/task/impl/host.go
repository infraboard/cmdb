package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/pkg/task"

	aliConn "github.com/infraboard/cmdb/provider/aliyun/connectivity"
	ecsOp "github.com/infraboard/cmdb/provider/aliyun/ecs"
	hwConn "github.com/infraboard/cmdb/provider/huawei/connectivity"
	hwEcsOp "github.com/infraboard/cmdb/provider/huawei/ecs"
	txConn "github.com/infraboard/cmdb/provider/txyun/connectivity"
	cvmOp "github.com/infraboard/cmdb/provider/txyun/cvm"
	vsConn "github.com/infraboard/cmdb/provider/vsphere/connectivity"
	vmOp "github.com/infraboard/cmdb/provider/vsphere/vm"
)

type SyncTaskCallback func(*task.Task)

func (s *service) syncHost(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager host.Pager
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

	hs := host.NewHostSet()
	switch secret.Vendor {
	case resource.VendorAliYun:
		s.log.Debugf("sync aliyun host ...")
		client := aliConn.NewAliCloudClient(secret.APIKey, secret.APISecret, t.Region)
		ec, err := client.EcsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := ecsOp.NewEcsOperater(ec)
		req := ecsOp.NewPageQueryRequest()
		req.Rate = secret.RequestRate
		pager = operater.PageQuery(req)
	case resource.VendorTencent:
		s.log.Debugf("sync txyun host ...")
		client := txConn.NewTencentCloudClient(secret.APIKey, secret.APISecret, t.Region)
		operater := cvmOp.NewCVMOperater(client.CvmClient())
		pager = operater.PageQuery()
	case resource.VendorHuaWei:
		s.log.Debugf("sync hwyun host ...")
		client := hwConn.NewHuaweiCloudClient(secret.APIKey, secret.APISecret, t.Region)
		ec, err := client.EcsClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := hwEcsOp.NewEcsOperater(ec)
		pager = operater.PageQuery()
	case resource.VendorVsphere:
		s.log.Debugf("sync vshpere host ...")
		client := vsConn.NewVsphereClient(secret.Address, secret.APIKey, secret.APISecret)
		ec, err := client.VimClient()
		if err != nil {
			t.Failed(err.Error())
			return
		}
		operater := vmOp.NewVmOperater(ec)
		hs, err = operater.Query()
		if err != nil {
			t.Failed(err.Error())
			return
		}
	default:
		t.Failed(fmt.Sprintf("unsuport vendor %s", secret.Vendor))
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
				hs.Add(p.Data.Items[i])
			}
		}
	}

	// 调用host服务保持数据
	for i := range hs.Items {
		target := hs.Items[i]
		_, err := s.host.SaveHost(ctx, target)
		if err != nil {
			t.AddDetailFailed(target.Name, err.Error())
		} else {
			t.AddDetailSucceed(target.Name, "")
		}
	}
}
