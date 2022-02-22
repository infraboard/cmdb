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

	// 保存资源数据
	_, err = stmt.Exec(
		base.Id, base.Vendor, base.Region, base.Zone, base.CreateAt, info.ExpireAt, info.Category, info.Type,
		info.Name, info.Description, info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount, info.PublicIPToString(),
		info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash, base.SecretId,
	)
	if err != nil {
		return fmt.Errorf("save host resource info error, %s", err)
	}

	if err := updateResourceTag(tx, base.Id, info.Tags); err != nil {
		return err
	}

	return nil
}

func UpdateResource(tx *sql.Tx, base *resource.Base, info *resource.Information) error {
	// 避免SQL注入, 请使用Prepare
	stmt, err := tx.Prepare(SQLUpdateResource)
	if err != nil {
		return fmt.Errorf("prepare update reousrce sql error, %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		info.ExpireAt, info.Category, info.Type, info.Name, info.Description,
		info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount,
		info.PublicIPToString(), info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash,
		base.SecretId, base.Id,
	)
	if err != nil {
		return err
	}

	if err := updateResourceTag(tx, base.Id, info.Tags); err != nil {
		return err
	}

	return nil
}

func DeleteResource(tx *sql.Tx, id string) error {
	stmt, err := tx.Prepare(SQLDeleteResource)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare(SQLDeleteResourceTag)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func updateResourceTag(tx *sql.Tx, resourceId string, tags []*resource.Tag) error {
	// 保存资源标签
	stmt, err := tx.Prepare(SQLInsertTag)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := range tags {
		t := tags[i]
		_, err = stmt.Exec(
			t.Key, t.Value, t.Describe, resourceId,
		)
		if err != nil {
			return fmt.Errorf("save resource tag error, %s", err)
		}
	}

	return nil
}
