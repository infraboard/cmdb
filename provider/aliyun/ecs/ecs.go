package ecs

import (
	"time"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperator(client *ecs.Client) *EcsOperator {
	return &EcsOperator{
		client:        client,
		log:           zap.L().Named("ALI ECS"),
		AccountGetter: &resource.AccountGetter{},
	}
}

// https://next.api.aliyun.com/api/Ecs/2014-05-26/CreateInstance?lang=GO&params={}
type EcsOperator struct {
	client *ecs.Client
	log    logger.Logger
	*resource.AccountGetter
}

func (o *EcsOperator) transferSet(items []*ecs.DescribeInstancesResponseBodyInstancesInstance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *EcsOperator) transferOne(ins *ecs.DescribeInstancesResponseBodyInstancesInstance) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_ALIYUN
	h.Base.Region = tea.StringValue(ins.RegionId)
	h.Base.Zone = tea.StringValue(ins.ZoneId)

	h.Base.CreateAt = o.parseTime(tea.StringValue(ins.CreationTime))
	h.Base.Id = tea.StringValue(ins.InstanceId)

	h.Information.ExpireAt = o.parseTime(tea.StringValue(ins.ExpiredTime))
	h.Information.Type = tea.StringValue(ins.InstanceType)
	h.Information.Name = tea.StringValue(ins.InstanceName)
	h.Information.Description = tea.StringValue(ins.Description)
	h.Information.Status = tea.StringValue(ins.Status)
	h.Information.Tags = o.transferTags(ins.Tags)
	h.Information.PublicIp = tea.StringSliceValue(ins.PublicIpAddress.IpAddress)
	h.Information.PrivateIp = o.parsePrivateIp(ins)
	h.Information.PayType = tea.StringValue(ins.InstanceChargeType)
	h.Information.SyncAccount = o.GetAccountId()

	h.Describe.Cpu = int64(tea.Int32Value(ins.Cpu))
	h.Describe.Memory = int64(tea.Int32Value(ins.Memory))
	h.Describe.GpuAmount = tea.Int32Value(ins.GPUAmount)
	h.Describe.GpuSpec = tea.StringValue(ins.GPUSpec)
	h.Describe.OsType = tea.StringValue(ins.OSType)
	h.Describe.OsName = tea.StringValue(ins.OSName)
	h.Describe.SerialNumber = tea.StringValue(ins.SerialNumber)
	h.Describe.ImageId = tea.StringValue(ins.ImageId)
	h.Describe.InternetMaxBandwidthOut = int64(tea.Int32Value(ins.InternetMaxBandwidthOut))
	h.Describe.InternetMaxBandwidthIn = int64(tea.Int32Value(ins.InternetMaxBandwidthIn))
	h.Describe.KeyPairName = []string{tea.StringValue(ins.KeyPairName)}
	h.Describe.SecurityGroups = tea.StringSliceValue(ins.SecurityGroupIds.SecurityGroupId)
	return h
}

func (o *EcsOperator) parsePrivateIp(ins *ecs.DescribeInstancesResponseBodyInstancesInstance) []string {
	ips := []string{}
	// 优先通过网卡查询主私网IP地址
	for _, nc := range ins.NetworkInterfaces.NetworkInterface {
		for _, ip := range nc.PrivateIpSets.PrivateIpSet {
			if tea.BoolValue(ip.Primary) {
				ips = append(ips, tea.StringValue(ip.PrivateIpAddress))
			}
		}
	}
	if len(ips) > 0 {
		return ips
	}

	// 查询InnerIpAddress属性
	if len(ins.InnerIpAddress.IpAddress) > 0 {
		return tea.StringSliceValue(ins.InnerIpAddress.IpAddress)
	}

	// 通过专有网络VPC属性查询内网Ip
	return tea.StringSliceValue(ins.VpcAttributes.PrivateIpAddress.IpAddress)
}

func (o *EcsOperator) transferTags(tags *ecs.DescribeInstancesResponseBodyInstancesInstanceTags) (ret []*resource.Tag) {
	if tags == nil {
		return nil
	}

	for i := range tags.Tag {
		ret = append(ret, resource.NewThirdTag(
			tea.StringValue(tags.Tag[i].TagKey),
			tea.StringValue(tags.Tag[i].TagValue),
		))
	}
	return
}

func (o *EcsOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02T15:04Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}
