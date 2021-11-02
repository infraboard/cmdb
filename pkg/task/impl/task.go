package impl

import (
	"context"

	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/pkg/task"
	"github.com/infraboard/mcube/exception"
)

func (s *service) CreatTask(ctx context.Context, req *task.CreateTaskRequst) (
	*task.Task, error) {
	var (
		resp *task.Task
		err  error
	)

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate sync request error, %s", err)
	}

	secret, err := s.secret.DescribeSecret(ctx, secret.NewDescribeSecretRequest(req.SecretId))
	if err != nil {
		return nil, err
	}

	// 如果不是vsphere 需要检查region
	if !secret.Vendor.Equal(resource.VendorVsphere) {
		if req.Region == "" {
			return nil, exception.NewBadRequest("region required")
		}
		if !secret.IsAllowRegion(req.Region) {
			return nil, exception.NewBadRequest("this secret not allow sync region %s", req.Region)
		}
	}

	switch req.ResourceType {
	case resource.HostResource:
		resp, err = s.syncHost(ctx, secret, req.Region)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
