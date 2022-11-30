package rest

import (
	"context"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/client/rest"
)

type ResourceService interface {
	Search(context.Context, *resource.SearchRequest) (*resource.ResourceSet, error)
}

type resourceImpl struct {
	client *rest.RESTClient
}

func (i *resourceImpl) Search(ctx context.Context, req *resource.SearchRequest) (
	*resource.ResourceSet, error) {
	set := resource.NewResourceSet()

	err := i.client.
		Get("resource/search").
		Do(ctx).
		Into(set)
	if err != nil {
		return nil, err
	}

	return set, nil
}
