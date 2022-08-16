package rds

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rds/v3/model"

	"github.com/infraboard/cmdb/apps/rds"
	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *RdsOperator) DescribeRds(ctx context.Context, r *provider.DescribeRequest) (
	*cmdbRds.Rds, error) {
	req := &model.ListInstancesRequest{
		Id:    tea.String(r.Id),
		Limit: tea.Int32(1),
	}
	set, err := o.Query(req)
	if err != nil {
		return nil, err
	}

	if set.Length() == 0 {
		return nil, exception.NewNotFound("rds %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *RdsOperator) PageQueryRds(req *provider.QueryRdsRequest) pager.Pager {
	return newPager(o)
}

// 查询数据库实例列表
// 参考: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=RDS&api=ListInstances
func (o *RdsOperator) Query(req *model.ListInstancesRequest) (*rds.Set, error) {
	set := rds.NewSet()

	resp, err := o.client.ListInstances(req)
	if err != nil {
		return nil, err
	}

	// 华为云返回的TotalCount不准确
	set.Total = int64(*resp.TotalCount)
	set.Items = o.transferSet(resp.Instances).Items

	return set, nil
}

func (o *RdsOperator) transferSet(list *[]model.InstanceResponse) *rds.Set {
	set := rds.NewSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *RdsOperator) transferOne(ins model.InstanceResponse) *rds.Rds {
	h := rds.NewDefaultRDS()
	b := h.Base
	b.Vendor = resource.VENDOR_HUAWEI
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
