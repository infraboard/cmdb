package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/cmdb/apps/task"
)

func (h *handler) CreatTask(r *restful.Request, w *restful.Response) {
	req := task.NewCreateTaskRequst()
	if err := request.GetDataFromRequest(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.task.CreatTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryTask(r *restful.Request, w *restful.Response) {
	query := task.NewQueryTaskRequestFromHTTP(r.Request)
	set, err := h.task.QueryTask(r.Request.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeTask(r *restful.Request, w *restful.Response) {
	req := task.NewDescribeTaskRequestWithId(r.PathParameter("id"))
	ins, err := h.task.DescribeTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeTaskRecord(r *restful.Request, w *restful.Response) {
	req := task.NewQueryTaskRecordRequest(r.PathParameter("id"))
	ins, err := h.task.QueryTaskRecord(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
