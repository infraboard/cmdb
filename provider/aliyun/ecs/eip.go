package ecs

import (
	"strconv"

	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/eip"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/aliyun/mapping"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
)

func (o *EcsOperator) PageQueryEip(req *provider.QueryEipRequest) pager.Pager {
	p := newEipPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询指定地域已创建的EIP
// 参考文档: https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeInstances?params={}
func (o *EcsOperator) queryEip(req *ecs.DescribeEipAddressesRequest) (*eip.EIPSet, error) {
	set := eip.NewEIPSet()

	resp, err := o.client.DescribeEipAddresses(req)
	if err != nil {
		return nil, err
	}

	if resp.Body != nil && resp.Body.EipAddresses != nil {
		set.Total = int64(tea.Int32Value(resp.Body.TotalCount))
		set.Items = o.transferEipSet(resp.Body.EipAddresses).Items
		o.log.Debugf("get %d eip, page number[%d], page size[%d] total[%d]",
			set.Length(),
			*req.PageNumber,
			*req.PageSize,
			set.Total)
	}

	return set, nil
}

func (o *EcsOperator) transferEipSet(items *ecs.DescribeEipAddressesResponseBodyEipAddresses) *eip.EIPSet {
	set := eip.NewEIPSet()
	for i := range items.EipAddress {
		set.Add(o.transferEip(items.EipAddress[i]))
	}
	return set
}

func (o *EcsOperator) transferEip(ins *ecs.DescribeEipAddressesResponseBodyEipAddressesEipAddress) *eip.EIP {
	h := eip.NewDefaultEip()
	base := h.Resource.Meta

	base.CreateAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.AllocationTime))
	base.Id = tea.StringValue(ins.AllocationId)

	info := h.Resource.Spec
	info.Vendor = resource.VENDOR_ALIYUN
	info.Region = tea.StringValue(ins.RegionId)
	info.Owner = o.GetAccountId()
	info.ExpireAt = utils.ParseDefaultSecondTime(tea.StringValue(ins.ExpiredTime))
	info.Type = tea.StringValue(ins.InstanceType)

	h.Resource.Status.PublicIp = []string{tea.StringValue(ins.IpAddress)}
	h.Resource.Status.Phase = praseEcsStatus(ins.Status)
	h.Resource.Cost.PayMode = mapping.PrasePayMode(ins.ChargeType)

	h.Describe.BandWidth, _ = strconv.ParseInt(tea.StringValue(ins.Bandwidth), 10, 64)
	h.Describe.InstanceId = tea.StringValue(ins.InstanceId)
	h.Describe.InstanceType = tea.StringValue(ins.InstanceType)
	return h
}
