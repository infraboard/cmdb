package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/resource"
)

func (s *service) addTag(ctx context.Context, tags []*resource.Tag) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start add tag tx error, %s", err)
	}

	stmt, err := tx.Prepare(SQLDeleteResourceTag)
	if err != nil {
		return fmt.Errorf("prepare update tag sql error, %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s *service) updateTag(ctx context.Context, tags []*resource.Tag) error {
	return nil
}

func (s *service) removeTag(ctx context.Context, tags []*resource.Tag) error {
	return nil
}
