package task

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
)

const (
	AppName = "task"
)

var (
	validate = validator.New()
)

func NewCreateTaskRequst() *CreateTaskRequst {
	return &CreateTaskRequst{
		Params: map[string]string{},
	}
}

func (req *CreateTaskRequst) Validate() error {
	return validate.Struct(req)
}

func NewQueryTaskRequestFromHTTP(r *http.Request) *QueryTaskRequest {
	qs := r.URL.Query()

	kw := qs.Get("keywords")

	return &QueryTaskRequest{
		Page:     request.NewPageRequestFromHTTP(r),
		Keywords: kw,
	}
}

func (req *QueryTaskRequest) OffSet() int64 {
	return req.Page.ComputeOffset()
}

func NewRecordSet() *RecordSet {
	return &RecordSet{
		Items: []*Record{},
	}
}

func (s *RecordSet) Add(item *Record) {
	s.Items = append(s.Items, item)
}

func NewDefaultTaskRecord() *Record {
	return &Record{}
}

func NewSyncSucceedRecord(taskId, instanceId, instanceName string) *Record {
	return &Record{
		TaskId:     taskId,
		IsSuccess:  true,
		InstanceId: instanceId,
		Name:       instanceName,
		CreateAt:   ftime.Now().Timestamp(),
	}
}

func NewSyncFailedRecord(taskId, instanceId, instanceName, message string) *Record {
	return &Record{
		TaskId:     taskId,
		IsSuccess:  false,
		CreateAt:   ftime.Now().Timestamp(),
		InstanceId: instanceId,
		Name:       instanceName,
		Message:    message,
	}
}

func NewDescribeTaskRequestWithId(id string) *DescribeTaskRequest {
	return &DescribeTaskRequest{
		Id: id,
	}
}

func NewQueryTaskRecordRequest(id string) *QueryTaskRecordRequest {
	return &QueryTaskRecordRequest{
		TaskId: id,
	}
}
