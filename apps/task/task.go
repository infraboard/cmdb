package task

import (
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

func NewDefaultTask() *Task {
	return &Task{
		Data:   &CreateTaskRequst{},
		Status: &Status{},
	}
}

func NewTaskFromReq(req *CreateTaskRequst) (*Task, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate sync request error, %s", err)
	}

	return &Task{
		Id:     xid.New().String(),
		Data:   req,
		Status: &Status{},
	}, nil
}

func (s *Task) Run() {
	s.Status.StartAt = ftime.Now().Timestamp()
	s.Status.Stage = Stage_RUNNING
}

func (s *Task) UpdateSecretDesc(desc string) {
	s.CredentialDescription = desc
}

func (s *Task) Completed() {
	s.Status.EndAt = ftime.Now().Timestamp()
	if s.Status.Stage != Stage_FAILED {
		if s.Status.TotalFailed == 0 {
			s.Status.Stage = Stage_SUCCESS
		} else {
			s.Status.Stage = Stage_WARNING
		}
	}
}

func (s *Task) Failed(message string) {
	s.Status.EndAt = time.Now().UnixMilli()
	s.Status.Stage = Stage_FAILED
	s.Status.Message = message
}

func (s *Task) AddDetail(d *Record) {
	if d.IsSuccess {
		s.Status.TotalSucceed++
	} else {
		s.Status.TotalFailed++
	}
}

func NewTaskSet() *TaskSet {
	return &TaskSet{
		Items: []*Task{},
	}
}

func (r *TaskSet) Add(item *Task) {
	r.Items = append(r.Items, item)
}
