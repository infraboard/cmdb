package cvm

import "github.com/infraboard/cmdb/apps/host"

var (
	STATUS_MAP = map[string]host.STATUS{
		"PENDING":               host.STATUS_PENDING,
		"LAUNCH_FAILED":         host.STATUS_LAUNCH_FAILED,
		"RUNNING":               host.STATUS_RUNNING,
		"STOPPED":               host.STATUS_STOPPED,
		"STARTING":              host.STATUS_STARTING,
		"SHUTDOWN":              host.STATUS_SHUTDOWN,
		"STOPPING":              host.STATUS_STOPPING,
		"REBOOTING":             host.STATUS_REBOOTING,
		"TERMINATING":           host.STATUS_TERMINATING,
		"EXPIRED":               host.STATUS_SHUTDOWN,
		"PROTECTIVELY_ISOLATED": host.STATUS_ERROR,
	}
)

func praseStatus(s *string) string {
	if s == nil {
		return ""
	}
	if v, ok := STATUS_MAP[*s]; ok {
		return v.String()
	}

	return *s
}
