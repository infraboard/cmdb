package impl

import (
	"context"

	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/pkg/syncer"
)

func (s *service) Sync(ctx context.Context, req *syncer.SyncRequest) (
	*syncer.SyncReponse, error) {
	var (
		resp *syncer.SyncReponse
		err  error
	)

	secret, err := s.DescribeSecret(ctx, syncer.NewDescribeSecretRequest(req.SecretId))
	if err != nil {
		return nil, err
	}

	switch req.ResourceType {
	case resource.HostResource:
		resp, err = s.syncHost(ctx, secret)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
