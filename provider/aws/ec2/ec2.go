package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *Ec2operator) PageQueryHost(req *provider.QueryHostRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}

func (o *Ec2operator) Query(ctx context.Context, req *ec2.DescribeInstancesInput) (*host.HostSet, error) {
	set := host.NewHostSet()
	result, err := o.client.DescribeInstances(ctx, req)
	if err != nil {
		o.log.Errorf("Got an error retrieving information about your Amazon EC2 instances: %s", err)
		return nil, err
	}
	// 设置Next Page token
	req.NextToken = result.NextToken
	set.Items = o.transferSet(result.Reservations).Items
	return set, nil
}

func (o *Ec2operator) transferSet(items []types.Reservation) *host.HostSet {
	set := host.NewHostSet()
	for _, item := range items {
		for i := range item.Instances {
			set.Add(o.transferOne(item.Instances[i]))
		}
	}
	return set
}

func (o *Ec2operator) transferOne(ins types.Instance) *host.Host {
	h := host.NewDefaultHost()
	b := h.Resource.Base
	b.Vendor = resource.VENDOR_AMAZON
	b.Zone = *ins.Placement.AvailabilityZone
	b.Id = *ins.InstanceId
	b.CreateAt = ins.LaunchTime.Unix()

	i := h.Resource.Information
	i.Type = string(ins.InstanceType)
	// h.Information.PayType = string(ins.Placement.Tenancy)
	i.Cpu = (*ins.CpuOptions.ThreadsPerCore) * (*ins.CpuOptions.CoreCount)

	// 判断tags中是否有NAME字段(实例名称), 不存在则取实例ID.
	if ParseTagName(ins.Tags) == "" {
		h.Resource.Information.Name = *ins.InstanceId
	} else {
		h.Resource.Information.Name = ParseTagName(ins.Tags)
	}

	h.Resource.Information.Status = string(ins.State.Name)
	h.Resource.Information.PublicIp = []string{*ins.PublicIpAddress}
	h.Resource.Information.PrivateIp = []string{*ins.PrivateIpAddress}

	h.Describe.OsName = *ins.PlatformDetails
	h.Describe.ImageId = *ins.ImageId
	h.Describe.KeyPairName = []string{*ins.KeyName}
	h.Describe.SecurityGroups = ParseGroup(ins.SecurityGroups)

	h.Resource.Tags = ParseTag(ins.Tags)
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

func ParseTag(items []types.Tag) []*resource.Tag {
	return nil
}

func (o *Ec2operator) QueryDisk(req *provider.QueryDiskRequest) pager.Pager {
	panic("not imple")
}

func (o *Ec2operator) DescribeHost(ctx context.Context, req *provider.DescribeHostRequest) (*host.Host, error) {
	return nil, fmt.Errorf("not impl")
}
