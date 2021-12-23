package rds

import (
	"time"

	hw_rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"

	"github.com/infraboard/cmdb/app/rds"
	"github.com/infraboard/cmdb/app/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperater(client *hw_rds.RdsClient) *RdsOperater {
	return &RdsOperater{
		client: client,
		log:    zap.L().Named("Huawei Rds"),
	}
}

type RdsOperater struct {
	client *hw_rds.RdsClient
	log    logger.Logger
}

func (o *RdsOperater) transferSet(list *[]model.InstanceResponse) *rds.Set {
	set := rds.NewSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *RdsOperater) transferOne(ins model.InstanceResponse) *rds.RDS {
	h := rds.NewDefaultRDS()
	h.Base.Vendor = resource.Vendor_HUAWEI
	h.Base.Region = ins.Region
	h.Base.CreateAt = o.parseTime(ins.Created)
	h.Base.InstanceId = ins.Id

	h.Information.Category = ins.Type
	h.Information.Name = ins.Name
	h.Information.Description = utils.PtrStrV(ins.Alias)
	h.Information.Status = ins.Status
	h.Information.Tags = o.transferTags(ins.Tags)
	h.Information.PrivateIp, h.Information.PublicIp = ins.PrivateIps, ins.PublicIps

	return h
}

func (o *RdsOperater) parseTime(t string) int64 {
	if t == "" {
		return 0
	}

	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

func (o *RdsOperater) transferTags(tags []model.TagResponse) map[string]string {
	return nil
}
