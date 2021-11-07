package rds

import (
	"github.com/infraboard/cmdb/app/resource"
)

func NewDefaultRds() *Rds {
	return &Rds{
		&resource.Base{
			ResourceType: resource.Type_RDS,
		},
		&resource.Information{},
		&Describe{},
	}
}

type Rds struct {
	*resource.Base
	*resource.Information
	*Describe
}

type Describe struct {
	Category           string   `json:"category"`             // 实例系列，取值：Basic：基础版；HighAvailability：高可用版；AlwaysOn：集群版；Finance：金融版（仅中国站支持）
	Engine             string   `json:"engine"`               // 引擎 比如 MYSQL, SQLServer, PGSQL
	EngineVersion      string   `json:"engine_version"`       // 实例版本
	Class              string   `json:"class"`                // 实例规格: 对应ALI(DBInstanceClass)
	ClassType          string   `json:"class_type"`           // 实例规格族，取值：s：共享型；x：通用型；d：独享套餐；h：独占物理机。
	NetType            string   `json:"net_type"`             // 实例是内网或外网 (Internet：外网/Intranet：内网)
	NetworkType        string   `json:"network_type"`         // 实例的网络类型 (Classic：经典网络/VPC：专有网络。)
	Type               string   `json:"type"`                 // 实例类型 Primary：主实例, Readonly：只读实例, Guard：灾备实例, Temp：临时实例
	CPU                int      `json:"cpu"`                  // CPU 核数
	Memory             int64    `json:"memory"`               // 实例内存，单位：M。
	DBMaxQuantity      int      `json:"db_max_quantity"`      // 一个实例下可创建最大数据库数量
	AccountMaxQuantity int      `json:"account_max_quantity"` // 可创建账号的最大数量。
	MaxConnections     int      `json:"max_connections"`      // 最大并发连接数
	MaxIOPS            int      `json:"max_iops"`             // 最大每秒IO请求次数
	Collation          string   `json:"collation"`            // 系统字符集排序规则
	TimeZone           string   `json:"timeZone"`             // 时区
	Storage            int64    `json:"storage"`              // 实例存储空间，单位：GB。
	StorageType        string   `json:"storage_type"`         // 实例储存类型 local_ssd/ephemeral_ssd：本地SSD盘, cloud_ssd：SSD云盘；cloud_essd：ESSD云盘
	SecurityIPMode     string   `json:"security_ip_mode"`     // 安全名单模式, 默认白名单
	SecurityIPList     []string `json:"security_ip_list"`     // IP白名单
	PayType            string   `json:"pay_type"`             // 付费方式
	ConnectionMode     string   `json:"connection_mode"`      // 实例的访问模式，取值：Standard：标准访问模式；Safe：数据库代理模式。
	IPType             string   `json:"ip_type"`              // IP类型
	LockMode           string   `json:"lock_mode"`            // 实例锁定模式; Unlock：正常；ManualLock：手动触发锁定；LockByExpiration：实例过期自动锁定；LockByRestoration：实例回滚前的自动锁定；LockByDiskQuota：实例空间满自动锁定
	LockReason         string   `json:"lock_reason"`          // 锁定原因
	Status             string   `json:"status"`               // 当前状态
	ExpireTime         int64    `json:"expire_time"`          // 过期时间
	Description        string   `json:"description"`          // 描述
	Port               int64    `json:"port"`                 // 端口
}

func NewRdsSet() *RdsSet {
	return &RdsSet{
		Items: []*Rds{},
	}
}

type RdsSet struct {
	Items []*Rds `json:"items"`
	Total int64  `json:"total"`
}

func (s *RdsSet) Add(item *Rds) {
	s.Items = append(s.Items, item)
}
