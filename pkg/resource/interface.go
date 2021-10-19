package resource

import "context"

type Service interface {
	Search(context.Context, *SearchRequest) (*ResourceSet, error)
}
