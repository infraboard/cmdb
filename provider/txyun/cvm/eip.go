package cvm

import (
	"context"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/eip"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

func (o *CVMOperator) PageQueryEip(req *provider.QueryRequest) pager.Pager {
	p := newEipPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询弹性公网IP列表 (VPC)
// 参考文档: https://console.cloud.tencent.com/api/explorer?Product=vpc&Version=2017-03-12&Action=DescribeAddresses&SignVersion=
func (o *CVMOperator) queryEip(ctx context.Context, req *vpc.DescribeAddressesRequest) (*eip.EIPSet, error) {
	resp, err := o.vpc.DescribeAddressesWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	set := o.transferEipSet(resp.Response)
	set.Total = int64(tea.Int64Value(resp.Response.TotalCount))
	return set, nil
}

func (o *CVMOperator) transferEipSet(items *vpc.DescribeAddressesResponseParams) *eip.EIPSet {
	set := eip.NewEIPSet()
	for i := range items.AddressSet {
		set.Add(o.transferEip(items.AddressSet[i]))
	}
	return set
}

func (o *CVMOperator) transferEip(ins *vpc.Address) *eip.EIP {
	r := eip.NewDefaultEip()
	base := r.Resource.Meta

	base.CreateAt = utils.ParseDefaultSecondTime(utils.PtrStrV(ins.CreatedTime))
	base.Id = utils.PtrStrV(ins.AddressId)

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_TENCENT
	info.Region = o.client.GetRegion()
	info.Owner = o.GetAccountId()
	info.Type = utils.PtrStrV(ins.AddressType)
	info.Name = utils.PtrStrV(ins.AddressName)

	r.Resource.Cost.PayMode = mapping.PrasePAY_MODE(utils.PtrStrV(ins.InternetChargeType))
	r.Resource.Cost.PayModeDetail = tea.StringValue(ins.InternetChargeType)

	r.Resource.Status.Phase = praseDiskStatus(ins.AddressStatus)
	r.Resource.Status.PublicIp = []string{tea.StringValue(ins.AddressIp)}
	r.Resource.Status.PrivateIp = []string{tea.StringValue(ins.PrivateAddressIp)}

	desc := r.Describe
	desc.BandWidth = int64(tea.Uint64Value(ins.Bandwidth))
	desc.InstanceId = tea.StringValue(ins.InstanceId)
	desc.Isp = tea.StringValue(ins.InternetServiceProvider)
	return r
}
