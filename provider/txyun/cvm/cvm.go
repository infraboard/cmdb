package cvm

import (
	"time"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewCVMOperator(client *cvm.Client) *CVMOperator {
	return &CVMOperator{
		client:        client,
		log:           zap.L().Named("Tx CVM"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type CVMOperator struct {
	client *cvm.Client
	log    logger.Logger
	*resource.AccountGetter
}

func (o *CVMOperator) transferSet(items []*cvm.Instance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CVMOperator) transferOne(ins *cvm.Instance) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_TENCENT
	h.Base.Region = o.client.GetRegion()
	h.Base.Zone = utils.PtrStrV(ins.Placement.Zone)
	h.Base.CreateAt = o.parseTime(utils.PtrStrV(ins.CreatedTime))
	h.Base.Id = utils.PtrStrV(ins.InstanceId)

	h.Information.ExpireAt = o.parseTime(utils.PtrStrV(ins.ExpiredTime))
	h.Information.Type = utils.PtrStrV(ins.InstanceType)
	h.Information.Name = utils.PtrStrV(ins.InstanceName)
	h.Information.Status = praseStatus(ins.InstanceState)
	h.Information.Tags = transferTags(ins.Tags)
	h.Information.PublicIp = utils.SlicePtrStrv(ins.PublicIpAddresses)
	h.Information.PrivateIp = utils.SlicePtrStrv(ins.PrivateIpAddresses)
	h.Information.PayType = utils.PtrStrV(ins.InstanceChargeType)
	h.Information.SyncAccount = o.GetAccountId()

	h.Describe.Cpu = utils.PtrInt64(ins.CPU)
	h.Describe.Memory = utils.PtrInt64(ins.Memory)
	h.Describe.OsName = utils.PtrStrV(ins.OsName)
	h.Describe.SerialNumber = utils.PtrStrV(ins.Uuid)
	h.Describe.ImageId = utils.PtrStrV(ins.ImageId)
	if ins.InternetAccessible != nil {
		h.Describe.InternetMaxBandwidthOut = utils.PtrInt64(ins.InternetAccessible.InternetMaxBandwidthOut)
	}
	h.Describe.KeyPairName = utils.SlicePtrStrv(ins.LoginSettings.KeyIds)
	h.Describe.SecurityGroups = utils.SlicePtrStrv(ins.SecurityGroupIds)
	return h
}

func transferTags(tags []*cvm.Tag) (ret []*resource.Tag) {
	for i := range tags {
		ret = append(ret, resource.NewThirdTag(
			utils.PtrStrV(tags[i].Key),
			utils.PtrStrV(tags[i].Value)),
		)
	}
	return
}

func (o *CVMOperator) parseTime(t string) int64 {
	if t == "" {
		return 0
	}

	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
