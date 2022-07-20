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

func (o *CDBOperator) transferOne(ins *cdb.InstanceInfo) *rds.Rds {
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
	info.Status = o.ParseStatus(ins.Status)
	info.PayType = o.ParsePayType(ins.PayType)

	desc := r.Describe
	desc.EngineType = "MySQL"
	desc.EngineVersion = utils.PtrStrV(ins.EngineVersion)
	// desc.ExportType = ins.DBInstanceNetType
	// desc.NetworkType = ins.InstanceNetworkType
	// desc.Type = ins.DBInstanceType
	desc.Type = o.ParseType(ins.InstanceType)
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
func (o *CDBOperator) ParseType(id *int64) string {
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

// 实例状态，可能的返回值：0-创建中；1-运行中；4-隔离中；5-已隔离
func (o *CDBOperator) ParseStatus(id *int64) string {
	if id == nil {
		return ""
	}
	switch *id {
	case 0:
		return "创建中"
	case 1:
		return "运行中"
	case 4:
		return "隔离中"
	case 5:
		return "已隔离"
	}
	return ""
}

// 付费类型，可能的返回值：0-包年包月；1-包年包月
func (o *CDBOperator) ParsePayType(id *int64) string {
	if id == nil {
		return ""
	}
	switch *id {
	case 0:
		return "包年包月"
	case 1:
		return "包年包月"
	}
	return ""
}
