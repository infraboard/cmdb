package impl

import (
	"context"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/syncer"
	"github.com/infraboard/mcube/exception"

	aliConn "github.com/infraboard/cmdb/provider/aliyun/connectivity"
	ecsOp "github.com/infraboard/cmdb/provider/aliyun/ecs"
	hwConn "github.com/infraboard/cmdb/provider/huawei/connectivity"
	hwEcsOp "github.com/infraboard/cmdb/provider/huawei/ecs"
	txConn "github.com/infraboard/cmdb/provider/txyun/connectivity"
	cvmOp "github.com/infraboard/cmdb/provider/txyun/cvm"
	vsConn "github.com/infraboard/cmdb/provider/vsphere/connectivity"
	vmOp "github.com/infraboard/cmdb/provider/vsphere/vm"
)

func (s *service) syncHost(ctx context.Context, secret *syncer.Secret, region string) (
	*syncer.SyncReponse, error) {
	var (
		pager host.Pager
	)

	hs := host.NewHostSet()
	switch secret.Vendor {
	case resource.VendorAliYun:
		client := aliConn.NewAliCloudClient(secret.APIKey, secret.APISecret, region)
		ec, err := client.EcsClient()
		if err != nil {
			return nil, err
		}
		operater := ecsOp.NewEcsOperater(ec)
		pager = operater.PageQuery()
	case resource.VendorTencent:
		client := txConn.NewTencentCloudClient(secret.APIKey, secret.APISecret, region)
		operater := cvmOp.NewCVMOperater(client.CvmClient())
		pager = operater.PageQuery()
	case resource.VendorHuaWei:
		client := hwConn.NewHuaweiCloudClient(secret.APIKey, secret.APISecret, region)
		ec, err := client.EcsClient()
		if err != nil {
			return nil, err
		}
		operater := hwEcsOp.NewEcsOperater(ec)
		pager = operater.PageQuery()
	case resource.VendorVsphere:
		client := vsConn.NewVsphereClient(secret.Address, secret.APISecret, secret.APISecret)
		ec, err := client.VimClient()
		if err != nil {
			return nil, err
		}
		operater := vmOp.NewVmOperater(ec)
		hs, err = operater.Query()
		if err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("unsuport vendor %s", secret.Vendor)
	}

	set := syncer.NewSyncReponse()
	// 分页查询数据
	if pager != nil {
		hasNext := true
		for hasNext {
			p := pager.Next()
			hasNext = p.HasNext

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
			set.AddFailed(target.Name, err.Error())
		} else {
			set.AddSucceed(target.Name, "")
		}
	}

	return set, nil
}
