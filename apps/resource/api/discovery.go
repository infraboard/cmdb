package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) DiscoveryPrometheus(r *restful.Request, w *restful.Response) {
	query, err := resource.NewSearchRequestFromHTTP(r.Request)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("new request error, %s", err))
		return
	}

	// 补充必须的标签
	query.AddTag(resource.NewPrometheusScrapeTag())
	query.WithTags = true

	set, err := h.service.Search(r.Request.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set.PrometheusFormat())
}
