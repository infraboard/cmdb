package sqlserver

import (
	"context"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	sqlserver "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver/v20180328"

	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *SQLServerOperator) PageQueryRds(req *provider.QueryRequest) pager.Pager {
	return newPager(20, o)
}

// 查询实例列表 (SQLServer)
// 参考: https://console.cloud.tencent.com/api/explorer?Product=sqlserver&Version=2018-03-28&Action=DescribeDBInstances
func (o *SQLServerOperator) Query(ctx context.Context, req *sqlserver.DescribeDBInstancesRequest) (*rds.RdsSet, error) {
	resp, err := o.client.DescribeDBInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferSet(resp.Response.DBInstances), nil
}

func (o *SQLServerOperator) transferSet(items []*sqlserver.DBInstance) *rds.RdsSet {
	set := rds.NewSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *SQLServerOperator) transferOne(ins *sqlserver.DBInstance) *rds.Rds {
	r := rds.NewDefaultRDS()

	b := r.Resource.Meta

	b.CreateAt = utils.ParseSecondMod1Time(tea.StringValue(ins.CreateTime))
	b.Id = utils.PtrStrV(ins.InstanceId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_TENCENT
	info.Region = utils.PtrStrV(ins.Region)
	info.Zone = utils.PtrStrV(ins.Zone)
	info.ExpireAt = utils.ParseSecondMod1Time(tea.StringValue(ins.EndTime))
	info.Name = utils.PtrStrV(ins.Name)
	info.Category = utils.PtrStrV(ins.Type)
	r.Resource.Status.Phase = praseStatus(ins.Status)
	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(fmt.Sprintf("%d", tea.Int64Value(ins.PayMode)))
	info.Cpu = int32(utils.PtrInt64(ins.Cpu))
	info.Memory = int32(utils.PtrInt64(ins.Memory))
	info.Storage = int32(utils.PtrInt64(ins.Storage))

	desc := r.Describe
	desc.EngineType = "SQLServer"
	desc.EngineVersion = utils.PtrStrV(ins.Version)
	desc.Type = o.ParseType(tea.StringValue(ins.InstanceType))

	desc.Port = utils.PtrInt64(ins.Vport)

	return r
}

// 实例类型，实例类型 HA-高可用 RO-只读实例 SI-基础版 BI-商业智能服务
func (o *SQLServerOperator) ParseType(id string) string {
	switch id {
	case "HA":
		return "高可用"
	case "RO":
		return "只读实例"
	case "SI":
		return "基础版"
	case "BI":
		return "商业智能服务"
	}
	return ""
}
