package resource

import (
	"context"
	"net/http"
	"strconv"
)

type Service interface {
	Search(context.Context, *SearchRequest) (*ResourceSet, error)
}

func NewSearchRequestFromHTTP(r *http.Request) *SearchRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &SearchRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}
