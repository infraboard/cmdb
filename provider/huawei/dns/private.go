package dns

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
	"github.com/infraboard/cmdb/apps/dns"
	"github.com/infraboard/cmdb/apps/resource"
)

// 查询内网Zone的列表
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=DNS&api=ListPrivateZones
func (o *DnsOperator) queryPrivateDomain(req *model.ListPrivateZonesRequest) (*dns.DomainSet, error) {
	set := dns.NewDomainSet()

	resp, err := o.client.ListPrivateZones(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.String())

	set.Items = o.transferPrivateZoneSet(resp).Items

	return set, nil
}

func (o *DnsOperator) transferPrivateZoneSet(list *model.ListPrivateZonesResponse) *dns.DomainSet {
	set := dns.NewDomainSet()

	items := *list.Zones
	for i := range items {
		set.Add(o.transferPrivateZone(items[i]))
	}
	return set
}

func (o *DnsOperator) transferPrivateZone(ins model.PrivateZoneResp) *dns.Domain {
	r := dns.NewDefaultDomain()

	b := r.Base
	b.Vendor = resource.VENDOR_HUAWEI
	return r
}
