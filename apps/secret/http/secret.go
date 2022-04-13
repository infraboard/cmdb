package http

import (
	"net/http"

	"github.com/infraboard/keyauth/apps/token"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/cmdb/apps/secret"
)

func (h *handler) QuerySecret(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := secret.NewQuerySecretRequestFromHTTP(r)
	req.WithNamespace(ctx.AuthInfo.(*token.Token))
	set, err := h.service.QuerySecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateSecret(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := secret.NewCreateSecretRequest()
	req.SetOwner(ctx.AuthInfo.(*token.Token))
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

func (h *handler) DescribeSecret(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)

	req := secret.NewDescribeSecretRequest(ctx.PS.ByName("id"))
	ins, err := h.service.DescribeSecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	ins.Data.Desense()
	response.Success(w, ins)
}

func (h *handler) DeleteSecret(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := secret.NewDeleteSecretRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DeleteSecret(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
