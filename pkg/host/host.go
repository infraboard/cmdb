package host

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/types/ftime"
)

const (
	PrivateIDC Vendor = iota
	Tencent
	AliYun
	HuaWei
)

type Vendor int

func NewDefaultHost() *Host {
	return &Host{
		&Base{},
		&Resource{},
		&Describe{},
	}
}

type Host struct {
	*Base
	*Resource
	*Describe
}

func (h *Host) Put(req *UpdateHostData) {
	h.Resource = req.Resource
	h.Describe = req.Describe
	h.UpdateAt = ftime.Now().Timestamp() // time, 13 时间戳
	h.GenHash()
}

func (h *Host) Patch(req *UpdateHostData) error {
	err := ObjectPatch(h.Resource, req.Resource)
	if err != nil {
		return err
	}

	err = ObjectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}

	h.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
	return nil
}

// patch JSON {a: 1, b： 2}， {b:20}  ===> {a:1, b:20}
func ObjectPatch(old, new interface{}) error {
	// {b: 20}
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}
	// {a:1, b:2}
	// {a:1, b: 20}
	return json.Unmarshal(newByte, old)
}

func (h *Host) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h.Resource)
	if err != nil {
		return err
	}
	hash.Write(b)
	h.ResourceHash = fmt.Sprintf("%x", hash.Sum(nil))

	b, err = json.Marshal(h.Describe)
	if err != nil {
		return err
	}
	hash.Reset()
	hash.Write(b)
	h.DescribeHash = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}

type Base struct {
	Id           string `json:"id"`            // 全局唯一Id
	SyncAt       int64  `json:"sync_at"`       // 同步时间
	Vendor       Vendor `json:"vendor"`        // 厂商
	Region       string `json:"region"`        // 地域
	Zone         string `json:"zone"`          // 区域
	CreateAt     int64  `json:"create_at"`     // 创建时间
	InstanceId   string `json:"instance_id"`   // 实例ID
	ResourceHash string `json:"resource_hash"` // 基础数据Hash
	DescribeHash string `json:"describe_hash"` // 描述数据Hash
}

type Resource struct {
	ExpireAt    int64             `json:"expire_at"`   // 过期时间
	Category    string            `json:"category"`    // 种类
	Type        string            `json:"type"`        // 规格
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Status      string            `json:"status"`      // 服务商中的状态
	Tags        map[string]string `json:"tags"`        // 标签
	UpdateAt    int64             `json:"update_at"`   // 更新时间
	SyncAccount string            `json:"sync_accout"` // 同步的账号
	PublicIP    string            `json:"public_ip"`   // 公网IP
	PrivateIP   string            `json:"private_ip"`  // 内网IP
	PayType     string            `json:"pay_type"`    // 实例付费方式
}

type Describe struct {
	ResourceId              string `json:"resource_id"`                // 关联Resource
	CPU                     int    `json:"cpu"`                        // 核数
	Memory                  int    `json:"memory"`                     // 内存
	GPUAmount               int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec                 string `json:"gpu_spec"`                   // GPU类型
	OSType                  string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName                  string `json:"os_name"`                    // 操作系统名称
	SerialNumber            string `json:"serial_number"`              // 序列号
	ImageID                 string `json:"image_id"`                   // 镜像ID
	InternetMaxBandwidthOut int    `json:"internet_max_bandwidth_out"` // 公网出带宽最大值，单位为 Mbps
	InternetMaxBandwidthIn  int    `json:"internet_max_bandwidth_in"`  // 公网入带宽最大值，单位为 Mbps
	KeyPairName             string `json:"key_pair_name,omitempty"`    // 秘钥对名称
	SecurityGroups          string `json:"security_groups"`            // 安全组  采用逗号分隔
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}
