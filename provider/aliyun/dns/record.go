package dns

import (
	"github.com/infraboard/cmdb/apps/dns"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"

	alidns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

func (o *DnsOperator) QueryRecord(req *provider.QueryRecordRequest) pager.Pager {
	p := newRecordPager(o)
	p.req.DomainName = &req.Domain
	p.SetRate(req.Rate)
	return p
}

// 根据传入参数获取指定主域名的所有解析记录列表
// 参考文档: https://next.api.aliyun.com/api/Alidns/2015-01-09/DescribeDomainRecords?params={}&tab=DEMO&lang=GO
func (o *DnsOperator) queryRecord(req *alidns.DescribeDomainRecordsRequest) (*dns.RecordSet, error) {
	set := dns.NewRecordSet()

	resp, err := o.dns.DescribeDomainRecords(req)
	if err != nil {
		return nil, err
	}

	set.Total = tea.Int64Value(resp.Body.TotalCount)
	set.Items = o.transferRecordSet(resp.Body.DomainRecords).Items
	return set, nil
}

func (o *DnsOperator) transferRecordSet(items *alidns.DescribeDomainRecordsResponseBodyDomainRecords) *dns.RecordSet {
	set := dns.NewRecordSet()
	for i := range items.Record {
		set.Add(o.transferRecord(items.Record[i]))
	}
	return set
}

func (o *DnsOperator) transferRecord(ins *alidns.DescribeDomainRecordsResponseBodyDomainRecordsRecord) *dns.Record {
	r := dns.NewDefaultRecord()
	r.Status = tea.StringValue(ins.Status)
	r.Type = tea.StringValue(ins.Type)
	r.Ttl = tea.Int64Value(ins.TTL)
	r.Id = tea.StringValue(ins.RecordId)
	r.Name = tea.StringValue(ins.RR)
	r.Value = tea.StringValue(ins.Value)
	r.Line = tea.StringValue(ins.Line)
	return r
}
