package impl

// func SaveResource(ctx context.Context, tx *sql.Tx, ins *resource.Resource) error {
// 	// 避免SQL注入, 请使用Prepare
// 	stmt, err := tx.PrepareContext(ctx, sqlInsertResource)
// 	if err != nil {
// 		return fmt.Errorf("prepare insert resource error, %s", err)
// 	}
// 	defer stmt.Close()

// 	base := ins.Meta
// 	info := ins.Spec
// 	// 保存资源数据
// 	_, err = stmt.ExecContext(ctx,
// 		base.Id, info.ResourceType, info.Vendor, info.Region, info.Zone, base.CreateAt, info.ExpireAt, info.Category, info.Type,
// 		info.Name, info.Description, ins.Status.Phase, info.UpdateAt, base.SyncAt, info.Owner, ins.Status.PublicAddressToString(),
// 		ins.Status.PrivateAddressToString(), ins.Cost.PAY_MODE, base.DescribeHash, base.ResourceHash, base.CredentialId,
// 		base.Domain, base.Namespace, base.Env, base.UsageMode,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("save host resource info error, %s", err)
// 	}

// 	if err := updateResourceTag(ctx, tx, base.Id, ins.Tags); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UpdateResource(ctx context.Context, tx *sql.Tx, base *resource.Meta, info *resource.Spec, tags []*resource.Tag) error {
// 	// 避免SQL注入, 请使用Prepare
// 	stmt, err := tx.PrepareContext(ctx, sqlUpdateResource)
// 	if err != nil {
// 		return fmt.Errorf("prepare update reousrce sql error, %s", err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx,
// 		info.ExpireAt, info.Category, info.Type, info.Name, info.Description,
// 		"", info.UpdateAt, base.SyncAt, info.Owner,
// 		"", "", "", base.DescribeHash, base.ResourceHash,
// 		base.CredentialId, base.Namespace, base.Env, base.UsageMode,
// 		base.Id,
// 	)

// 	if err != nil {
// 		return fmt.Errorf("update resource base info error, %s", err)
// 	}

// 	if err := updateResourceTag(ctx, tx, base.Id, tags); err != nil {
// 		return fmt.Errorf("update resource tag error, %s", err)
// 	}

// 	return nil
// }

// func DeleteResource(ctx context.Context, tx *sql.Tx, id string) error {
// 	stmt, err := tx.PrepareContext(ctx, sqlDeleteResource)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	stmt, err = tx.PrepareContext(ctx, sqlDeleteResourceTag)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func updateResourceTag(ctx context.Context, tx *sql.Tx, resourceId string, tags []*resource.Tag) error {
// 	// 保存资源标签
// 	stmt, err := tx.PrepareContext(ctx, sqlInsertOrUpdateResourceTag)
// 	if err != nil {
// 		return fmt.Errorf("prepare update resource tag error, %s", err)
// 	}
// 	defer stmt.Close()

// 	for i := range tags {
// 		t := tags[i]
// 		if t.Weight == 0 {
// 			t.Weight = 1
// 		}
// 		_, err = stmt.ExecContext(ctx,
// 			t.Purpose, t.Key, t.Value, t.Describe, resourceId, t.Weight, time.Now().UnixMilli(),
// 			t.Describe, t.Weight,
// 		)
// 		if err != nil {
// 			return fmt.Errorf("save resource tag error, %s", err)
// 		}
// 	}

// 	return nil
// }
