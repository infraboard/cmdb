package cvm

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
)

func (o *CVMOperator) DescribeDisk(ctx context.Context, r *provider.DescribeRequest) (
	*disk.Disk, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := cbs.NewDescribeDisksRequest()
	req.Limit = tea.Uint64(1)
	req.DiskIds = tea.StringSlice([]string{r.Id})

	set, err := o.QueryDisk(ctx, req)
	if err != nil {
		return nil, err
	}

	if set.Length() == 0 {
		return nil, exception.NewNotFound("disk %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *CVMOperator) PageQueryDisk(req *provider.QueryRequest) pager.Pager {
	p := newDiskPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询云硬盘列表
// 查看实例列表: https://console.cloud.tencent.com/api/explorer?Product=cbs&Version=2017-03-12&Action=DescribeDisks&SignVersion=
func (o *CVMOperator) QueryDisk(ctx context.Context, req *cbs.DescribeDisksRequest) (*disk.DiskSet, error) {
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

	h.Resource.Meta.CreateAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.CreateTime))
	h.Resource.Meta.Id = utils.PtrStrV(ins.DiskId)

	h.Resource.Spec.Vendor = resource.VENDOR_TENCENT
	h.Resource.Spec.Region = o.client.GetRegion()
	h.Resource.Spec.Zone = utils.PtrStrV(ins.Placement.Zone)
	h.Resource.Spec.Owner = o.GetAccountId()
	h.Resource.Spec.ExpireAt = utils.ParseTime("2006-01-02 15:04:05", utils.PtrStrV(ins.DeadlineTime))
	h.Resource.Spec.Type = utils.PtrStrV(ins.DiskType)
	h.Resource.Spec.Name = utils.PtrStrV(ins.DiskName)
	h.Resource.Status.Phase = praseDiskStatus(ins.DiskState)
	h.Resource.Cost.PayMode = mapping.PrasePAY_MODE(tea.StringValue(ins.DiskChargeType))

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
