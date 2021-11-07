package task

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type Service interface {
	QueryTask(context.Context, *QueryTaskRequest) (*TaskSet, error)
	CreatTask(context.Context, *CreateTaskRequst) (*Task, error)
}

func NewCreateTaskRequst() *CreateTaskRequst {
	return &CreateTaskRequst{}
}

func (req *CreateTaskRequst) Validate() error {
	return validate.Struct(req)
}

func NewQueryTaskRequestFromHTTP(r *http.Request) *QueryTaskRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &QueryTaskRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

func (req *QueryTaskRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}
