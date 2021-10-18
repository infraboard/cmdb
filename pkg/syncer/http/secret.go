package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/cmdb/pkg/syncer"
)

func (h *handler) QuerySecret(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := syncer.NewQuerySecretRequest()
	set, err := h.service.QuerySecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateSecret(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := syncer.NewCreateSecretRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateSecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
