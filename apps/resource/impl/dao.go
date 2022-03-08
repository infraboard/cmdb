package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/sqlbuilder"
)

func QueryTag(ctx context.Context, db *sql.DB, resourceIds []string) (
	tags []*resource.Tag, err error) {
	if len(resourceIds) == 0 {
		return
	}

	args, pos := []interface{}{}, []string{}
	for _, id := range resourceIds {
		args = append(args, id)
		pos = append(pos, "?")
	}

	query := sqlbuilder.NewQuery(sqlQueryResourceTag)
	inWhere := fmt.Sprintf("resource_id IN (%s)", strings.Join(pos, ","))
	query.Where(inWhere, args...)
	querySQL, args := query.BuildQuery()
	zap.L().Debugf("sql: %s", querySQL)

	queryStmt, err := db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query resource tag error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		ins := resource.NewDefaultTag()
		err := rows.Scan(
			&ins.Key, &ins.Value, &ins.Describe, &ins.ResourceId, &ins.Weight, &ins.Type,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query resource tag error, %s", err.Error())
		}
		tags = append(tags, ins)
	}

	return
}

func (s *service) addTag(ctx context.Context, resourceId string, tags []*resource.Tag) (
	err error) {
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

func (s *service) removeTag(ctx context.Context, resourceId string, tags []*resource.Tag) (
	err error) {
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

	stmt, err = tx.Prepare(sqlDeleteResourceTag)
	if err != nil {
		err = fmt.Errorf("prepare delete tag sql error, %s", err)
		return
	}
	defer stmt.Close()

	for i := range tags {
		if _, err = stmt.Exec(resourceId, tags[i].Key, tags[i].Value); err != nil {
			err = fmt.Errorf("save resource tag error, %s", err)
			return
		}
	}

	return
}
