package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
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

func (h *handler) DescribeTask(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := task.NewDescribeTaskRequestWithId(ctx.PS.ByName("id"))
	ins, err := h.task.DescribeTask(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeTaskRecord(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := task.NewQueryTaskRecordRequest(ctx.PS.ByName("id"))
	ins, err := h.task.QueryTaskRecord(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
