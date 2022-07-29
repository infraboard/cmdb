package elb

import "github.com/infraboard/cmdb/apps/lb"

var (
	// ONLINE
	// FROZEN
	ELB_STATUS_MAP = map[string]lb.STATUS{
		"ONLINE": lb.STATUS_RUNNING,
		"FROZEN": lb.STATUS_LOCKED,
	}
)

func praseElbStatus(s string) string {
	if v, ok := ELB_STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}
