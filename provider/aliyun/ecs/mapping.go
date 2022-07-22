package ecs

import "github.com/infraboard/cmdb/apps/host"

var (
	STATUS_MAP = map[string]host.STATUS{
		"Pending":  host.STATUS_PENDING,
		"Running":  host.STATUS_RUNNING,
		"Starting": host.STATUS_STARTING,
		"Stopping": host.STATUS_STOPPING,
		"Stopped":  host.STATUS_STOPPED,
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
