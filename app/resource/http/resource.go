package http

import (
	"net/http"

	"github.com/infraboard/cmdb/app/resource"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) SearchResource(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := resource.NewSearchRequestFromHTTP(r)
	set, err := h.service.Search(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
