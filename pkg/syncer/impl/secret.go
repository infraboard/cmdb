package impl

import (
	"context"

	"github.com/infraboard/cmdb/pkg/syncer"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
)

const (
	insertSecretSQL = `INSERT INTO secret (
		id,create_at,description,vendor,crendential_type,api_key,api_secret
	) VALUES (?,?,?,?,?,?,?);`

	querySecretSQL = `SELECT * FROM secret`
)

func (s *service) CreateSecret(ctx context.Context, req *syncer.CreateSecretRequest) (
	*syncer.Secret, error) {
	ins, err := syncer.NewSecret(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create secret error, %s", err)
	}

	stmt, err := s.db.Prepare(insertSecretSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ins.Id, ins.CreateAt, ins.Description, ins.Vendor,
		ins.CrendentialType, ins.APIKey, ins.APISecret,
	)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) QuerySecret(ctx context.Context, req *syncer.QuerySecretRequest) (
	*syncer.SecretSet, error) {
	query := sqlbuilder.NewQuery(querySecretSQL)

	querySQL, args := query.Order("create_at").Desc().BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query secret error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := syncer.NewSecretSet()
	for rows.Next() {
		ins := syncer.NewDefaultSecret()
		err := rows.Scan(
			&ins.Id, &ins.CreateAt, &ins.Description, &ins.Vendor, &ins.CrendentialType,
			&ins.APIKey, &ins.APISecret,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
		}
		set.Add(ins)
	}
	return set, nil
}
