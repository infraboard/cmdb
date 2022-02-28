package http

import (
	"net/http"

	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) SearchResource(w http.ResponseWriter, r *http.Request) {
	query := resource.NewSearchRequestFromHTTP(r)
	set, err := h.service.Search(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) AddTag(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := resource.NewUpdateTagRequest(ctx.PS.ByName("id"), resource.UpdateAction_ADD)
	if err := request.GetDataFromRequest(r, &req.Tags); err != nil {
		response.Failed(w, err)
		return
	}
	set, err := h.service.UpdateTag(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) RemoveTag(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := resource.NewUpdateTagRequest(ctx.PS.ByName("id"), resource.UpdateAction_REMOVE)
	if err := request.GetDataFromRequest(r, &req.Tags); err != nil {
		response.Failed(w, err)
		return
	}
	set, err := h.service.UpdateTag(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
