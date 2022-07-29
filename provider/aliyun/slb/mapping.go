package slb

import "github.com/infraboard/cmdb/apps/lb"

var (
	// inactive: 实例已停止，此状态的实例监听不会再转发流量。
	// active: 实例运行中，实例创建后，默认状态为active。
	// locked: 实例已锁定。当负载均衡实例到期后，但到期时间未超过7天时，负载均衡实例进入锁定状态。
	//         此种状态下，您不能对负载均衡实例进行任何操作，并且实例不再会进行流量转发，但会保留实例的IP和其它配置。
	SLB_STATUS_MAP = map[string]lb.STATUS{
		"inactive": lb.STATUS_STOPPED,
		"active":   lb.STATUS_RUNNING,
		"locked":   lb.STATUS_LOCKED,
	}
)

func praseSlbStatus(s *string) string {
	if s == nil {
		return ""
	}

	if v, ok := SLB_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return *s
}
