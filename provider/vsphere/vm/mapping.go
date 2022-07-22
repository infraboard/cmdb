package vm

import "github.com/infraboard/cmdb/apps/host"

var (
	STATUS_MAP = map[string]host.STATUS{
		"poweredOff": host.STATUS_STOPPED,
		"poweredOn":  host.STATUS_RUNNING,
		"suspended":  host.STATUS_SUSPENDED,
	}
)

func praseStatus(s string) string {
	if v, ok := STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}
