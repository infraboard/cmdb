package dns

import (
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/pager"
)

// 根据传入参数获取指定主域名的所有解析记录列表
// 参考文档: https://next.api.aliyun.com/api/Alidns/2015-01-09/DescribeDomainRecords?params={}&tab=DEMO&lang=GO
func (o *DnsOperator) QueryRecord(req *provider.QueryRecordRequest) pager.Pager {
	return nil
}
