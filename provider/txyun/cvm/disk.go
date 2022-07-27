package cvm

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
)

func (o *CVMOperator) QueryDisk(req *provider.QueryDiskRequest) pager.Pager {
	p := newDiskPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询云硬盘列表
// 查看实例列表: https://console.cloud.tencent.com/api/explorer?Product=cbs&Version=2017-03-12&Action=DescribeDisks&SignVersion=
func (o *CVMOperator) queryDisk(ctx context.Context, req *cbs.DescribeDisksRequest) (*disk.DiskSet, error) {
	resp, err := o.cbs.DescribeDisksWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	set := o.transferDiskSet(resp.Response)
	set.Total = int64(tea.Uint64Value(resp.Response.TotalCount))
	return set, nil
}

func (o *CVMOperator) transferDiskSet(items *cbs.DescribeDisksResponseParams) *disk.DiskSet {
	set := disk.NewDiskSet()
	for i := range items.DiskSet {
		set.Add(o.transferDisk(items.DiskSet[i]))
	}
	return set
}

func (o *CVMOperator) transferDisk(ins *cbs.Disk) *disk.Disk {
	h := disk.NewDefaultDisk()
	h.Base.Vendor = resource.VENDOR_TENCENT
	h.Base.Region = o.client.GetRegion()
	h.Base.Zone = utils.PtrStrV(ins.Placement.Zone)
	h.Base.CreateAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.CreateTime))
	h.Base.Id = utils.PtrStrV(ins.InstanceId)

	h.Information.ExpireAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.DeadlineTime))
	h.Information.Type = utils.PtrStrV(ins.DiskType)
	h.Information.Name = utils.PtrStrV(ins.DiskName)
	h.Information.Status = praseStatus(ins.DiskState)
	h.Information.PayType = utils.PtrStrV(ins.DiskChargeType)
	h.Information.SyncAccount = o.GetAccountId()

	desc := h.Describe
	desc.Type = tea.StringValue(ins.DiskUsage)
	desc.InstanceId = tea.StringValue(ins.InstanceId)
	desc.Size = tea.Uint64Value(ins.DiskSize)
	desc.DeleteWithInstance = tea.BoolValue(ins.DeleteWithInstance)
	desc.Portable = tea.BoolValue(ins.Portable)
	desc.Encrypted = tea.BoolValue(ins.Encrypt)
	desc.MultiAttach = tea.BoolValue(ins.Shareable)
	return h
}
