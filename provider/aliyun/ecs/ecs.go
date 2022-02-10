package ecs

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperater(client *ecs.Client) *EcsOperater {
	return &EcsOperater{
		client:        client,
		log:           zap.L().Named("ALI ECS"),
		AccountGetter: &resource.AccountGetter{},
	}
}

type EcsOperater struct {
	client *ecs.Client
	log    logger.Logger
	*resource.AccountGetter
}

func (o *EcsOperater) transferSet(items []ecs.Instance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *EcsOperater) transferOne(ins ecs.Instance) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_ALIYUN
	h.Base.Region = ins.RegionId
	h.Base.Zone = ins.ZoneId

	h.Base.CreateAt = o.parseTime(ins.CreationTime)
	h.Base.Id = ins.InstanceId

	h.Information.ExpireAt = o.parseTime(ins.ExpiredTime)
	h.Information.Type = ins.InstanceType
	h.Information.Name = ins.InstanceName
	h.Information.Description = ins.Description
	h.Information.Status = ins.Status
	h.Information.Tags = o.transferTags(ins.Tags.Tag)
	h.Information.PublicIp = ins.PublicIpAddress.IpAddress
	h.Information.PrivateIp = ins.InnerIpAddress.IpAddress
	h.Information.PayType = ins.InstanceChargeType
	h.Information.SyncAccount = o.GetAccountId()

	h.Describe.Cpu = int64(ins.CPU)
	h.Describe.Memory = int64(ins.Memory)
	h.Describe.GpuAmount = int32(ins.GPUAmount)
	h.Describe.GpuSpec = ins.GPUSpec
	h.Describe.OsType = ins.OsType
	h.Describe.OsName = ins.OSName
	h.Describe.SerialNumber = ins.SerialNumber
	h.Describe.ImageId = ins.ImageId
	h.Describe.InternetMaxBandwidthOut = int64(ins.InternetMaxBandwidthOut)
	h.Describe.InternetMaxBandwidthIn = int64(ins.InternetMaxBandwidthIn)
	h.Describe.KeyPairName = []string{ins.KeyPairName}
	h.Describe.SecurityGroups = ins.SecurityGroupIds.SecurityGroupId
	return h
}

func (o *EcsOperater) transferTags(tags []ecs.Tag) map[string]string {
	return nil
}

func (o *EcsOperater) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
