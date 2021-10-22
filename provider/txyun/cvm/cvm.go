package cvm

import (
	"time"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCVMOperater(client *cvm.Client) *CVMOperater {
	return &CVMOperater{
		client: client,
		log:    zap.L().Named("Tx CVM"),
	}
}

type CVMOperater struct {
	client *cvm.Client
	log    logger.Logger
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
	h.Base.Vendor = resource.VendorTencent
	h.Base.Region = o.client.GetRegion()
	h.Base.Zone = utils.PtrStrV(ins.Placement.Zone)
	h.Base.CreateAt = o.parseTime(utils.PtrStrV(ins.CreatedTime))
	h.Base.InstanceId = utils.PtrStrV(ins.InstanceId)

	h.Information.ExpireAt = o.parseTime(utils.PtrStrV(ins.ExpiredTime))
	h.Information.Type = utils.PtrStrV(ins.InstanceType)
	h.Information.Name = utils.PtrStrV(ins.InstanceName)
	h.Information.Status = utils.PtrStrV(ins.InstanceState)
	h.Information.Tags = transferTags(ins.Tags)
	h.Information.PublicIP = utils.SlicePtrStrv(ins.PublicIpAddresses)
	h.Information.PrivateIP = utils.SlicePtrStrv(ins.PrivateIpAddresses)
	h.Information.PayType = utils.PtrStrV(ins.InstanceChargeType)

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

func (o *CVMOperater) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
