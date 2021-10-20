package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	cmdbRds "github.com/infraboard/cmdb/pkg/rds"
	"github.com/infraboard/cmdb/pkg/resource"
)

func NewRdsOperater(client *rds.Client) *RdsOperater {
	return &RdsOperater{
		client: client,
	}
}

type RdsOperater struct {
	client *rds.Client
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
	r.Base.CreateAt = ins.CreateTime
	r.Base.InstanceId = ins.DBInstanceId

	r.Information.ExpireAt = ins.ExpireTime
	r.Information.Type = ins.DBInstanceType
	r.Information.Description = ins.DBInstanceDescription
	r.Information.Status = ins.DBInstanceStatus
	r.Information.PayType = ins.PayType

	return r
}
