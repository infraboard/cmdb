package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/credential"
)

func (s *service) deleteSecret(ctx context.Context, ins *credential.Secret) error {
	if ins == nil {
		return fmt.Errorf("credential is nil")
	}

	stmt, err := s.db.PrepareContext(ctx, deleteSecretSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, ins.Id)
	if err != nil {
		return err
	}

	return nil
}
