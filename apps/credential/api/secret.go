package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/cmdb/apps/credential"
)

func (h *handler) QuerySecret(r *restful.Request, w *restful.Response) {
	req := credential.NewQuerySecretRequestFromHTTP(r.Request)
	req.WithNamespace(r.Attribute("token").(*token.Token))
	set, err := h.service.QuerySecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateSecret(r *restful.Request, w *restful.Response) {
	req := credential.NewCreateSecretRequest()
	req.SetOwner(r.Attribute("token").(*token.Token))
	if err := request.GetDataFromRequest(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.CreateSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeSecret(r *restful.Request, w *restful.Response) {
	req := credential.NewDescribeSecretRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins.Data.Desense()
	response.Success(w, ins)
}

func (h *handler) DeleteSecret(r *restful.Request, w *restful.Response) {
	req := credential.NewDeleteSecretRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
