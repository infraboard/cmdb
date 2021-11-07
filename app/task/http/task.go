package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/cmdb/app/task"
)

func (h *handler) CreatTask(w http.ResponseWriter, r *http.Request) {
	req := task.NewCreateTaskRequst()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.task.CreatTask(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryTask(w http.ResponseWriter, r *http.Request) {
	query := task.NewQueryTaskRequestFromHTTP(r)
	set, err := h.task.QueryTask(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
