package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/infraboard/cmdb/apps/rds"
)

func (s *service) save(ctx context.Context, h *rds.Rds) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				s.log.Errorf("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				s.log.Errorf("commit error, %s", err)
			}
		}
	}()

	// 生成描写信息的Hash
	if err := h.GenHash(); err != nil {
		return err
	}

	// err = impl.SaveResource(ctx, tx, h.Resource)
	// if err != nil {
	// 	return err
	// }

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(insertRdsSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	desc := h.Describe
	_, err = stmt.Exec(
		h.Resource.Meta.Id, desc.EngineType, desc.EngineVersion, desc.InstanceClass, desc.ClassType, desc.ExportType,
		desc.NetworkType, desc.Type, desc.DbMaxQuantity, desc.AccountMaxQuantity, desc.MaxConnections,
		desc.MaxIops, desc.Collation, desc.TimeZone, desc.StorageType, desc.SecurityIpMode,
		desc.SecurityIpListToString(), desc.ConnectionMode, desc.IpType, desc.DeployMode,
		desc.Port, desc.ExtraToJson(),
	)
	if err != nil {
		return fmt.Errorf("save rds resource describe error, %s", err)
	}

	return err
}
