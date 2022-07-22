package ecs

import "github.com/infraboard/cmdb/apps/host"

var (
	STATUS_MAP = map[string]host.STATUS{
		"ERROR":         host.STATUS_ERROR,
		"RESIZE":        host.STATUS_PENDING,
		"REVERT_RESIZE": host.STATUS_PENDING,
		"VERIFY_RESIZE": host.STATUS_PENDING,
		"MIGRATING":     host.STATUS_PENDING,
		"ACTIVE":        host.STATUS_RUNNING,
		"SHUTOFF":       host.STATUS_STOPPED,
		"BUILD":         host.STATUS_PENDING,
		"REBUILD":       host.STATUS_PENDING,
		"REBOOT":        host.STATUS_REBOOTING,
		"HARD_REBOOT":   host.STATUS_REBOOTING,
		"TERMINATING":   host.STATUS_TERMINATING,
		"DELETED":       host.STATUS_ARCHIVED,
	}
)

func praseStatus(s string) string {
	if v, ok := STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}
