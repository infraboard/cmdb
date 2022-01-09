package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/cmdb/apps/host"
)

func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request) {
	query := host.NewQueryHostRequestFromHTTP(r)
	set, err := h.service.QueryHost(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request) {
	ins := host.NewDefaultHost()
	if err := request.GetDataFromRequest(r, ins); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.SaveHost(r.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := host.NewDescribeHostRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DescribeHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := host.NewDeleteHostRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DeleteHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PutHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := host.NewUpdateHostRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.UpdateHostData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) PatchHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := host.NewUpdateHostRequest(ctx.PS.ByName("id"))
	req.UpdateMode = host.UpdateMode_PATCH

	if err := request.GetDataFromRequest(r, req.UpdateHostData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
