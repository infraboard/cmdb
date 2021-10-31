package impl

import (
	"context"

	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
)

const (
	queryResourceSQL = `SELECT * FROM resource`
)

func (s *service) Search(ctx context.Context, req *resource.SearchRequest) (
	*resource.ResourceSet, error) {
	query := sqlbuilder.NewQuery(queryResourceSQL)

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

	querySQL, args := query.Order("sync_at").Desc().Limit(req.OffSet(), uint(req.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query host error, %s", err.Error())
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
		err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
			&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name, &ins.Description,
			&ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
			&publicIPList, &privateIPList, &ins.PayType, &ins.DescribeHash, &ins.ResourceHash,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
		}
		ins.LoadPrivateIPString(privateIPList)
		ins.LoadPublicIPString(publicIPList)
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
