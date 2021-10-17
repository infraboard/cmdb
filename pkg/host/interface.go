package host

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

type Service interface {
	SaveHost(context.Context, *Host) (*Host, error)
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

func NewQueryHostRequestFromHTTP(r *http.Request) *QueryHostRequest {
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
	return &QueryHostRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

type QueryHostRequest struct {
	PageSize   uint64 `json:"page_size,omitempty"`
	PageNumber uint64 `json:"page_number,omitempty"`
	Keywords   string `json:"keywords"`
}

func (req *QueryHostRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewDescribeHostRequestWithID(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

type DescribeHostRequest struct {
	Id string `json:"id" validate:"required"`
}

func NewDeleteHostRequestWithID(id string) *DeleteHostRequest {
	return &DeleteHostRequest{Id: id}
}

type DeleteHostRequest struct {
	Id string `json:"id" validate:"required"`
}

type UpdateMode int

const (
	PUT UpdateMode = iota
	PATCH
)

func NewUpdateHostRequest(id string) *UpdateHostRequest {
	return &UpdateHostRequest{
		Id:             id,
		UpdateMode:     PUT,
		UpdateHostData: &UpdateHostData{},
	}
}

type UpdateHostRequest struct {
	Id             string          `json:"id" validate:"required"`
	UpdateMode     UpdateMode      `json:"update_mode"`
	UpdateHostData *UpdateHostData `json:"data" validate:"required"`
}

func (req *UpdateHostRequest) Validate() error {
	return validate.Struct(req)
}

type UpdateHostData struct {
	*Resource
	*Describe
}
