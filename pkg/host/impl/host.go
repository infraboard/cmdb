package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/infraboard/cmdb/pkg/host"
)

const (
	insertResourceSQL = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertHostSQL = `INSERT INTO host (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	updateResourceSQL = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?
	WHERE id = ?`
	updateHostSQL = `UPDATE host SET 
		cpu=?,memory=?,gpu_amount=?,gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?`

	queryHostSQL      = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
	deleteHostSQL     = `DELETE FROM host WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)

func (s *service) SaveHost(ctx context.Context, h *host.Host) (
	*host.Host, error) {
	h.Id = xid.New().String()
	h.ResourceId = h.Id
	h.SyncAt = ftime.Now().Timestamp()

	if err := s.save(ctx, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (s *service) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)

	if req.Keywords != "" {
		query.Where("r.name LIKE ?", "%"+req.Keywords+"%")
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

	set := host.NewHostSet()
	for rows.Next() {
		ins := host.NewDefaultHost()
		err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
			&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name, &ins.Description,
			&ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
			&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.DescribeHash, &ins.ResourceHash, &ins.ResourceId,
			&ins.CPU, &ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
			&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
			&ins.KeyPairName, &ins.SecurityGroups,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
		}
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

	oldRH, oldDH := ins.ResourceHash, ins.DescribeHash

	switch req.UpdateMode {
	case host.PATCH:
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

	if oldRH != ins.ResourceHash {
		// 避免SQL注入, 请使用Prepare
		stmt, err = tx.Prepare(updateResourceSQL)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(
			ins.ExpireAt, ins.Category, ins.Type, ins.Name, ins.Description,
			ins.Status, ins.UpdateAt, ins.SyncAt, ins.SyncAccount,
			ins.PublicIP, ins.PrivateIP, ins.PayType, ins.DescribeHash, ins.ResourceHash,
			ins.ResourceId,
		)
		if err != nil {
			return nil, err
		}
	} else {
		s.log.Debug("resource data hash not changed, needn't update")
	}

	if oldDH != ins.DescribeHash {
		// 避免SQL注入, 请使用Prepare
		stmt, err = tx.Prepare(updateHostSQL)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(
			ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec, ins.OSType, ins.OSName,
			ins.ImageID, ins.InternetMaxBandwidthOut,
			ins.InternetMaxBandwidthIn, ins.KeyPairName, ins.SecurityGroups,
			ins.Id,
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
	querySQL, args := query.Where("id = ?", req.Id).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query host error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := host.NewDefaultHost()
	err = queryStmt.QueryRow(args...).Scan(
		&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
		&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name, &ins.Description,
		&ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
		&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.DescribeHash, &ins.ResourceHash, &ins.ResourceId,
		&ins.CPU, &ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
		&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
		&ins.KeyPairName, &ins.SecurityGroups,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("describe host error, %s", err.Error())
	}

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
