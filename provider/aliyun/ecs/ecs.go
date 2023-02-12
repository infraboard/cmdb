package ecs

import (
	"context"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
)

func (o *EcsOperator) DescribeHost(ctx context.Context, req *provider.DescribeRequest) (*host.Host, error) {
	r := &ecs.DescribeInstancesRequest{
		RegionId:   o.client.RegionId,
		PageNumber: tea.Int32(1),
		PageSize:   tea.Int32(1),
	}
	r.InstanceIds = tea.String(`["` + req.Id + `"]`)
	hs, err := o.queryInstance(r)
	if err != nil {
		return nil, err
	}
	if hs.Length() == 0 {
		return nil, exception.NewNotFound("instance %s not found", req.Id)
	}

	return hs.Items[0], nil
}

func (o *EcsOperator) PageQueryHost(req *provider.QueryRequest) pager.Pager {
	p := newEcsPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询一台或多台ECS实例的详细信息
// 参考文档: https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeInstances?params={}
func (o *EcsOperator) queryInstance(req *ecs.DescribeInstancesRequest) (*host.HostSet, error) {
	set := host.NewHostSet()

	resp, err := o.client.DescribeInstances(req)
	if err != nil {
		return nil, err
	}

	if resp.Body != nil && resp.Body.Instances != nil && resp.Body.Instances.Instance != nil {
		set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
		set.Items = o.transferInstanceSet(resp.Body.Instances.Instance).Items
		o.log.Debugf("get %d host, page number[%d], page size[%d] total[%d]",
			set.Length(),
			*req.PageNumber,
			*req.PageSize,
			set.Total)
	}

	return set, nil
}

func (o *EcsOperator) transferInstanceSet(items []*ecs.DescribeInstancesResponseBodyInstancesInstance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferInstance(items[i]))
	}
	return set
}

func (o *EcsOperator) transferInstance(ins *ecs.DescribeInstancesResponseBodyInstancesInstance) *host.Host {
	h := host.NewDefaultHost()
	b := h.Resource.Meta

	b.CreateAt = utils.ParseDefaultMiniteTime(tea.StringValue(ins.CreationTime))
	b.Id = tea.StringValue(ins.InstanceId)
	b.SerialNumber = tea.StringValue(ins.SerialNumber)

	i := h.Resource.Spec
	i.ExpireAt = utils.ParseDefaultMiniteTime(tea.StringValue(ins.ExpiredTime))
	i.Type = tea.StringValue(ins.InstanceType)
	i.Name = tea.StringValue(ins.InstanceName)
	i.Description = tea.StringValue(ins.Description)
	i.Vendor = resource.VENDOR_ALIYUN
	i.Region = tea.StringValue(ins.RegionId)
	i.Zone = tea.StringValue(ins.ZoneId)
	i.Owner = o.GetAccountId()

	h.Resource.Status.PublicAddress = tea.StringSliceValue(ins.PublicIpAddress.IpAddress)
	h.Resource.Status.PrivateAddress = o.parsePrivateIp(ins)

	h.Resource.Status.Phase = praseEcsStatus(ins.Status)
	h.Resource.Cost.PayMode = mapping.PrasePAY_MODE(ins.InstanceChargeType)

	i.Cpu = tea.Int32Value(ins.Cpu)
	i.Memory = tea.Int32Value(ins.Memory)
	i.Gpu = tea.Int32Value(ins.GPUAmount)

	if ins.EipAddress != nil {
		h.Resource.Status.PublicAddress = []string{tea.StringValue(ins.EipAddress.IpAddress)}
		i.BandWidth = tea.Int32Value(ins.EipAddress.Bandwidth)
	}

	h.Resource.Spec.Tags = o.transferTags(ins.Tags)

	h.Describe.GpuSpec = tea.StringValue(ins.GPUSpec)
	h.Describe.OsType = tea.StringValue(ins.OSType)
	h.Describe.OsName = tea.StringValue(ins.OSName)

	h.Describe.ImageId = tea.StringValue(ins.ImageId)
	h.Describe.InternetMaxBandwidthOut = int64(tea.Int32Value(ins.InternetMaxBandwidthOut))
	h.Describe.InternetMaxBandwidthIn = int64(tea.Int32Value(ins.InternetMaxBandwidthIn))
	h.Describe.KeyPairName = []string{tea.StringValue(ins.KeyPairName)}
	h.Describe.SecurityGroups = tea.StringSliceValue(ins.SecurityGroupIds.SecurityGroupId)
	return h
}

func (o *EcsOperator) parsePrivateIp(ins *ecs.DescribeInstancesResponseBodyInstancesInstance) []string {
	ips := []string{}

	if ins.NetworkInterfaces == nil {
		return ips
	}

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
		ret = append(ret, resource.NewGroupTag(
			tea.StringValue(tags.Tag[i].TagKey),
			tea.StringValue(tags.Tag[i].TagValue),
		))
	}
	return
}
