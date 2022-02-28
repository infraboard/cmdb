package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/infraboard/cmdb/apps/resource"
)

func (s *service) addTag(ctx context.Context, resourceId string, tags []*resource.Tag) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start add tag tx error, %s", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	stmt, err = tx.Prepare(SQLInsertOrUpdateResourceTag)
	if err != nil {
		return fmt.Errorf("prepare update tag sql error, %s", err)
	}
	defer stmt.Close()

	if err := updateResourceTag(tx, resourceId, tags); err != nil {
		return err
	}

	return nil
}

func (s *service) removeTag(ctx context.Context, tags []*resource.Tag) error {
	return nil
}
