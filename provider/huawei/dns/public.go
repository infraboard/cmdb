package dns

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
	"github.com/infraboard/cmdb/apps/dns"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *DnsOperator) QueryDomain(req *provider.QueryDomainRequest) pager.Pager {
	p := newPrivateZonePager(o)
	p.SetRate(req.Rate)
	return p
}

// 查询公网Zone的列表
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=DNS&api=ListPublicZones
func (o *DnsOperator) queryPublicDomain(req *model.ListPublicZonesRequest) (*dns.DomainSet, error) {
	set := dns.NewDomainSet()

	resp, err := o.client.ListPublicZones(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	set.Items = o.transferPublicZoneSet(resp).Items

	return set, nil
}

func (o *DnsOperator) transferPublicZoneSet(list *model.ListPublicZonesResponse) *dns.DomainSet {
	set := dns.NewDomainSet()

	items := *list.Zones
	for i := range items {
		set.Add(o.transferPublicZone(items[i]))
	}
	return set
}

func (o *DnsOperator) transferPublicZone(ins model.PublicZoneResp) *dns.Domain {
	r := dns.NewDefaultDomain()

	b := r.Resource.Base
	b.Vendor = resource.VENDOR_HUAWEI
	return r
}
