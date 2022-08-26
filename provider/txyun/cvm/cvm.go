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
	h.Resource.Base.Vendor = resource.VENDOR_TENCENT
	h.Resource.Base.Region = o.client.GetRegion()
	h.Resource.Base.Zone = utils.PtrStrV(ins.Placement.Zone)
	h.Resource.Base.CreateAt = utils.ParseDefaultSecondTime(utils.PtrStrV(ins.CreatedTime))
	h.Resource.Base.Id = utils.PtrStrV(ins.InstanceId)

	h.Resource.Information.ExpireAt = utils.ParseDefaultSecondTime(utils.PtrStrV(ins.ExpiredTime))
	h.Resource.Information.Type = utils.PtrStrV(ins.InstanceType)
	h.Resource.Information.Name = utils.PtrStrV(ins.InstanceName)
	h.Resource.Information.Status = praseCvmStatus(ins.InstanceState)
	h.Resource.Information.PublicIp = utils.SlicePtrStrv(ins.PublicIpAddresses)
	h.Resource.Information.PrivateIp = utils.SlicePtrStrv(ins.PrivateIpAddresses)
	h.Resource.Information.PayMode = mapping.PrasePayMode(tea.StringValue(ins.InstanceChargeType))
	h.Resource.Information.Owner = o.GetAccountId()

	h.Resource.Tags = transferTags(ins.Tags)

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
