package rds

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	hw_rds "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRdsOperator(client *hw_rds.RdsClient) *RdsOperator {
	return &RdsOperator{
		client: client,
		log:    zap.L().Named("Huawei Rds"),
	}
}

type RdsOperator struct {
	client *hw_rds.RdsClient
	log    logger.Logger
}

func (o *RdsOperator) transferSet(list *[]model.InstanceResponse) *rds.Set {
	set := rds.NewSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *RdsOperator) transferOne(ins model.InstanceResponse) *rds.RDS {
	h := rds.NewDefaultRDS()
	b := h.Base
	b.Vendor = resource.Vendor_HUAWEI
	b.Region = ins.Region
	b.CreateAt = o.parseTime(ins.Created)
	b.Id = ins.Id

	i := h.Information
	i.ExpireAt = o.parseTime(utils.PtrStrV(ins.ExpirationTime))
	i.Category = ins.Type
	i.Name = ins.Name
	i.Description = utils.PtrStrV(ins.Alias)
	i.Status = ins.Status
	i.Tags = o.transferTags(ins.Tags)
	i.PrivateIp, i.PublicIp = ins.PrivateIps, ins.PublicIps
	i.PayType = o.getEnumValue(ins.ChargeInfo.ChargeMode)
	i.Category = ins.FlavorRef

	d := h.Describe
	cpu, _ := strconv.ParseInt(utils.PtrStrV(ins.Cpu), 10, 32)
	mem, _ := strconv.ParseInt(utils.PtrStrV(ins.Mem), 10, 64)

	d.EngineType = o.getEnumValue(ins.Datastore.Type)
	d.EngineVersion = ins.Datastore.Version
	d.Cpu = int32(cpu)
	d.Memory = mem * 1024
	d.TimeZone = ins.TimeZone
	d.MaxIops = utils.PtrInt64(ins.MaxIops)

	d.StorageType = o.getEnumValue(ins.Volume.Type)
	d.StorageCapacity = int64(ins.Volume.Size)
	d.Port = int64(ins.Port)
	return h
}

func (o *RdsOperator) parseTime(t string) int64 {
	if t == "" {
		return 0
	}

	ts, err := time.Parse("2006-01-02T15:04:05+0000", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

func (o *RdsOperator) transferTags(tags []model.TagResponse) []*resource.Tag {
	return nil
}

func (o *RdsOperator) getEnumValue(m json.Marshaler) string {
	vb, err := m.MarshalJSON()
	if err != nil {
		o.log.Errorf("marshal enum error, %s", err)
		return ""
	}

	new := []byte{}
	new = bytes.ReplaceAll(vb, []byte("\""), []byte(""))
	new = bytes.ReplaceAll(new, []byte("\n"), []byte(""))
	return string(new)
}
