package ecs

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2/model"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
)

func (o *EcsOperator) DescribeDisk(ctx context.Context, r *provider.DescribeRequest) (
	*disk.Disk, error) {
	if err := r.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	req := &model.ListVolumesRequest{
		Id:    tea.String(r.Id),
		Limit: tea.Int32(1),
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

func (o *EcsOperator) PageQueryDisk(req *provider.QueryDiskRequest) pager.Pager {
	p := newDiskPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询所有云硬盘详情
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=EVS&api=ListVolumes
func (o *EcsOperator) QueryDisk(req *model.ListVolumesRequest) (*disk.DiskSet, error) {
	set := disk.NewDiskSet()

	resp, err := o.evs.ListVolumes(req)
	if err != nil {
		return nil, err
	}

	set.Items = o.transferDiskSet(resp).Items

	return set, nil
}

func (o *EcsOperator) transferDiskSet(list *model.ListVolumesResponse) *disk.DiskSet {
	set := disk.NewDiskSet()

	if list.Volumes == nil {
		return set
	}

	items := *list.Volumes

	for i := range items {
		set.Add(o.transferDisk(items[i]))
	}
	return set
}

func (o *EcsOperator) transferDisk(ins model.VolumeDetail) *disk.Disk {
	h := disk.NewDefaultDisk()

	b := h.Base
	b.Vendor = resource.VENDOR_HUAWEI
	b.Zone = ins.AvailabilityZone
	b.Id = ins.Id
	b.CreateAt = utils.ParseTime("2006-01-02T15:04:05.999999", ins.CreatedAt)

	info := h.Information
	info.Name = ins.Name
	info.Description = ins.Description
	info.Status = praseDiskStatus(ins.Status)
	info.Category = ins.ServiceType
	info.Type = ins.VolumeType

	desc := h.Describe
	if ins.Bootable == "true" {
		desc.Type = "system"
	} else {
		desc.Type = "data"
	}
	if len(ins.Attachments) > 0 {
		desc.AttachedTime = utils.ParseTime("2006-01-02T15:04:05.999999", ins.Attachments[0].AttachedAt)
		desc.Device = ins.Attachments[0].Device
		desc.InstanceId = ins.Attachments[0].ServerId
	}

	desc.MultiAttach = ins.Multiattach
	desc.Size = uint64(ins.Size)
	desc.Encrypted = tea.BoolValue(ins.Encrypted)
	return h
}
