package rds

import (
	"context"
	"strconv"
	"strings"
	"time"

	rds "github.com/alibabacloud-go/rds-20140815/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	cmdbRds "github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *RdsOperator) DescribeRds(ctx context.Context, req *provider.DescribeRequest) (*cmdbRds.Rds, error) {
	descReq := &rds.DescribeDBInstanceAttributeRequest{
		DBInstanceId: &req.Id,
	}

	detail, err := o.client.DescribeDBInstanceAttribute(descReq)
	if err != nil {
		if v, ok := err.(*tea.SDKError); ok {
			if tea.IntValue(v.StatusCode) == 404 {
				return nil, exception.NewNotFound(tea.StringValue(v.Message)).WithData(v.Data)
			}
		}
		return nil, err
	}

	set := o.transferSet(detail.Body.Items)
	if set.Length() == 0 {
		return nil, exception.NewNotFound("ins %s not found", req.Id)
	}

	return set.Items[0], nil
}

func (o *RdsOperator) PageQueryRds(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 查询RDS实例列表
// 参考文档: https://next.api.aliyun.com/api/Rds/2014-08-15/DescribeDBInstances?params={}&lang=GO
func (o *RdsOperator) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.RdsSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}
	req.NextToken = resp.Body.NextToken

	set := cmdbRds.NewSet()
	// 更详细的数据 需要通过DescribeDBInstanceAttribute获取，比如cpu和内存信息
	for _, ins := range resp.Body.Items.DBInstance {
		descReq := &rds.DescribeDBInstanceAttributeRequest{}
		descReq.DBInstanceId = ins.DBInstanceId
		detail, err := o.client.DescribeDBInstanceAttribute(descReq)
		if err != nil {
			return nil, err
		}
		set.AddSet(o.transferSet(detail.Body.Items))
	}

	set.Total = int64(tea.Int32Value(resp.Body.TotalRecordCount))
	return set, nil
}

func (o *RdsOperator) transferSet(items *rds.DescribeDBInstanceAttributeResponseBodyItems) *cmdbRds.RdsSet {
	set := cmdbRds.NewSet()
	for i := range items.DBInstanceAttribute {
		set.Add(o.transferOne(items.DBInstanceAttribute[i]))
	}
	return set
}

func (o *RdsOperator) transferOne(ins *rds.DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) *cmdbRds.Rds {
	r := cmdbRds.NewDefaultRDS()

	b := r.Resource.Meta
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreationTime))
	b.Id = tea.StringValue(ins.DBInstanceId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_ALIYUN
	info.Region = tea.StringValue(ins.RegionId)
	info.Zone = tea.StringValue(ins.ZoneId)
	info.ExpireAt = o.parseTime(tea.StringValue(ins.ExpireTime))
	info.Name = tea.StringValue(ins.DBInstanceDescription)
	info.Type = tea.StringValue(ins.DBInstanceType)
	info.Description = tea.StringValue(ins.DBInstanceDescription)

	info.Category = tea.StringValue(ins.Category)
	cpu, _ := strconv.Atoi(tea.StringValue(ins.DBInstanceCPU))
	info.Cpu = int32(cpu)
	info.Memory = int32(tea.Int64Value(ins.DBInstanceMemory))

	info.Storage = tea.Int32Value(ins.DBInstanceStorage)

	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(ins.PayType)

	r.Resource.Status.Phase = tea.StringValue(ins.DBInstanceStatus)
	r.Resource.Status.LockMode = tea.StringValue(ins.LockMode)
	r.Resource.Status.LockReason = tea.StringValue(ins.LockReason)

	desc := r.Describe
	desc.EngineType = tea.StringValue(ins.Engine)
	desc.EngineVersion = tea.StringValue(ins.EngineVersion)
	desc.InstanceClass = tea.StringValue(ins.DBInstanceClass)
	desc.ClassType = tea.StringValue(ins.DBInstanceClass)
	desc.ExportType = tea.StringValue(ins.DBInstanceNetType)
	desc.NetworkType = tea.StringValue(ins.InstanceNetworkType)
	desc.Type = tea.StringValue(ins.DBInstanceType)

	desc.DbMaxQuantity = int64(tea.Int32Value(ins.DBMaxQuantity))
	desc.AccountMaxQuantity = int64(tea.Int32Value(ins.AccountMaxQuantity))
	desc.MaxConnections = int64(tea.Int32Value(ins.MaxConnections))
	desc.MaxIops = int64(tea.Int32Value(ins.MaxIOPS))
	desc.Collation = tea.StringValue(ins.Collation)
	desc.TimeZone = tea.StringValue(ins.TimeZone)
	desc.StorageType = tea.StringValue(ins.DBInstanceStorageType)
	desc.SecurityIpMode = tea.StringValue(ins.SecurityIPMode)
	desc.SecurityIpList = strings.Split(tea.StringValue(ins.SecurityIPList), ",")
	desc.ConnectionMode = tea.StringValue(ins.ConnectionMode)
	desc.IpType = tea.StringValue(ins.IPType)
	port, _ := strconv.Atoi(tea.StringValue(ins.Port))
	desc.Port = int64(port)
	return r
}

func (o *RdsOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
