package dns

import (
	"context"
	"fmt"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/infraboard/cmdb/apps/dns"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/pager"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func (o *DnsOperator) QueryRecord(req *provider.QueryRecordRequest) pager.Pager {
	p := newRecordPager(o)
	p.SetRate(float64(req.Rate))
	p.req.Domain = &req.Domain
	return p
}

// 获取某个域名下的解析记录
// 参考: https://console.cloud.tencent.com/api/explorer?Product=dnspod&Version=2021-03-23&Action=DescribeRecordList&SignVersion=
func (o *DnsOperator) queryRecord(ctx context.Context, req *dnspod.DescribeRecordListRequest) (*dns.RecordSet, error) {
	resp, err := o.client.DescribeRecordListWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return o.transferRecordSet(resp.Response.RecordList), nil
}

func (o *DnsOperator) transferRecordSet(items []*dnspod.RecordListItem) *dns.RecordSet {
	set := dns.NewRecordSet()
	for i := range items {
		set.Add(o.transferRecord(items[i]))
	}
	return set
}

func (o *DnsOperator) transferRecord(ins *dnspod.RecordListItem) *dns.Record {
	r := dns.NewDefaultRecord()
	r.Id = fmt.Sprintf("%d", tea.Uint64Value(ins.RecordId))
	r.Status = tea.StringValue(ins.Status)
	r.Value = tea.StringValue(ins.Value)
	r.UpdatedOn = utils.ParseSecondMod1Time(tea.StringValue(ins.UpdatedOn))
	r.Name = tea.StringValue(ins.Name)
	r.Line = tea.StringValue(ins.Line)
	r.Type = tea.StringValue(ins.Type)
	r.Remark = tea.StringValue(ins.Remark)
	r.Ttl = int64(tea.Uint64Value(ins.TTL))
	r.Mx = int64(tea.Uint64Value(ins.MX))
	return r
}
