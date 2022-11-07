package ecs

import (
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
	"github.com/infraboard/mcube/pager"

	"github.com/infraboard/cmdb/apps/eip"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
)

func (o *EcsOperator) PageQueryEip(req *provider.QueryRequest) pager.Pager {
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

	set.Items = o.transferEipSet(resp).Items

	if set.Length() > 0 {
		req.Marker = &set.Items[0].Resource.Meta.Id
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
	r := eip.NewDefaultEip()

	b := r.Resource.Meta
	b.Id = tea.StringValue(ins.Id)
	b.CreateAt = utils.ParseDefaultSecondTime(ins.CreateTime.String())

	info := r.Resource.Spec
	info.Vendor = resource.VENDOR_HUAWEI
	info.Region = tea.StringValue(ins.Profile.RegionId)
	info.Owner = tea.StringValue(ins.Profile.UserId)
	sd, _ := ins.Status.MarshalJSON()
	r.Resource.Status.Phase = string(sd)
	info.Name = tea.StringValue(ins.Alias)
	bt, _ := ins.BandwidthShareType.MarshalJSON()
	info.Category = string(bt)
	info.Type = tea.StringValue(ins.Type)

	status, err := ins.Status.MarshalJSON()
	if err == nil {
		r.Resource.Status.Phase = praseEIPStatus(strings.Trim(strings.TrimSpace(string(status)), `"`))
	} else {
		o.log.Errorf("status marshal json error, %s", err)
	}

	r.Resource.Status.PublicAddress = []string{tea.StringValue(ins.PublicIpAddress)}
	r.Resource.Status.PrivateAddress = []string{tea.StringValue(ins.PrivateIpAddress)}

	desc := r.Describe
	desc.BandWidth = int64(tea.Int32Value(ins.BandwidthSize))
	return r
}
