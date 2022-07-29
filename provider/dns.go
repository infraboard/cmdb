package provider

import "github.com/infraboard/mcube/pager"

type DnsOperator interface {
	QueryDomain(req *QueryDomainRequest) pager.Pager
	QueryRecord(req *QueryRecordRequest) pager.Pager
}

func NewQueryDomainRequest() *QueryDomainRequest {
	return &QueryDomainRequest{
		Rate: 5,
	}
}

type QueryDomainRequest struct {
	Rate float64 `json:"rate"`
}

func NewQueryRecordRequest(domain string) *QueryRecordRequest {
	return &QueryRecordRequest{
		Domain: domain,
		Rate:   5,
	}
}

type QueryRecordRequest struct {
	Domain string  `json:"domain"`
	Rate   float64 `json:"rate"`
}
