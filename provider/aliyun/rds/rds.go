package rds

import (
	"strconv"
	"strings"
	"time"

	rds "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRdsOperator(client *rds.Client) *RdsOperator {
	return &RdsOperator{
		client: client,
		log:    zap.L().Named("ALI RDS"),
	}
}

type RdsOperator struct {
	client *rds.Client
	log    logger.Logger
}

func (o *RdsOperator) transferSet(items *rds.DescribeDBInstanceAttributeResponseBodyItems) *cmdbRds.Set {
	set := cmdbRds.NewSet()
	for i := range items.DBInstanceAttribute {
		set.Add(o.transferOne(items.DBInstanceAttribute[i]))
	}
	return set
}

func (o *RdsOperator) transferOne(ins *rds.DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) *cmdbRds.Rds {
	r := cmdbRds.NewDefaultRDS()

	b := r.Base
	b.Vendor = resource.VENDOR_ALIYUN
	b.Region = tea.StringValue(ins.RegionId)
	b.Zone = tea.StringValue(ins.ZoneId)
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreationTime))
	b.Id = tea.StringValue(ins.DBInstanceId)

	info := r.Information
	info.ExpireAt = o.parseTime(tea.StringValue(ins.ExpireTime))
	info.Name = tea.StringValue(ins.DBInstanceDescription)
	info.Type = tea.StringValue(ins.DBInstanceType)
	info.Description = tea.StringValue(ins.DBInstanceDescription)
	info.Status = tea.StringValue(ins.DBInstanceStatus)
	info.PayType = tea.StringValue(ins.PayType)
	info.Category = tea.StringValue(ins.Category)

	desc := r.Describe
	desc.EngineType = tea.StringValue(ins.Engine)
	desc.EngineVersion = tea.StringValue(ins.EngineVersion)
	desc.InstanceClass = tea.StringValue(ins.DBInstanceClass)
	desc.ClassType = tea.StringValue(ins.DBInstanceClass)
	desc.ExportType = tea.StringValue(ins.DBInstanceNetType)
	desc.NetworkType = tea.StringValue(ins.InstanceNetworkType)
	desc.Type = tea.StringValue(ins.DBInstanceType)

	cpu, _ := strconv.Atoi(tea.StringValue(ins.DBInstanceCPU))
	desc.Cpu = int32(cpu)
	desc.Memory = tea.Int64Value(ins.DBInstanceMemory)
	desc.DbMaxQuantity = int64(tea.Int32Value(ins.DBMaxQuantity))
	desc.AccountMaxQuantity = int64(tea.Int32Value(ins.AccountMaxQuantity))
	desc.MaxConnections = int64(tea.Int32Value(ins.MaxConnections))
	desc.MaxIops = int64(tea.Int32Value(ins.MaxIOPS))
	desc.Collation = tea.StringValue(ins.Collation)
	desc.TimeZone = tea.StringValue(ins.TimeZone)
	desc.StorageCapacity = int64(tea.Int32Value(ins.DBInstanceStorage))
	desc.StorageType = tea.StringValue(ins.DBInstanceStorageType)
	desc.SecurityIpMode = tea.StringValue(ins.SecurityIPMode)
	desc.SecurityIpList = strings.Split(tea.StringValue(ins.SecurityIPList), ",")
	desc.ConnectionMode = tea.StringValue(ins.ConnectionMode)
	desc.IpType = tea.StringValue(ins.IPType)
	desc.LockMode = tea.StringValue(ins.LockMode)
	desc.LockReason = tea.StringValue(ins.LockReason)
	port, _ := strconv.Atoi(tea.StringValue(ins.Port))
	desc.Port = int64(port)
	return r
}

func (o *RdsOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
