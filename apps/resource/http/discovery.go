package http

import (
	"net/http"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) DiscoveryPrometheus(w http.ResponseWriter, r *http.Request) {
	query, err := resource.NewSearchRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("new request error, %s", err))
		return
	}

	set, err := h.service.Search(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set.PrometheusFormat())
}
