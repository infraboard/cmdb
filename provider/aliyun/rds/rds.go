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
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

// 查询RDS实例列表
// 参考文档: https://next.api.aliyun.com/api/Rds/2014-08-15/DescribeDBInstances?params={}&lang=GO
func (o *RdsOperator) Query(req *rds.DescribeDBInstancesRequest) (*cmdbRds.Set, error) {
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

func (o *RdsOperator) QueryRds(req *provider.QueryRdsRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

func (o *RdsOperator) DescribeRds(ctx context.Context, req *provider.DescribeRdsRequest) (*cmdbRds.Rds, error) {
	descReq := &rds.DescribeDBInstanceAttributeRequest{
		DBInstanceId: &req.Id,
	}

	detail, err := o.client.DescribeDBInstanceAttribute(descReq)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(detail.Body.Items)
	if set.Length() == 0 {
		return nil, exception.NewNotFound("ins %s not found", req.Id)
	}

	return set.Items[0], nil
}

func (o *RdsOperator) transferSet(items *rds.DescribeDBInstanceAttributeResponseBodyItems) *cmdbRds.Set {
	set := cmdbRds.NewSet()
	for i := range items.DBInstanceAttribute {
		set.Add(o.transferOne(items.DBInstanceAttribute[i]))
	}
	return set
}

func (o *RdsOperator) transferOne(ins *rds.DescribeDBInstanceAttributeResponseBodyItemsDBInstanceAttribute) *cmdbRds.Rds {
	r := cmdbRds.NewDefaultRDS()

	b := r.Base
	b.Vendor = resource.VENDOR_ALIYUN
	b.Region = tea.StringValue(ins.RegionId)
	b.Zone = tea.StringValue(ins.ZoneId)
	b.CreateAt = o.parseTime(tea.StringValue(ins.CreationTime))
	b.Id = tea.StringValue(ins.DBInstanceId)

	info := r.Information
	info.ExpireAt = o.parseTime(tea.StringValue(ins.ExpireTime))
	info.Name = tea.StringValue(ins.DBInstanceDescription)
	info.Type = tea.StringValue(ins.DBInstanceType)
	info.Description = tea.StringValue(ins.DBInstanceDescription)
	info.Status = tea.StringValue(ins.DBInstanceStatus)
	info.PayType = tea.StringValue(ins.PayType)
	info.Category = tea.StringValue(ins.Category)

	desc := r.Describe
	desc.EngineType = tea.StringValue(ins.Engine)
	desc.EngineVersion = tea.StringValue(ins.EngineVersion)
	desc.InstanceClass = tea.StringValue(ins.DBInstanceClass)
	desc.ClassType = tea.StringValue(ins.DBInstanceClass)
	desc.ExportType = tea.StringValue(ins.DBInstanceNetType)
	desc.NetworkType = tea.StringValue(ins.InstanceNetworkType)
	desc.Type = tea.StringValue(ins.DBInstanceType)

	cpu, _ := strconv.Atoi(tea.StringValue(ins.DBInstanceCPU))
	desc.Cpu = int32(cpu)
	desc.Memory = tea.Int64Value(ins.DBInstanceMemory)
	desc.DbMaxQuantity = int64(tea.Int32Value(ins.DBMaxQuantity))
	desc.AccountMaxQuantity = int64(tea.Int32Value(ins.AccountMaxQuantity))
	desc.MaxConnections = int64(tea.Int32Value(ins.MaxConnections))
	desc.MaxIops = int64(tea.Int32Value(ins.MaxIOPS))
	desc.Collation = tea.StringValue(ins.Collation)
	desc.TimeZone = tea.StringValue(ins.TimeZone)
	desc.StorageCapacity = int64(tea.Int32Value(ins.DBInstanceStorage))
	desc.StorageType = tea.StringValue(ins.DBInstanceStorageType)
	desc.SecurityIpMode = tea.StringValue(ins.SecurityIPMode)
	desc.SecurityIpList = strings.Split(tea.StringValue(ins.SecurityIPList), ",")
	desc.ConnectionMode = tea.StringValue(ins.ConnectionMode)
	desc.IpType = tea.StringValue(ins.IPType)
	desc.LockMode = tea.StringValue(ins.LockMode)
	desc.LockReason = tea.StringValue(ins.LockReason)
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
