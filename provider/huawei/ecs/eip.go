package ecs

import (
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/apps/eip"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
)

func (o *EcsOperator) PageQueryEip(req *provider.QueryEipRequest) pager.Pager {
	p := newEipPager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询弹性公网IP列表
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=EIP&api=ListPublicips&version=v2
func (o *EcsOperator) queryEip(req *model.ListPublicipsRequest) (*eip.EIPSet, error) {
	set := eip.NewEIPSet()

	resp, err := o.eip.ListPublicips(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("####")
	fmt.Println(resp.String())

	set.Items = o.transferEipSet(resp).Items

	if set.Length() > 0 {
		req.Marker = &set.Items[0].Base.Id
	}

	return set, nil
}

func (o *EcsOperator) transferEipSet(list *model.ListPublicipsResponse) *eip.EIPSet {
	set := eip.NewEIPSet()

	if list.Publicips == nil {
		return set
	}

	items := *list.Publicips

	for i := range items {
		set.Add(o.transferEip(items[i]))
	}
	return set
}

func (o *EcsOperator) transferEip(ins model.PublicipShowResp) *eip.EIP {
	h := eip.NewDefaultEip()

	b := h.Base
	b.Vendor = resource.VENDOR_HUAWEI
	b.Id = tea.StringValue(ins.Id)
	b.Region = tea.StringValue(ins.Profile.RegionId)
	b.CreateAt = utils.ParseDefaultSecondTime(ins.CreateTime.String())

	info := h.Information
	sd, _ := ins.Status.MarshalJSON()
	info.Status = string(sd)
	info.Name = tea.StringValue(ins.Alias)
	bt, _ := ins.BandwidthShareType.MarshalJSON()
	info.Category = string(bt)
	info.Type = tea.StringValue(ins.Type)
	info.PublicIp = []string{tea.StringValue(ins.PublicIpAddress)}
	info.PrivateIp = []string{tea.StringValue(ins.PrivateIpAddress)}
	info.SyncAccount = tea.StringValue(ins.Profile.UserId)

	desc := h.Describe
	desc.BandWidth = int64(tea.Int32Value(ins.BandwidthSize))
	return h
}
