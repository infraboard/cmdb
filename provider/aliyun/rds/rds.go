package rds

import (
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRdsOperater(client *rds.Client) *RdsOperater {
	return &RdsOperater{
		client: client,
		log:    zap.L().Named("ALI RDS"),
	}
}

type RdsOperater struct {
	client *rds.Client
	log    logger.Logger
}

func (o *RdsOperater) transferSet(items []rds.DBInstanceAttribute) *cmdbRds.Set {
	set := cmdbRds.NewSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *RdsOperater) transferOne(ins rds.DBInstanceAttribute) *cmdbRds.RDS {
	r := cmdbRds.NewDefaultRDS()

	b := r.Base
	b.Vendor = resource.Vendor_ALIYUN
	b.Region = ins.RegionId
	b.Zone = ins.ZoneId
	b.CreateAt = o.parseTime(ins.CreationTime)
	b.InstanceId = ins.DBInstanceId

	info := r.Information
	info.ExpireAt = o.parseTime(ins.ExpireTime)
	info.Name = ins.DBInstanceDescription
	info.Type = ins.DBInstanceType
	info.Description = ins.DBInstanceDescription
	info.Status = ins.DBInstanceStatus
	info.PayType = ins.PayType
	info.Category = ins.Category

	desc := r.Describe
	desc.EngineType = ins.Engine
	desc.EngineVersion = ins.EngineVersion
	desc.InstanceClass = ins.DBInstanceClass
	desc.ClassType = ins.DBInstanceClass
	desc.ExportType = ins.DBInstanceNetType
	desc.NetworkType = ins.InstanceNetworkType
	desc.Type = ins.DBInstanceType

	cpu, _ := strconv.Atoi(ins.DBInstanceCPU)
	desc.Cpu = int32(cpu)
	desc.Memory = ins.DBInstanceMemory
	desc.DbMaxQuantity = int64(ins.DBMaxQuantity)
	desc.AccountMaxQuantity = int64(ins.AccountMaxQuantity)
	desc.MaxConnections = int64(ins.MaxConnections)
	desc.MaxIops = int64(ins.MaxIOPS)
	desc.Collation = ins.Collation
	desc.TimeZone = ins.TimeZone
	desc.StorageCapacity = int64(ins.DBInstanceStorage)
	desc.StorageType = ins.DBInstanceStorageType
	desc.SecurityIpMode = ins.SecurityIPMode
	desc.SecurityIpList = strings.Split(ins.SecurityIPList, ",")
	desc.ConnectionMode = ins.ConnectionMode
	desc.IpType = ins.IPType
	desc.LockMode = ins.LockMode
	desc.LockReason = ins.LockReason
	port, _ := strconv.Atoi(ins.Port)
	desc.Port = int64(port)
	return r
}

func (o *RdsOperater) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
