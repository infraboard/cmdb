package http

import (
	"net/http"

	"github.com/infraboard/cmdb/pkg/secret"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) ListCrendentialType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := []utils.EnumDescribe{
		{Value: secret.CrendentialType_CrendentialAPIKey.String(), Describe: "API凭证"},
		{Value: secret.CrendentialType_CrendentialPassword.String(), Describe: "用户名密码"},
	}
	response.Success(w, resp)
}