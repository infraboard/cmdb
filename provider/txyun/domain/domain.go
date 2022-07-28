package domain

import (
	"context"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/domain"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func (o *DnsOperator) QueryDomain(req *provider.QueryDomainRequest) pager.Pager {
	p := newDomainPager(o)
	p.SetRate(float64(req.Rate))
	return p
}

// 获取域名列表
// 参考: https://console.cloud.tencent.com/api/explorer?Product=dnspod&Version=2021-03-23&Action=DescribeDomainList&SignVersion=
func (o *DnsOperator) queryDomain(ctx context.Context, req *dnspod.DescribeDomainListRequest) (*domain.DomainSet, error) {
	resp, err := o.client.DescribeDomainListWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferDomainSet(resp.Response.DomainList), nil
}

func (o *DnsOperator) transferDomainSet(items []*dnspod.DomainListItem) *domain.DomainSet {
	set := domain.NewDomainSet()
	for i := range items {
		set.Add(o.transferDomain(items[i]))
	}
	return set
}

func (o *DnsOperator) transferDomain(ins *dnspod.DomainListItem) *domain.Domain {
	r := domain.NewDefaultDomain()

	b := r.Base
	b.Vendor = resource.VENDOR_TENCENT
	b.CreateAt = utils.ParseSecondMod1Time(tea.StringValue(ins.CreatedOn))
	b.Id = fmt.Sprintf("%d", ins.DomainId)

	info := r.Information
	info.ExpireAt = utils.ParseSecondMod1Time(tea.StringValue(ins.VipEndAt))
	info.Name = tea.StringValue(ins.Name)
	info.Type = tea.StringValue(ins.Grade)
	info.Status = tea.StringValue(ins.Status)
	info.Category = tea.StringValue(ins.GradeTitle)
	info.Description = tea.StringValue(ins.Remark)

	return r
}
