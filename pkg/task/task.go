package task

import "github.com/infraboard/cmdb/pkg/resource"

func NewTask() *Task {
	return &Task{
		Details: []*Detail{},
	}
}

// 同个区域的同一种资源一次只能有1个task run
type Task struct {
	ID           string        `json:"id"`            // 任务id
	Region       string        `json:"region"`        // 同步的区域
	ResourceType resource.Type `json:"resource_type"` // 同步的资源
	SecretID     string        `json:"secret_id"`     // 关联secret
	StartAt      int64         `json:"start_at"`      // 开始同步的时间
	EndAt        int64         `json:"end_at"`        // 同步结束时间
	TotolSucceed int64         `json:"total_succeed"` // 成功的条数
	TotalFailed  int64         `json:"total_failed"`  // 失败的条数
	Details      []*Detail     `json:"details"`       // 同步详情
}

type Detail struct {
	Name      string `json:"name"`       // 资源名称
	IsSuccess bool   `json:"is_success"` // 是否同步成功
	Message   string `json:"message"`    // 同步失败原因
}

func (s *Task) AddFailed(name, message string) {
	s.Details = append(s.Details, &Detail{
		IsSuccess: false,
		Name:      name,
		Message:   message,
	})
	s.TotalFailed++
}

func (s *Task) AddSucceed(name, message string) {
	s.Details = append(s.Details, &Detail{
		IsSuccess: true,
		Name:      name,
		Message:   message,
	})
	s.TotolSucceed++
}
