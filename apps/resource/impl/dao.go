package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/infraboard/cmdb/apps/resource"
	"gorm.io/gorm"
)

func (s *service) BuildQuery(ctx context.Context, req *resource.SearchRequest) (*gorm.DB, error) {
	query := s.db.WithContext(ctx).
		Table("resource_meta m").
		Joins("LEFT JOIN resource_spec spsc ON spec.resource_id=m.id").
		Joins("LEFT JOIN resource_cost cost ON cost.resource_id=m.id").
		Joins("LEFT JOIN resource_status status ON status.resource_id=m.id")

	if req.Domain != "" {
		query = query.Where("m.domain = ?", req.Domain)
	}
	if req.Namespace != "" {
		query = query.Where("m.namespace = ?", req.Namespace)
	}
	if req.Env != "" {
		query = query.Where("m.env = ?", req.Env)
	}
	if req.UsageMode != nil {
		query = query.Where("m.usage_mode = ?", *req.UsageMode)
	}
	if req.Vendor != nil {
		query = query.Where("spec.vendor = ?", req.Vendor)
	}
	if req.Owner != "" {
		query = query.Where("spec.owner = ?", req.Owner)
	}
	if req.Type != nil {
		query = query.Where("spec.type = ?", req.Type)
	}
	if req.Status != "" {
		query = query.Where("status.phase = ?", req.Status)
	}
	if req.HasTag() {
		query = query.Joins("RIGHT JOIN resource_tag tag ON tag.resource_id = m.id")
		for i := range req.Tags {
			selector := req.Tags[i]
			if selector.Key == "" {
				continue
			}

			// 添加key过滤条件
			query = query.Where("tag.t_key LIKE ?", strings.ReplaceAll(selector.Key, ".*", "%"))

			// 添加Value过滤条件
			condtions := []string{}
			args := []interface{}{}
			for _, v := range selector.Values {
				condtions = append(condtions, fmt.Sprintf("t.t_value %s ?", selector.Opertor))
				args = append(args, strings.ReplaceAll(v, ".*", "%"))
			}
			if len(condtions) > 0 {
				vwhere := fmt.Sprintf("( %s )", strings.Join(condtions, selector.RelationShip()))
				query = query.Where(vwhere, args...)
			}
		}
	}
	if req.Keywords != "" {
		query = query.Where("r.name LIKE ? OR r.id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	return query, nil
}

// func QueryTag(ctx context.Context, db *sql.DB, resourceIds []string) (
// 	tags []*resource.Tag, err error) {
// 	if len(resourceIds) == 0 {
// 		return
// 	}

// 	args, pos := []interface{}{}, []string{}
// 	for _, id := range resourceIds {
// 		args = append(args, id)
// 		pos = append(pos, "?")
// 	}

// 	query := sqlbuilder.NewQuery(sqlQueryResourceTag)
// 	inWhere := fmt.Sprintf("resource_id IN (%s)", strings.Join(pos, ","))
// 	query.Where(inWhere, args...)
// 	querySQL, args := query.BuildQuery()
// 	zap.L().Debugf("sql: %s", querySQL)

// 	queryStmt, err := db.PrepareContext(ctx, querySQL)
// 	if err != nil {
// 		return nil, exception.NewInternalServerError("prepare query resource tag error, %s", err.Error())
// 	}
// 	defer queryStmt.Close()

// 	rows, err := queryStmt.QueryContext(ctx, args...)
// 	if err != nil {
// 		return nil, exception.NewInternalServerError(err.Error())
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		ins := resource.NewDefaultTag()
// 		err := rows.Scan(
// 			&ins.Key, &ins.Value, &ins.Describe, "", &ins.Weight, &ins.Purpose,
// 		)
// 		if err != nil {
// 			return nil, exception.NewInternalServerError("query resource tag error, %s", err.Error())
// 		}
// 		tags = append(tags, ins)
// 	}

// 	return
// }

// func (s *service) addTag(ctx context.Context, resourceId string, tags []*resource.Tag) (
// 	err error) {
// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return fmt.Errorf("start add tag tx error, %s", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 			return
// 		}
// 		tx.Commit()
// 	}()

// 	err = updateResourceTag(ctx, tx, resourceId, tags)
// 	return
// }

// func (s *service) removeTag(ctx context.Context, resourceId string, tags []*resource.Tag) (
// 	err error) {
// 	var (
// 		stmt *sql.Stmt
// 	)

// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return fmt.Errorf("start add tag tx error, %s", err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 			return
// 		}
// 		tx.Commit()
// 	}()

// 	stmt, err = tx.PrepareContext(ctx, sqlDeleteResourceTag)
// 	if err != nil {
// 		err = fmt.Errorf("prepare delete tag sql error, %s", err)
// 		return
// 	}
// 	defer stmt.Close()

// 	for i := range tags {
// 		if _, err = stmt.ExecContext(ctx, resourceId, tags[i].Key, tags[i].Value); err != nil {
// 			err = fmt.Errorf("save resource tag error, %s", err)
// 			return
// 		}
// 	}

// 	return
// }
