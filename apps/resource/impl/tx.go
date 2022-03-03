package impl

import (
	"database/sql"
	"fmt"

	"github.com/infraboard/cmdb/apps/resource"
)

func SaveResource(tx *sql.Tx, base *resource.Base, info *resource.Information) error {
	// 避免SQL注入, 请使用Prepare
	stmt, err := tx.Prepare(sqlInsertResource)
	if err != nil {
		return fmt.Errorf("prepare insert resource error, %s", err)
	}
	defer stmt.Close()

	// 保存资源数据
	_, err = stmt.Exec(
		base.Id, base.ResourceType, base.Vendor, base.Region, base.Zone, base.CreateAt, info.ExpireAt, info.Category, info.Type,
		info.Name, info.Description, info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount, info.PublicIPToString(),
		info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash, base.SecretId,
		base.Domain, base.Namespace, base.Env, base.UsageMode,
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
	stmt, err := tx.Prepare(sqlUpdateResource)
	if err != nil {
		return fmt.Errorf("prepare update reousrce sql error, %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		info.ExpireAt, info.Category, info.Type, info.Name, info.Description,
		info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount,
		info.PublicIPToString(), info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash,
		base.SecretId, base.Namespace, base.Env, base.UsageMode,
		base.Id,
	)
	if err != nil {
		return fmt.Errorf("update resource base info error, %s", err)
	}

	if err := updateResourceTag(tx, base.Id, info.Tags); err != nil {
		return fmt.Errorf("update resource tag error, %s", err)
	}

	return nil
}

func DeleteResource(tx *sql.Tx, id string) error {
	stmt, err := tx.Prepare(sqlDeleteResource)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare(sqlDeleteResourceTag)
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
	stmt, err := tx.Prepare(sqlInsertOrUpdateResourceTag)
	if err != nil {
		return fmt.Errorf("prepare update resource tag error, %s", err)
	}
	defer stmt.Close()

	for i := range tags {
		t := tags[i]
		if t.Weight == 0 {
			t.Weight = 1
		}
		_, err = stmt.Exec(
			t.Type, t.Key, t.Value, t.Describe, resourceId, t.Weight,
			t.Describe, t.Weight,
		)
		if err != nil {
			return fmt.Errorf("save resource tag error, %s", err)
		}
	}

	return nil
}
