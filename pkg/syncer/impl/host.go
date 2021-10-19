package impl

import (
	"context"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/syncer"
	"github.com/infraboard/mcube/exception"

	aliConn "github.com/infraboard/cmdb/provider/aliyun/connectivity"
	ecsOp "github.com/infraboard/cmdb/provider/aliyun/ecs"
	txConn "github.com/infraboard/cmdb/provider/txyun/connectivity"
	cvmOp "github.com/infraboard/cmdb/provider/txyun/cvm"
)

func (s *service) syncHost(ctx context.Context, secret *syncer.Secret) (
	*syncer.SyncReponse, error) {
	var pager host.Pager

	switch secret.Vendor {
	case resource.VendorAliYun:
		client := aliConn.NewAliCloudClient(secret.APIKey, secret.APISecret, secret.Region)
		ec, err := client.EcsClient()
		if err != nil {
			return nil, err
		}
		operater := ecsOp.NewEcsOperater(ec)
		pager = operater.PageQuery()
	case resource.VendorTencent:
		client := txConn.NewTencentCloudClient(secret.APIKey, secret.APISecret, secret.Region)
		operater := cvmOp.NewCVMOperater(client.CvmClient())
		pager = operater.PageQuery()
	default:
		return nil, exception.NewBadRequest("unsuport vendor %s", secret.Vendor)
	}

	set := syncer.NewSyncReponse()
	// 分页查询数据
	hasNext := true
	for hasNext {
		p := pager.Next()
		hasNext = p.HasNext

		// 调用host服务保持数据
		for i := range p.Data.Items {
			ins, err := s.host.SaveHost(ctx, p.Data.Items[i])
			if err != nil {
				set.AddFailed(ins.Name, err.Error())
			} else {
				set.AddSucceed(ins.Name, "")
			}
		}

	}

	return set, nil
}
