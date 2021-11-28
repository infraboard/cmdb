package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/app/host"
	"github.com/infraboard/cmdb/app/resource"
)

type Ec2Operater struct {
	client *ec2.Client
	log    logger.Logger
}

// NewEc2Operator Ec2Operator
func NewEc2Operator(client *ec2.Client) *Ec2Operater {
	return &Ec2Operater{
		client: client,
		log:    zap.L().Named("AWS EC2"),
	}
}

func (o *Ec2Operater) transferSet(items []types.Reservation) *host.HostSet {
	set := host.NewHostSet()
	for _, item := range items {
		for i, _ := range item.Instances {
			set.Add(o.transferOne(item.Instances[i]))
		}
	}
	return set
}

func (o *Ec2Operater) transferOne(ins types.Instance) *host.Host {
	h := host.NewDefaultHost()
	h.Base.Vendor = resource.Vendor_AMAZON
	h.Base.Zone = *ins.Placement.AvailabilityZone
	h.Base.InstanceId = *ins.InstanceId
	h.Base.CreateAt = ins.LaunchTime.Unix()
	h.Information.Type = string(ins.InstanceType)
	// 判断tags中是否有NAME字段(实例名称), 不存在则取实例ID.
	if ParseTagName(ins.Tags) == "" {
		h.Information.Name = *ins.InstanceId
	} else {
		h.Information.Name = ParseTagName(ins.Tags)
	}
	h.Information.Status = string(ins.State.Name)
	h.Information.Tags = ParseTag(ins.Tags)
	h.Information.PublicIp = []string{*ins.PublicIpAddress}
	h.Information.PrivateIp = []string{*ins.PrivateIpAddress}
	h.Information.PayType = string(ins.Placement.Tenancy)
	h.Describe.Cpu = int64((*ins.CpuOptions.ThreadsPerCore) * (*ins.CpuOptions.CoreCount))
	h.Describe.OsName = *ins.PlatformDetails
	h.Describe.ImageId = *ins.ImageId
	h.Describe.KeyPairName = []string{*ins.KeyName}
	h.Describe.SecurityGroups = ParseGroup(ins.SecurityGroups)
	return h
}

func ParseTagName(items []types.Tag) string {
	for i := range items {
		if *items[i].Key == "Name" {
			return *items[i].Value
		}
	}
	return ""
}

func ParseGroup(items []types.GroupIdentifier) []string {
	var sg []string
	for i := range items {
		sg = append(sg, *items[i].GroupName)
	}
	return sg
}

func ParseTag(items []types.Tag) map[string]string {
	return nil
}
