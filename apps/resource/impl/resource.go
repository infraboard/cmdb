package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
)

func (s *service) Search(ctx context.Context, req *resource.SearchRequest) (
	*resource.ResourceSet, error) {
	query := sqlbuilder.NewQuery(SQLQueryResource)
	s.buildQuery(query, req)

	set := resource.NewResourceSet()

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCountWith("COUNT(DISTINCT r.id)")
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	// 获取分页数据
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
		publicIPList, privateIPList                          string
		tagKeys, tagValues, tagDescribe, tagWeighs, tagTypes string
	)

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
			&tagKeys, &tagValues, &tagDescribe, &tagWeighs, &tagTypes,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query resource error, %s", err.Error())
		}
		info.LoadPrivateIPString(privateIPList)
		info.LoadPublicIPString(publicIPList)
		if err := info.LoadTags(tagKeys, tagValues, tagDescribe, tagWeighs, tagTypes); err != nil {
			s.log.Error("load tags error, %s", err)
		}
		set.Add(ins)
	}

	return set, nil
}

func (s *service) buildQuery(query *sqlbuilder.Query, req *resource.SearchRequest) {
	if req.Keywords != "" {
		if req.ExactMatch {
			// 精确匹配
			query.Where("r.name = ? OR r.id = ? OR r.private_ip = ? OR r.public_ip = ?",
				req.Keywords,
				req.Keywords,
				req.Keywords,
				req.Keywords,
			)
		} else {
			// 模糊匹配
			query.Where("r.name LIKE ? OR r.id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
				"%"+req.Keywords+"%",
				req.Keywords,
				req.Keywords+"%",
				req.Keywords+"%",
			)
		}
	}

	if req.Domain != "" {
		query.Where("r.domain = ?", req.Domain)
	}
	if req.Namespace != "" {
		query.Where("r.namespace = ?", req.Namespace)
	}
	if req.Env != "" {
		query.Where("r.env = ?", req.Env)
	}
	if req.UsageMode != nil {
		query.Where("r.usage_mode = ?", req.UsageMode)
	}
	if req.Vendor != nil {
		query.Where("r.vendor = ?", req.Vendor)
	}
	if req.SyncAccount != "" {
		query.Where("r.sync_accout = ?", req.SyncAccount)
	}
	if req.Type != nil {
		query.Where("r.resource_type = ?", req.Type)
	}
	if req.Status != "" {
		query.Where("r.status = ?", req.Status)
	}

	// Tag过滤
	for k, v := range req.GroupByKey() {
		inset := []string{}
		if len(v) == 0 {
			continue
		}
		for i := range v {
			inset = append(inset, v[i].Value)
		}
		query.Where(fmt.Sprintf("t.t_key=? AND t.t_value IN (%s)", strings.Join(inset, ",")), k)
	}
}

func (s *service) UpdateTag(ctx context.Context, req *resource.UpdateTagRequest) (
	*resource.Resource, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update tag request error, %s", err)
	}
	switch req.Action {
	case resource.UpdateAction_ADD:
		s.addTag(ctx, req.Id, req.Tags)
	case resource.UpdateAction_REMOVE:
		s.removeTag(ctx, req.Id, req.Tags)
	default:
		return nil, fmt.Errorf("unknow update tag action: %s", req.Action)
	}
	return nil, nil
}
