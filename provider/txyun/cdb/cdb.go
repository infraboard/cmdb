package cdb

import (
	"time"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/infraboard/cmdb/apps/rds"
	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCDBOperator(client *cdb.Client) *CDBOperator {
	return &CDBOperator{
		client: client,
		log:    zap.L().Named("Tx CDB"),
	}
}

type CDBOperator struct {
	client *cdb.Client
	log    logger.Logger
}

func (o *CDBOperator) transferSet(items []*cdb.InstanceInfo) *rds.Set {
	set := rds.NewSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CDBOperator) transferOne(ins *cdb.InstanceInfo) *rds.RDS {
	r := cmdbRds.NewDefaultRDS()

	b := r.Base
	b.Vendor = resource.Vendor_TENCENT
	b.Region = utils.PtrStrV(ins.Region)
	b.Zone = utils.PtrStrV(ins.Zone)
	b.CreateAt = o.parseTime(utils.PtrStrV(ins.CreateTime))
	b.Id = utils.PtrStrV(ins.InstanceId)

	info := r.Information
	info.ExpireAt = o.parseTime(utils.PtrStrV(ins.DeadlineTime))
	info.Name = utils.PtrStrV(ins.InstanceName)
	info.Category = utils.PtrStrV(ins.DeviceType)
	// info.Type = ins.DBInstanceType
	// info.Description = ins.DBInstanceDescription
	// info.Status = ins.DBInstanceStatus
	// info.PayType = ins.PayType

	desc := r.Describe
	desc.EngineType = "MySQL"
	desc.EngineVersion = utils.PtrStrV(ins.EngineVersion)
	desc.InstanceClass = o.ParseInstanceType(ins.InstanceType)
	// desc.ClassType = ins.DBInstanceClass
	// desc.ExportType = ins.DBInstanceNetType
	// desc.NetworkType = ins.InstanceNetworkType
	// desc.Type = ins.DBInstanceType
	desc.Cpu = int32(utils.PtrInt64(ins.Cpu))
	desc.Memory = utils.PtrInt64(ins.Memory)
	desc.StorageCapacity = utils.PtrInt64(ins.Volume)
	desc.Port = utils.PtrInt64(ins.WanPort)

	return r
}

func (o *CDBOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

// 实例类型，可能的返回值：1-主实例；2-灾备实例；3-只读实例
func (o *CDBOperator) ParseInstanceType(id *int64) string {
	if id == nil {
		return ""
	}

	switch *id {
	case 1:
		return "主实例"
	case 2:
		return "灾备实例"
	case 3:
		return "只读实例"
	}

	return ""
}
