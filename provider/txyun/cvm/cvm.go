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

func (o *CVMOperator) PageQueryHost(req *provider.QueryHostRequest) pager.Pager {
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

func (o *CVMOperator) DescribeHost(ctx context.Context, req *provider.DescribeHostRequest) (*host.Host, error) {
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
	h := host.NewDefaultHost()
	b := h.Resource.Base
	b.Vendor = resource.VENDOR_TENCENT
	b.Region = o.client.GetRegion()
	b.Zone = utils.PtrStrV(ins.Placement.Zone)
	b.CreateAt = utils.ParseDefaultSecondTime(utils.PtrStrV(ins.CreatedTime))
	b.Id = utils.PtrStrV(ins.InstanceId)

	i := h.Resource.Information
	i.ExpireAt = utils.ParseDefaultSecondTime(utils.PtrStrV(ins.ExpiredTime))
	i.Type = utils.PtrStrV(ins.InstanceType)
	i.Name = utils.PtrStrV(ins.InstanceName)
	i.Status = praseCvmStatus(ins.InstanceState)
	i.PublicIp = utils.SlicePtrStrv(ins.PublicIpAddresses)
	if ins.InternetAccessible != nil {
		i.BandWidth = int32(tea.Int64Value(ins.InternetAccessible.InternetMaxBandwidthOut))
	}
	i.PrivateIp = utils.SlicePtrStrv(ins.PrivateIpAddresses)
	i.PayMode = mapping.PrasePayMode(tea.StringValue(ins.InstanceChargeType))
	i.Owner = o.GetAccountId()
	i.Cpu = int32(utils.PtrInt64(ins.CPU))
	i.Memory = int32(utils.PtrInt64(ins.Memory))
	i.SerialNumber = utils.PtrStrV(ins.Uuid)

	h.Resource.Tags = transferTags(ins.Tags)

	h.Describe.OsName = utils.PtrStrV(ins.OsName)

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
