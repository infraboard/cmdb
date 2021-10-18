package syncer

import "github.com/infraboard/cmdb/pkg/host"

const (
	CrendentialAPIKey CrendentialType = iota
	CrendentialPassword
)

type CrendentialType int

type Rule struct {
	Vendor          host.Vendor     `json:"vendor"`
	Region          string          `json:"region"`
	CrendentialType CrendentialType `json:"crendential_type"`
	APIKey          string          `json:"api_key"`
	APISecret       string          `json:"api_secret"`
}
