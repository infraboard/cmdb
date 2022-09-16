package ecs

import (
	"context"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *EcsOperator) DescribeDisk(ctx context.Context, r *provider.DescribeRequest) (
	*disk.Disk, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := &ecs.DescribeDisksRequest{
		RegionId: o.client.RegionId,
		PageSize: tea.Int32(1),
		DiskIds:  tea.String(`["` + r.Id + `"]`),
	}

	set, err := o.QueryDisk(req)
	if err != nil {
		return nil, err
	}

	if set.Length() == 0 {
		return nil, exception.NewNotFound("disk %s not found", r.Id)
	}

	return set.Items[0], nil
}

func (o *EcsOperator) PageQueryDisk(req *provider.QueryRequest) pager.Pager {
	p := newDiskPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询一块或多块已经创建的块存储（包括云盘以及本地盘）
// 参考文档: https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeDisks?params={}&lang=GO
func (o *EcsOperator) QueryDisk(req *ecs.DescribeDisksRequest) (*disk.DiskSet, error) {
	set := disk.NewDiskSet()

	resp, err := o.client.DescribeDisks(req)
	if err != nil {
		return nil, err
	}
	req.NextToken = resp.Body.NextToken

	set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
	set.Items = o.transferDiskSet(resp.Body.Disks).Items

	return set, nil
}

func (o *EcsOperator) transferDiskSet(items *ecs.DescribeDisksResponseBodyDisks) *disk.DiskSet {
	set := disk.NewDiskSet()
	for i := range items.Disk {
		set.Add(o.transferDisk(items.Disk[i]))
	}
	return set
}

func (o *EcsOperator) transferDisk(ins *ecs.DescribeDisksResponseBodyDisksDisk) *disk.Disk {
	h := disk.NewDefaultDisk()
	h.Resource.Meta.CreateAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.CreationTime))
	h.Resource.Meta.Id = tea.StringValue(ins.DiskId)

	h.Resource.Spec.Owner = o.GetAccountId()
	h.Resource.Spec.Vendor = resource.VENDOR_ALIYUN
	h.Resource.Spec.Region = tea.StringValue(ins.RegionId)
	h.Resource.Spec.Zone = tea.StringValue(ins.ZoneId)
	h.Resource.Spec.ExpireAt = utils.ParseDefaultMiniteTime(tea.StringValue(ins.ExpiredTime))
	h.Resource.Spec.Type = tea.StringValue(ins.Type)
	h.Resource.Spec.Name = tea.StringValue(ins.DiskName)
	h.Resource.Spec.Description = tea.StringValue(ins.Description)
	h.Resource.Status.Phase = praseDiskStatus(ins.Status)
	h.Resource.Cost.PayMode = mapping.PrasePayMode(ins.DiskChargeType)

	h.Resource.Tags = o.transferDiskTags(ins.Tags)

	h.Describe.InstanceId = tea.StringValue(ins.InstanceId)
	h.Describe.Size = uint64(tea.Int32Value(ins.Size))
	h.Describe.Device = tea.StringValue(ins.Device)
	h.Describe.AttachedTime = utils.ParseDefaultSecondTime(tea.StringValue(ins.AttachedTime))
	h.Describe.Portable = tea.BoolValue(ins.Portable)
	h.Describe.DetachedTime = utils.ParseDefaultSecondTime(tea.StringValue(ins.DetachedTime))
	h.Describe.Encrypted = tea.BoolValue(ins.Encrypted)
	h.Describe.EnableAutoSnapshot = tea.BoolValue(ins.EnableAutoSnapshot)
	h.Describe.DeleteWithInstance = tea.BoolValue(ins.DeleteWithInstance)
	h.Describe.Iops = tea.Int32Value(ins.IOPS)
	return h
}

func (o *EcsOperator) transferDiskTags(tags *ecs.DescribeDisksResponseBodyDisksDiskTags) (ret []*resource.Tag) {
	if tags == nil {
		return nil
	}

	for i := range tags.Tag {
		ret = append(ret, resource.NewGroupTag(
			tea.StringValue(tags.Tag[i].TagKey),
			tea.StringValue(tags.Tag[i].TagValue),
		))
	}
	return
}
