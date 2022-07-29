package dns

import (
	dom "github.com/alibabacloud-go/domain-20180129/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/dns"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

func (o *DnsOperator) QueryDomain(req *provider.QueryDomainRequest) pager.Pager {
	p := newDomainPager(o)
	p.SetRate(req.Rate)
	return p
}

// 分页查询自己账户下的域名列表
// 参考: https://next.api.aliyun.com/api/Domain/2018-01-29/QueryDomainList?params={}
func (o *DnsOperator) queryDomain(req *dom.QueryDomainListRequest) (*dns.DomainSet, error) {
	set := dns.NewDomainSet()

	resp, err := o.client.QueryDomainList(req)
	if err != nil {
		return nil, err
	}

	set.Total = int64(tea.Int32Value(resp.Body.TotalItemNum))
	set.Items = o.transferDomainSet(resp.Body.Data).Items
	return set, nil
}

func (o *DnsOperator) transferDomainSet(items *dom.QueryDomainListResponseBodyData) *dns.DomainSet {
	set := dns.NewDomainSet()
	for i := range items.Domain {
		set.Add(o.transferDomain(items.Domain[i]))
	}
	return set
}

func (o *DnsOperator) transferDomain(ins *dom.QueryDomainListResponseBodyDataDomain) *dns.Domain {
	r := dns.NewDefaultDomain()

	b := r.Base
	b.Vendor = resource.VENDOR_ALIYUN
	b.CreateAt = tea.Int64Value(ins.RegistrationDateLong) / 1000
	b.Id = tea.StringValue(ins.InstanceId)

	info := r.Information
	info.ExpireAt = tea.Int64Value(ins.ExpirationDateLong) / 1000
	info.Name = tea.StringValue(ins.DomainName)
	info.Type = tea.StringValue(ins.RegistrantType)
	info.Status = tea.StringValue(ins.DomainStatus)
	info.Category = tea.StringValue(ins.ProductId)
	return r
}
