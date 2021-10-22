package host

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/infraboard/cmdb/pkg/resource"
	"github.com/infraboard/mcube/types/ftime"
)

func NewDefaultHost() *Host {
	return &Host{
		&resource.Base{
			ResourceType: resource.HostResource,
		},
		&resource.Information{},
		&Describe{},
	}
}

type Host struct {
	*resource.Base
	*resource.Information
	*Describe
}

func (h *Host) Put(req *UpdateHostData) {
	h.Information = req.Information
	h.Describe = req.Describe
	h.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
}

func (h *Host) Patch(req *UpdateHostData) error {
	err := ObjectPatch(h.Information, req.Information)
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

func ObjectPatch(old, new interface{}) error {
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}
	return json.Unmarshal(newByte, old)
}

func (h *Host) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h.Information)
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

type Describe struct {
	ResourceId              string   `json:"resource_id"`                // 关联Resource
	CPU                     int64    `json:"cpu"`                        // 核数
	Memory                  int64    `json:"memory"`                     // 内存
	GPUAmount               int      `json:"gpu_amount"`                 // GPU数量
	GPUSpec                 string   `json:"gpu_spec"`                   // GPU类型
	OSType                  string   `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName                  string   `json:"os_name"`                    // 操作系统名称
	SerialNumber            string   `json:"serial_number"`              // 序列号
	ImageID                 string   `json:"image_id"`                   // 镜像ID
	InternetMaxBandwidthOut int64    `json:"internet_max_bandwidth_out"` // 公网出带宽最大值，单位为 Mbps
	InternetMaxBandwidthIn  int64    `json:"internet_max_bandwidth_in"`  // 公网入带宽最大值，单位为 Mbps
	KeyPairName             []string `json:"key_pair_name,omitempty"`    // 秘钥对名称
	SecurityGroups          []string `json:"security_groups"`            // 安全组  采用逗号分隔
}

func (d *Describe) KeyPairNameToString() string {
	return strings.Join(d.KeyPairName, ",")
}

func (d *Describe) SecurityGroupsToString() string {
	return strings.Join(d.SecurityGroups, ",")
}

func (d *Describe) LoadKeyPairNameString(s string) {
	if s != "" {
		d.KeyPairName = strings.Split(s, ",")
	}
}

func (d *Describe) LoadSecurityGroupsString(s string) {
	if s != "" {
		d.SecurityGroups = strings.Split(s, ",")
	}
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

type HostSet struct {
	Items []*Host `json:"items"`
	Total int64   `json:"total"`
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}

func (s *HostSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
