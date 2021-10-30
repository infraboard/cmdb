package resource

import "strings"

func LoadTypeFromString(t string) Type {
	switch t {
	case "host":
		return HostResource
	case "rds":
		return RdsResource
	default:
		return Unsuport
	}
}

const (
	Unsuport Type = iota
	HostResource
	RdsResource
)

type Type int

const (
	VendorAliYun Vendor = iota
	VendorTencent
	VendorHuaWei
	VendorVsphere
)

type Vendor int

func NewDefaultResource() *Resource {
	return &Resource{
		Base:        &Base{},
		Information: &Information{},
	}
}

type Resource struct {
	*Base
	*Information
}

type Base struct {
	Id           string `json:"id"`            // 全局唯一Id
	SyncAt       int64  `json:"sync_at"`       // 同步时间
	SecretID     string `json:"secret_id"`     // 用于同步的凭证ID
	Vendor       Vendor `json:"vendor"`        // 厂商
	ResourceType Type   `json:"resource_type"` // 资源类型
	Region       string `json:"region"`        // 地域
	Zone         string `json:"zone"`          // 区域
	CreateAt     int64  `json:"create_at"`     // 创建时间
	InstanceId   string `json:"instance_id"`   // 实例ID
	ResourceHash string `json:"resource_hash"` // 基础数据Hash
	DescribeHash string `json:"describe_hash"` // 描述数据Hash
}

type Information struct {
	ExpireAt    int64             `json:"expire_at"`   // 过期时间
	Category    string            `json:"category"`    // 种类
	Type        string            `json:"type"`        // 规格
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Status      string            `json:"status"`      // 服务商中的状态
	Tags        map[string]string `json:"tags"`        // 标签
	UpdateAt    int64             `json:"update_at"`   // 更新时间
	SyncAccount string            `json:"sync_accout"` // 同步的账号
	PublicIP    []string          `json:"public_ip"`   // 公网IP
	PrivateIP   []string          `json:"private_ip"`  // 内网IP
	PayType     string            `json:"pay_type"`    // 实例付费方式
}

func (i *Information) PrivateIPToString() string {
	return strings.Join(i.PrivateIP, ",")
}

func (i *Information) PublicIPToString() string {
	return strings.Join(i.PublicIP, ",")
}

func (i *Information) LoadPrivateIPString(s string) {
	if s != "" {
		i.PrivateIP = strings.Split(s, ",")
	}
}

func (i *Information) LoadPublicIPString(s string) {
	if s != "" {
		i.PublicIP = strings.Split(s, ",")
	}
}

type SearchRequest struct {
	PageSize     uint64
	PageNumber   uint64
	Vendor       Vendor
	ResourceType Type
	Keywords     string
}

func (req *SearchRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Items: []*Resource{},
	}
}

type ResourceSet struct {
	Items []*Resource `json:"items"`
	Total int64       `json:"total"`
}

func (r *ResourceSet) Add(item *Resource) {
	r.Items = append(r.Items, item)
}
