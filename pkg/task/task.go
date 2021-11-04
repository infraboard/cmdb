package task

import (
	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

const (
	StatusPendding Status = iota
	// 任务允许中
	StatusRunning
	// 执行成功
	StatusSuccess
	// 执行失败
	StatusFailed
	// 部分成功
	StatusWarning
)

type Status int

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

// 同个区域的同一种资源一次只能有1个task run
type Task struct {
	Id                string        `json:"id"`                 // 任务id
	Region            string        `json:"region"`             // 同步的区域
	ResourceType      resource.Type `json:"resource_type"`      // 同步的资源
	SecretId          string        `json:"secret_id"`          // 关联secret
	SecretDescription string        `json:"secret_description"` // secret描述
	Timeout           int           `json:"timeout"`            // 任务超时时间
	Status            Status        `json:"status"`             // 任务状态
	Message           string        `json:"message"`            // 失败时的异常信息
	StartAt           int64         `json:"start_at"`           // 开始同步的时间
	EndAt             int64         `json:"end_at"`             // 同步结束时间
	TotolSucceed      int64         `json:"total_succeed"`      // 成功的条数
	TotalFailed       int64         `json:"total_failed"`       // 失败的条数
	Details           []*Detail     `json:"details"`            // 同步详情
}

type Detail struct {
	Name      string `json:"name"`       // 资源名称
	IsSuccess bool   `json:"is_success"` // 是否同步成功
	Message   string `json:"message"`    // 同步失败原因
}

func (s *Task) Run() {
	s.StartAt = ftime.Now().Timestamp()
	s.Status = StatusRunning
}

func (s *Task) UpdateSecretDesc(desc string) {
	s.SecretDescription = desc
}

func (s *Task) Completed() {
	s.EndAt = ftime.Now().Timestamp()
	if s.Status != StatusFailed {
		if s.TotalFailed == 0 {
			s.Status = StatusSuccess
		} else {
			s.Status = StatusWarning
		}
	}
}

func (s *Task) Failed(message string) {
	s.Status = StatusFailed
	s.Message = message
}

func (s *Task) AddDetailFailed(name, message string) {
	s.Details = append(s.Details, &Detail{
		IsSuccess: false,
		Name:      name,
		Message:   message,
	})
	s.TotalFailed++
}

func (s *Task) AddDetailSucceed(name, message string) {
	s.Details = append(s.Details, &Detail{
		IsSuccess: true,
		Name:      name,
		Message:   message,
	})
	s.TotolSucceed++
}

func NewTaskSet() *TaskSet {
	return &TaskSet{
		Items: []*Task{},
	}
}

type TaskSet struct {
	Items []*Task `json:"items"`
	Total int     `json:"total"`
}

func (r *TaskSet) Add(item *Task) {
	r.Items = append(r.Items, item)
}
