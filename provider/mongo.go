package provider

import "github.com/infraboard/mcube/pager"

type MongoOperator interface {
	PageQueryMongo(req *QueryRequest) pager.Pager
}
