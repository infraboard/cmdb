package impl

import (
	"database/sql"
	"fmt"

	"github.com/infraboard/cmdb/apps/resource"
)

func SaveResource(tx *sql.Tx, base *resource.Base, info *resource.Information) error {
	// 避免SQL注入, 请使用Prepare
	stmt, err := tx.Prepare(SQLInsertResource)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		base.Id, base.Vendor, base.Region, base.Zone, base.CreateAt, info.ExpireAt, info.Category, info.Type,
		info.Name, info.Description, info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount, info.PublicIPToString(),
		info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash, base.SecretId,
	)
	if err != nil {
		return fmt.Errorf("save host resource info error, %s", err)
	}

	return nil
}
