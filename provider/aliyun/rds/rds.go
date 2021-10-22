package rds

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/pkg/rds"
	"github.com/infraboard/cmdb/pkg/resource"
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

func (o *RdsOperater) transferSet(items []rds.DBInstance) *cmdbRds.RdsSet {
	set := cmdbRds.NewRdsSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *RdsOperater) transferOne(ins rds.DBInstance) *cmdbRds.Rds {
	r := cmdbRds.NewDefaultRds()
	r.Base.Vendor = resource.VendorAliYun
	r.Base.Region = ins.RegionId
	r.Base.Zone = ins.ZoneId
	r.Base.CreateAt = o.parseTime(ins.CreateTime)
	r.Base.InstanceId = ins.DBInstanceId

	r.Information.ExpireAt = o.parseTime(ins.ExpireTime)
	r.Information.Type = ins.DBInstanceType
	r.Information.Description = ins.DBInstanceDescription
	r.Information.Status = ins.DBInstanceStatus
	r.Information.PayType = ins.PayType

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
