package nlb

import (
	"strings"

	"github.com/infraboard/cmdb/apps/lb"
)

var (

	// 网络型负载均衡实例状态。取值：
	// Inactive：已停止，表示实例监听不会再转发流量。
	// Active：运行中。
	// Provisioning：创建中。
	// Configuring：变配中。
	// Deleting：删除中。
	// Deleted：已删除。
	SLB_STATUS_MAP = map[string]lb.STATUS{
		"inactive：":    lb.STATUS_STOPPED,
		"active":       lb.STATUS_RUNNING,
		"provisioning": lb.STATUS_PENDING,
		"configuring":  lb.STATUS_RUNNING,
		"deleting":     lb.STATUS_DELETING,
		"deleted":      lb.STATUS_DESTROYED,
	}
)

func praseSlbStatus(s *string) string {
	if s == nil {
		return ""
	}

	if v, ok := SLB_STATUS_MAP[strings.ToLower(*s)]; ok {
		return v.String()
	}

	return *s
}
