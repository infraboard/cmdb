package impl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource/impl"
)

func (s *service) save(ctx context.Context, h *host.Host) error {
	if h.Resource.Meta.SyncAt != 0 {
		h.Resource.Meta.SyncAt = time.Now().UnixMicro()
	}

	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	// wiki: https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
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

	// 保存资源基础信息
	err = impl.SaveResource(ctx, tx, h.Resource)
	if err != nil {
		return err
	}

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.PrepareContext(ctx, insertHostSQL)
	if err != nil {
		return fmt.Errorf("prepare insert host sql error, %s", err)
	}
	defer stmt.Close()

	desc := h.Describe
	_, err = stmt.ExecContext(ctx,
		h.Resource.Meta.Id, desc.GpuSpec, desc.OsType, desc.OsName,
		desc.ImageId, desc.InternetMaxBandwidthOut,
		desc.InternetMaxBandwidthIn, desc.KeyPairNameToString(), desc.SecurityGroupsToString(),
	)
	if err != nil {
		return fmt.Errorf("save host resource describe error, %s", err)
	}

	return err
}

func (s *service) update(ctx context.Context, ins *host.Host) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	// 更新资源基本信息
	if err := impl.UpdateResource(ctx, tx, ins.Resource.Meta, ins.Resource.Spec, ins.Resource.Tags); err != nil {
		return err
	}

	// 更新实例信息
	stmt, err = tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return fmt.Errorf("prepare update host sql error, %s", err)
	}
	defer stmt.Close()

	base := ins.Resource.Meta
	desc := ins.Describe
	_, err = stmt.ExecContext(ctx,
		desc.GpuSpec, desc.OsType, desc.OsName,
		desc.ImageId, desc.InternetMaxBandwidthOut,
		desc.InternetMaxBandwidthIn, desc.KeyPairNameToString(), desc.SecurityGroupsToString(),
		base.Id,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *service) delete(ctx context.Context, req *host.ReleaseHostRequest) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	stmt, err = tx.PrepareContext(ctx, deleteHostSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, req.Id)
	if err != nil {
		return err
	}

	if err := impl.DeleteResource(ctx, tx, req.Id); err != nil {
		return err
	}

	return tx.Commit()
}
