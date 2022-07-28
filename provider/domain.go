package provider

import "github.com/infraboard/mcube/pager"

type DomainOperator interface {
	QueryDomain(req *QueryDomainRequest) pager.Pager
}

func NewQueryDomainRequest() *QueryDomainRequest {
	return &QueryDomainRequest{
		Rate: 5,
	}
}

type QueryDomainRequest struct {
	Rate float64 `json:"rate"`
}
