package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/cmdb/app/host"
	"github.com/infraboard/cmdb/app/resource/impl"
)

func (s *service) SaveHost(ctx context.Context, h *host.Host) (
	*host.Host, error) {
	h.Base.Id = xid.New().String()
	h.Describe.ResourceId = h.Base.Id
	h.Base.SyncAt = ftime.Now().Timestamp()

	if err := s.save(ctx, h); err != nil {
		return nil, err
	}
	return h, nil
}

func (s *service) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)

	if req.Keywords != "" {
		query.Where("r.name LIKE ? OR r.id = ? OR r.instance_id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords,
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	querySQL, args := query.Order("sync_at").Desc().Limit(req.OffSet(), uint(req.Page.PageSize)).BuildQuery()
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

	set := host.NewHostSet()
	var (
		publicIPList, privateIPList, keyPairNameList, securityGroupsList string
	)
	for rows.Next() {
		ins := host.NewDefaultHost()
		base := ins.Base
		info := ins.Information
		desc := ins.Describe
		err := rows.Scan(
			&base.Id, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
			&info.Category, &info.Type, &base.InstanceId, &info.Name, &info.Description,
			&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
			&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
			&base.SecretId, &desc.ResourceId,
			&desc.Cpu, &desc.Memory, &desc.GpuAmount, &desc.GpuSpec, &desc.OsType, &desc.OsName,
			&desc.SerialNumber, &desc.ImageId, &desc.InternetMaxBandwidthOut, &desc.InternetMaxBandwidthIn,
			&keyPairNameList, &securityGroupsList,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
		}
		info.LoadPrivateIPString(privateIPList)
		info.LoadPublicIPString(publicIPList)
		desc.LoadKeyPairNameString(keyPairNameList)
		desc.LoadSecurityGroupsString(securityGroupsList)
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

func (s *service) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (
	*host.Host, error) {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 检测参数合法性
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update host error, %s", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("start tx error, %s", err)
	}

	// 查询出该条实例的数据
	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	oldRH, oldDH := ins.Base.ResourceHash, ins.Base.DescribeHash

	switch req.UpdateMode {
	case host.UpdateMode_PATCH:
		ins.Patch(req.UpdateHostData)
	default:
		ins.Put(req.UpdateHostData)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if oldRH != ins.Base.ResourceHash {
		// 避免SQL注入, 请使用Prepare
		stmt, err = tx.Prepare(impl.SQLDeleteResource)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		base := ins.Base
		info := ins.Information
		_, err = stmt.Exec(
			info.ExpireAt, info.Category, info.Type, info.Name, info.Description,
			info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount,
			info.PublicIp, info.PrivateIp, info.PayType, base.DescribeHash, base.ResourceHash,
			ins.Base.SecretId, ins.Describe.ResourceId,
		)
		if err != nil {
			return nil, err
		}
	} else {
		s.log.Debug("resource data hash not changed, needn't update")
	}

	if oldDH != ins.Base.DescribeHash {
		// 避免SQL注入, 请使用Prepare
		stmt, err = tx.Prepare(updateHostSQL)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		base := ins.Base
		desc := ins.Describe
		_, err = stmt.Exec(
			desc.Cpu, desc.Memory, desc.GpuAmount, desc.GpuSpec, desc.OsType, desc.OsName,
			desc.ImageId, desc.InternetMaxBandwidthOut,
			desc.InternetMaxBandwidthIn, desc.KeyPairNameToString(), desc.SecurityGroupsToString(),
			base.Id,
		)
		if err != nil {
			return nil, err
		}
	} else {
		s.log.Debug("describe data hash not changed, needn't update")
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (
	*host.Host, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)
	cond, val := req.Where()
	querySQL, args := query.Where(cond, val).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query host error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := host.NewDefaultHost()
	var (
		publicIPList, privateIPList, keyPairNameList, securityGroupsList string
	)
	base := ins.Base
	info := ins.Information
	desc := ins.Describe
	err = queryStmt.QueryRow(args...).Scan(
		&base.Id, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
		&info.Category, &info.Type, &base.InstanceId, &info.Name, &info.Description,
		&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
		&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash,
		&base.SecretId, &desc.ResourceId,
		&desc.Cpu, &desc.Memory, &desc.GpuAmount, &desc.GpuSpec, &desc.OsType, &desc.OsName,
		&desc.SerialNumber, &desc.ImageId, &desc.InternetMaxBandwidthOut, &desc.InternetMaxBandwidthIn,
		&keyPairNameList, &securityGroupsList,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("describe host error, %s", err.Error())
	}

	info.LoadPrivateIPString(privateIPList)
	info.LoadPublicIPString(publicIPList)
	desc.LoadKeyPairNameString(keyPairNameList)
	desc.LoadSecurityGroupsString(securityGroupsList)

	return ins, nil
}

func (s *service) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (
	*host.Host, error) {
	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.delete(ctx, req); err != nil {
		return nil, err
	}

	return ins, nil
}
