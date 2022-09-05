package cdb

import (
	"context"
	"fmt"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/infraboard/cmdb/apps/rds"
	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *CDBOperator) DescribeRds(ctx context.Context, r *provider.DescribeRequest) (
	*cmdbRds.Rds, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := cdb.NewDescribeDBInstancesRequest()
	req.InstanceIds = tea.StringSlice([]string{r.Id})
	req.Limit = tea.Uint64(1)

	set, err := o.Query(ctx, req)
	if err != nil {
		return nil, err
	}
	if set.Length() == 0 {
		return nil, exception.NewNotFound("rds %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *CDBOperator) PageQueryRds(req *provider.QueryRdsRequest) pager.Pager {
	return newPager(20, o)
}

// 查询实例列表 (CDB)
// 参考: https://console.cloud.tencent.com/api/explorer?Product=cdb&Version=2017-03-20&Action=DescribeDBInstances&SignVersion=
func (o *CDBOperator) Query(ctx context.Context, req *cdb.DescribeDBInstancesRequest) (*rds.Set, error) {
	resp, err := o.client.DescribeDBInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response.Items), nil
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

	b := r.Resource.Meta

	b.CreateAt = o.parseTime(utils.PtrStrV(ins.CreateTime))
	b.Id = utils.PtrStrV(ins.InstanceId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_TENCENT
	info.Region = utils.PtrStrV(ins.Region)
	info.Zone = utils.PtrStrV(ins.Zone)
	info.ExpireAt = o.parseTime(utils.PtrStrV(ins.DeadlineTime))
	info.Name = utils.PtrStrV(ins.InstanceName)
	info.Category = utils.PtrStrV(ins.DeviceType)
	r.Resource.Status.Phase = praseStatus(ins.Status)
	r.Resource.Cost.PayMode = mapping.PrasePayMode(fmt.Sprintf("%d", tea.Int64Value(ins.PayType)))
	info.Cpu = int32(utils.PtrInt64(ins.Cpu))
	info.Memory = int32(utils.PtrInt64(ins.Memory))
	info.Storage = int32(utils.PtrInt64(ins.Volume))

	// 补充其他状态
	if ins.TaskStatus != nil && *ins.TaskStatus != 0 {
		r.Resource.Status.Phase = praseTaskStatus(ins.TaskStatus)
	}

	desc := r.Describe
	desc.EngineType = "MySQL"
	desc.EngineVersion = utils.PtrStrV(ins.EngineVersion)
	desc.Type = o.ParseType(ins.InstanceType)

	desc.Port = utils.PtrInt64(ins.WanPort)

	return r
}

func (o *CDBOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixMilli()
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

// 付费类型，可能的返回值：0-包年包月；1-包年包月
func (o *CDBOperator) ParsePayMode(id *int64) string {
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
