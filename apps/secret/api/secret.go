package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/restful/response"

	"github.com/infraboard/cmdb/apps/secret"
)

func (h *handler) QuerySecret(r *restful.Request, w *restful.Response) {
	req := secret.NewQuerySecretRequestFromHTTP(r.Request)
	req.WithNamespace(r.Attribute("token").(*token.Token))
	set, err := h.service.QuerySecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateSecret(r *restful.Request, w *restful.Response) {
	req := secret.NewCreateSecretRequest()
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
	req := secret.NewDescribeSecretRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins.Spec.Desense()
	response.Success(w, ins)
}

func (h *handler) DeleteSecret(r *restful.Request, w *restful.Response) {
	req := secret.NewDeleteSecretRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteSecret(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
