package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/app/secret"
)

const (
	deleteSecretSQL = `DELETE FROM secret WHERE id = ?;`
)

func (s *service) deleteSecret(ctx context.Context, ins *secret.Secret) error {
	if ins == nil {
		return fmt.Errorf("secret is nil")
	}

	stmt, err := s.db.Prepare(deleteSecretSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ins.Id)
	if err != nil {
		return err
	}

	return nil
}
