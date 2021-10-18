package syncer

import "github.com/infraboard/cmdb/pkg/host"

const (
	CrendentialAPIKey CrendentialType = iota
	CrendentialPassword
)

type CrendentialType int

type Secret struct {
	Id          string `json:"id"`          // 全局唯一Id
	Description string `json:"description"` // 描述
	CreateAt    string `json:"create_at"`   // 创建时间

	Vendor          host.Vendor     `json:"vendor"`
	Region          string          `json:"region"`
	CrendentialType CrendentialType `json:"crendential_type"`
	APIKey          string          `json:"api_key"`
	APISecret       string          `json:"api_secret"`
}
