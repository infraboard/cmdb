package task

import (
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

func NewDefaultTask() *Task {
	return &Task{
		Details: []*Detail{},
	}
}

func NewTaskFromReq(req *CreateTaskRequst) (*Task, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate sync request error, %s", err)
	}

	return &Task{
		Id:           xid.New().String(),
		Region:       req.Region,
		ResourceType: req.ResourceType,
		SecretId:     req.SecretId,
		Details:      []*Detail{},
	}, nil
}

func (s *Task) Run() {
	s.StartAt = ftime.Now().Timestamp()
	s.Status = Status_RUNNING
}

func (s *Task) UpdateSecretDesc(desc string) {
	s.SecretDescription = desc
}

func (s *Task) Completed() {
	s.EndAt = ftime.Now().Timestamp()
	if s.Status != Status_FAILED {
		if s.TotalFailed == 0 {
			s.Status = Status_SUCCESS
		} else {
			s.Status = Status_WARNING
		}
	}
}

func (s *Task) Failed(message string) {
	s.Status = Status_FAILED
	s.Message = message
}

func (s *Task) AddDetail(d *Detail) {
	if d.IsSuccess {
		s.TotalSucceed++
	} else {
		s.TotalFailed++
	}
	s.Details = append(s.Details, d)
}

func NewTaskSet() *TaskSet {
	return &TaskSet{
		Items: []*Task{},
	}
}

func (r *TaskSet) Add(item *Task) {
	r.Items = append(r.Items, item)
}
