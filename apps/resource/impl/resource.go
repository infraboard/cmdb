package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
)

func (s *service) Search(ctx context.Context, req *resource.SearchRequest) (
	*resource.ResourceSet, error) {
	query := sqlbuilder.NewQuery(SQLQueryResource)

	if req.Keywords != "" {
		query.Where("name LIKE ? OR id = ? OR instance_id = ? OR private_ip LIKE ? OR public_ip LIKE ?",
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords,
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	if req.Vendor != nil {
		query.Where("vendor = ?", req.Vendor)
	}

	querySQL, args := query.Order("sync_at").Desc().Limit(req.Page.ComputeOffset(), uint(req.Page.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query resource error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var (
		publicIPList, privateIPList string
	)
	set := resource.NewResourceSet()
	for rows.Next() {
		ins := resource.NewDefaultResource()
		base := ins.Base
		info := ins.Information
		err := rows.Scan(
			&base.Id, &base.ResourceType, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
			&info.Category, &info.Type, &info.Name, &info.Description,
			&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
			&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
			&base.SecretId, &base.Domain, &base.Namespace, &base.Env, &base.UsageMode,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query resource error, %s", err.Error())
		}
		info.LoadPrivateIPString(privateIPList)
		info.LoadPublicIPString(publicIPList)
		set.Add(ins)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return set, nil
}

func (s *service) UpdateTag(ctx context.Context, req *resource.UpdateTagRequest) (
	*resource.Resource, error) {
	switch req.Action {
	case resource.UpdateAction_ADD:
		s.addTag(ctx, req.Id, req.Tags)
	case resource.UpdateAction_REMOVE:
		s.removeTag(ctx, req.Tags)
	default:
		return nil, fmt.Errorf("unknow update tag action: %s", req.Action)
	}
	return nil, nil
}
