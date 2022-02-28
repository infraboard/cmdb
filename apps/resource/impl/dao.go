package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/infraboard/cmdb/apps/resource"
)

func (s *service) addTag(ctx context.Context, resourceId string, tags []*resource.Tag) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start add tag tx error, %s", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	err = updateResourceTag(tx, resourceId, tags)
	return
}

func (s *service) removeTag(ctx context.Context, tags []*resource.Tag) (err error) {
	var (
		stmt *sql.Stmt
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
		tx.Commit()
	}()

	stmt, err = tx.Prepare(SQLDeleteResourceTag)
	if err != nil {
		err = fmt.Errorf("prepare delete tag sql error, %s", err)
		return
	}
	defer stmt.Close()

	for i := range tags {
		if _, err = stmt.Exec(tags[i].Id); err != nil {
			err = fmt.Errorf("save resource tag error, %s", err)
			return
		}
	}

	return
}
