package ec2

import (
	"context"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *Ec2operator) PageQueryHost(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}

// 参考文档: https://docs.aws.amazon.com/zh_cn/AWSEC2/latest/APIReference/API_DescribeInstances.html
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
	b := h.Resource.Meta
	b.Id = *ins.InstanceId
	b.CreateAt = ins.LaunchTime.Unix()

	i := h.Resource.Spec
	i.Vendor = resource.VENDOR_AMAZON
	i.Zone = *ins.Placement.AvailabilityZone
	i.Type = string(ins.InstanceType)
	// h.Spec.PayType = string(ins.Placement.Tenancy)
	i.Cpu = tea.Int32Value(ins.CpuOptions.ThreadsPerCore) * tea.Int32Value(ins.CpuOptions.CoreCount)

	// 判断tags中是否有NAME字段(实例名称), 不存在则取实例ID.
	if ParseTagName(ins.Tags) == "" {
		h.Resource.Spec.Name = tea.StringValue(ins.InstanceId)
	} else {
		h.Resource.Spec.Name = ParseTagName(ins.Tags)
	}

	h.Resource.Status.Phase = string(ins.State.Name)
	h.Resource.Status.PublicAddress = []string{tea.StringValue(ins.PublicIpAddress)}
	h.Resource.Status.PrivateAddress = []string{tea.StringValue(ins.PrivateIpAddress)}

	h.Describe.OsName = *ins.PlatformDetails
	h.Describe.ImageId = *ins.ImageId
	h.Describe.KeyPairName = []string{tea.StringValue(ins.KeyName)}
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

func (o *Ec2operator) QueryDisk(req *provider.QueryRequest) pager.Pager {
	panic("not imple")
}

func (o *Ec2operator) DescribeHost(ctx context.Context, req *provider.DescribeRequest) (*host.Host, error) {
	return nil, fmt.Errorf("not impl")
}
