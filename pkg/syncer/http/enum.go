package http

import (
	"net/http"

	"github.com/infraboard/cmdb/pkg/syncer"
	"github.com/infraboard/cmdb/utils"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) ListCrendentialType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := []utils.EnumDescribe{
		{Value: syncer.CrendentialAPIKey.String(), Describe: "API凭证"},
		{Value: syncer.CrendentialPassword.String(), Describe: "用户名密码"},
	}
	response.Success(w, resp)
}
