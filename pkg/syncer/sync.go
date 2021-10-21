package syncer

import "github.com/infraboard/cmdb/pkg/resource"

func NewSyncRequest(secretId string) *SyncRequest {
	return &SyncRequest{
		SecretId: secretId,
	}
}

type SyncRequest struct {
	SecretId     string
	ResourceType resource.Type
	Region       string
}

func NewSyncReponse() *SyncReponse {
	return &SyncReponse{
		Details: []*SyncDetail{},
	}
}

type SyncReponse struct {
	TotolSucceed int64         `json:"total_succeed"`
	TotalFailed  int64         `json:"total_failed"`
	Details      []*SyncDetail `json:"details"`
}

func (s *SyncReponse) AddFailed(name, message string) {
	s.Details = append(s.Details, &SyncDetail{
		IsSuccess: false,
		Name:      name,
		Message:   message,
	})
	s.TotalFailed++
}

func (s *SyncReponse) AddSucceed(name, message string) {
	s.Details = append(s.Details, &SyncDetail{
		IsSuccess: true,
		Name:      name,
		Message:   message,
	})
	s.TotolSucceed++
}

type SyncDetail struct {
	Name      string `json:"name"`       // 资源名称
	IsSuccess bool   `json:"is_success"` // 是否同步成功
	Message   string `json:"message"`    // 同步失败原因
}
