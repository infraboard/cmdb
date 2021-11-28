package host

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

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

func (req *QueryHostRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewDescribeHostRequestWithID(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		DescribeBy: DescribeBy_HOST_ID,
		Value:      id,
	}
}

func NewDeleteHostRequestWithID(id string) *DeleteHostRequest {
	return &DeleteHostRequest{Id: id}
}

func NewUpdateHostRequest(id string) *UpdateHostRequest {
	return &UpdateHostRequest{
		Id:             id,
		UpdateMode:     UpdateMode_PUT,
		UpdateHostData: &UpdateHostData{},
	}
}

func (req *UpdateHostRequest) Validate() error {
	return validate.Struct(req)
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewHostSet(),
	}
}

type PagerResult struct {
	Data    *HostSet
	Err     error
	HasNext bool
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}
