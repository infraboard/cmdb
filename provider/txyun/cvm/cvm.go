package cvm

import (
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/utils"
)

func NewCVMOperater(client *cvm.Client) *CVMOperater {
	return &CVMOperater{
		client: client,
	}
}

type CVMOperater struct {
	client *cvm.Client
}

func (o *CVMOperater) transferSet(items []*cvm.Instance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CVMOperater) transferOne(ins *cvm.Instance) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = host.Tencent
	h.Base.Region = o.client.GetRegion()
	h.Base.Zone = utils.PtrStrV(ins.Placement.Zone)
	h.Base.CreateAt = utils.PtrStrV(ins.CreatedTime)
	h.Base.InstanceId = utils.PtrStrV(ins.InstanceId)

	h.Resource.ExpireAt = utils.PtrStrV(ins.ExpiredTime)
	h.Resource.Type = utils.PtrStrV(ins.InstanceType)
	h.Resource.Name = utils.PtrStrV(ins.InstanceName)
	h.Resource.Status = utils.PtrStrV(ins.InstanceState)
	h.Resource.Tags = transferTags(ins.Tags)
	h.Resource.PublicIP = utils.SlicePtrStrv(ins.PublicIpAddresses)
	h.Resource.PrivateIP = utils.SlicePtrStrv(ins.PrivateIpAddresses)
	h.Resource.PayType = utils.PtrStrV(ins.InstanceChargeType)

	h.Describe.CPU = utils.PtrInt64(ins.CPU)
	h.Describe.Memory = utils.PtrInt64(ins.Memory)
	h.Describe.OSName = utils.PtrStrV(ins.OsName)
	h.Describe.SerialNumber = utils.PtrStrV(ins.Uuid)
	h.Describe.ImageID = utils.PtrStrV(ins.ImageId)
	if ins.InternetAccessible != nil {
		h.Describe.InternetMaxBandwidthOut = utils.PtrInt64(ins.InternetAccessible.InternetMaxBandwidthOut)
	}
	h.Describe.KeyPairName = utils.SlicePtrStrv(ins.LoginSettings.KeyIds)
	h.Describe.SecurityGroups = utils.SlicePtrStrv(ins.SecurityGroupIds)
	return h
}

func transferTags(tags []*cvm.Tag) map[string]string {
	return nil
}
