package impl

import (
	"context"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

func (s *service) SyncRDS(ctx context.Context, h *rds.Rds) (
	*rds.Rds, error) {
	h.Base.Id = xid.New().String()
	h.Base.SyncAt = ftime.Now().Timestamp()

	if err := s.save(ctx, h); err != nil {
		return nil, err
	}
	return h, nil
}
