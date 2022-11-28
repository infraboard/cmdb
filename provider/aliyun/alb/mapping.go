package alb

import (
	"strings"

	"github.com/infraboard/cmdb/apps/lb"
)

var (

	// Inactive： 已停止，监听不再转发流量。
	// Active:：运行中。
	// Provisioning：创建中。
	// Configuring：变配中。
	// CreateFailed：创建失败，此时不会产生费用，实例只能被删除。系统默认清理最近1天创建失败的实例。
	SLB_STATUS_MAP = map[string]lb.STATUS{
		"inactive：":    lb.STATUS_STOPPED,
		"active":       lb.STATUS_RUNNING,
		"provisioning": lb.STATUS_PENDING,
		"configuring":  lb.STATUS_RUNNING,
		"createFailed": lb.STATUS_CREATE_FAILED,
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
