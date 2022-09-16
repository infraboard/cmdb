package cvm

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pager"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func (o *CVMOperator) PageQueryHost(req *provider.QueryRequest) pager.Pager {
	p := newPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查看实例列表
// 查看实例列表: https://console.cloud.tencent.com/api/explorer?Product=cvm&Version=2017-03-12&Action=DescribeInstances&SignVersion=
func (o *CVMOperator) QueryCVM(ctx context.Context, req *cvm.DescribeInstancesRequest) (*host.HostSet, error) {
	resp, err := o.client.DescribeInstancesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Response.InstanceSet)
	set.Total = utils.PtrInt64(resp.Response.TotalCount)
	return set, nil
}

func (o *CVMOperator) DescribeHost(ctx context.Context, req *provider.DescribeRequest) (*host.Host, error) {
	query := cvm.NewDescribeInstancesRequest()
	query.InstanceIds = []*string{tea.String(req.Id)}
	query.Limit = tea.Int64(1)
	hs, err := o.QueryCVM(ctx, query)
	if err != nil {
		return nil, err
	}
	if hs.Length() == 0 {
		return nil, exception.NewNotFound("instance %s not found", err)
	}

	return hs.Items[0], nil
}

func (o *CVMOperator) transferSet(items []*cvm.Instance) *host.HostSet {
	set := host.NewHostSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CVMOperator) transferOne(ins *cvm.Instance) *host.Host {
	r := host.NewDefaultHost()
	b := r.Resource.Meta
	b.CreateAt = utils.ParseDefaultSecondTime(utils.PtrStrV(ins.CreatedTime))
	b.Id = utils.PtrStrV(ins.InstanceId)
	b.SerialNumber = utils.PtrStrV(ins.Uuid)

	i := r.Resource.Spec
	i.Vendor = resource.VENDOR_TENCENT
	i.Region = o.client.GetRegion()
	i.Zone = utils.PtrStrV(ins.Placement.Zone)
	i.Owner = o.GetAccountId()
	i.ExpireAt = utils.ParseDefaultSecondTime(utils.PtrStrV(ins.ExpiredTime))
	i.Type = utils.PtrStrV(ins.InstanceType)
	i.Name = utils.PtrStrV(ins.InstanceName)

	if ins.InternetAccessible != nil {
		i.BandWidth = int32(tea.Int64Value(ins.InternetAccessible.InternetMaxBandwidthOut))
	}
	i.Cpu = int32(utils.PtrInt64(ins.CPU))
	i.Memory = int32(utils.PtrInt64(ins.Memory))

	r.Resource.Cost.PayMode = mapping.PrasePayMode(tea.StringValue(ins.InstanceChargeType))
	r.Resource.Status.PublicIp = utils.SlicePtrStrv(ins.PublicIpAddresses)
	r.Resource.Status.PrivateIp = utils.SlicePtrStrv(ins.PrivateIpAddresses)
	r.Resource.Status.Phase = praseCvmStatus(ins.InstanceState)

	r.Resource.Tags = transferTags(ins.Tags)

	r.Describe.OsName = utils.PtrStrV(ins.OsName)

	r.Describe.ImageId = utils.PtrStrV(ins.ImageId)
	if ins.InternetAccessible != nil {
		r.Describe.InternetMaxBandwidthOut = utils.PtrInt64(ins.InternetAccessible.InternetMaxBandwidthOut)
	}
	r.Describe.KeyPairName = utils.SlicePtrStrv(ins.LoginSettings.KeyIds)
	r.Describe.SecurityGroups = utils.SlicePtrStrv(ins.SecurityGroupIds)
	return r
}

func transferTags(tags []*cvm.Tag) (ret []*resource.Tag) {
	for i := range tags {
		ret = append(ret, resource.NewGroupTag(
			utils.PtrStrV(tags[i].Key),
			utils.PtrStrV(tags[i].Value)),
		)
	}
	return
}
