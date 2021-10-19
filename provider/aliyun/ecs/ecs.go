package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/infraboard/cmdb/pkg/host"
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperater(client *ecs.Client) *EcsOperater {
	return &EcsOperater{
		client: client,
		log:    zap.L().Named("ALI ECS"),
	}
}

type EcsOperater struct {
	client *ecs.Client
	log    logger.Logger
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
	h.Base.Vendor = resource.VendorAliYun
	h.Base.Region = ins.RegionId
	h.Base.Zone = ins.ZoneId
	h.Base.CreateAt = ins.CreationTime
	h.Base.InstanceId = ins.InstanceId

	h.Information.ExpireAt = ins.ExpiredTime
	h.Information.Type = ins.InstanceType
	h.Information.Name = ins.InstanceName
	h.Information.Description = ins.Description
	h.Information.Status = ins.Status
	h.Information.Tags = o.transferTags(ins.Tags.Tag)
	h.Information.PublicIP = ins.PublicIpAddress.IpAddress
	h.Information.PrivateIP = ins.InnerIpAddress.IpAddress
	h.Information.PayType = ins.InstanceChargeType

	h.Describe.CPU = int64(ins.CPU)
	h.Describe.Memory = int64(ins.Memory)
	h.Describe.GPUAmount = ins.GPUAmount
	h.Describe.GPUSpec = ins.GPUSpec
	h.Describe.OSType = ins.OsType
	h.Describe.OSName = ins.OSName
	h.Describe.SerialNumber = ins.SerialNumber
	h.Describe.ImageID = ins.ImageId
	h.Describe.InternetMaxBandwidthOut = int64(ins.InternetMaxBandwidthOut)
	h.Describe.InternetMaxBandwidthIn = int64(ins.InternetMaxBandwidthIn)
	h.Describe.KeyPairName = []string{ins.KeyPairName}
	h.Describe.SecurityGroups = ins.SecurityGroupIds.SecurityGroupId
	return h
}

func (o *EcsOperater) transferTags(tags []ecs.Tag) map[string]string {
	return nil
}
