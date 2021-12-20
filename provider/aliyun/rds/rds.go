package rds

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/app/rds"
	"github.com/infraboard/cmdb/app/resource"
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

func (o *RdsOperater) transferSet(items []rds.DBInstance) *cmdbRds.Set {
	set := cmdbRds.NewSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *RdsOperater) transferOne(ins rds.DBInstance) *cmdbRds.RDS {
	r := cmdbRds.NewDefaultRDS()

	b := r.Base
	b.Vendor = resource.Vendor_ALIYUN
	b.Region = ins.RegionId
	b.Zone = ins.ZoneId
	b.CreateAt = o.parseTime(ins.CreateTime)
	b.InstanceId = ins.DBInstanceId

	info := r.Information
	info.ExpireAt = o.parseTime(ins.ExpireTime)
	info.Type = ins.DBInstanceType
	info.Description = ins.DBInstanceDescription
	info.Status = ins.DBInstanceStatus
	info.PayType = ins.PayType

	desc := r.Describe
	desc.Category = ins.Category
	desc.EngineType = ins.Engine
	desc.EngineVersion = ins.EngineVersion
	desc.InstanceClass = ins.DBInstanceClass
	desc.ClassType = ins.DBInstanceClass
	desc.ExportType = ins.DBInstanceNetType
	desc.NetworkType = ins.InstanceNetworkType
	desc.Type = ins.DBInstanceType
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
