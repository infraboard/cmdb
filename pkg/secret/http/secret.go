package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/cmdb/pkg/secret"
)

func (h *handler) QuerySecret(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := secret.NewQuerySecretRequestFromHTTP(r)
	set, err := h.service.QuerySecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateSecret(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := secret.NewCreateSecretRequest()
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

func (h *handler) DescribeSecret(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := secret.NewDescribeSecretRequest(ps.ByName("id"))
	ins, err := h.service.DescribeSecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins.Desense()
	response.Success(w, ins)
}

func (h *handler) DeleteSecret(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := secret.NewDeleteSecretRequestWithID(ps.ByName("id"))
	set, err := h.service.DeleteSecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
