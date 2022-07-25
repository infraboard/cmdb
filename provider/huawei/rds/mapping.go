package rds

import (
	"strings"

	"github.com/infraboard/cmdb/apps/rds"
)

var (
	STATUS_MAP = map[string]rds.STATUS{
		"BUILD":                   rds.STATUS_PENDING,
		"ACTIVE":                  rds.STATUS_RUNNING,
		"FAILED":                  rds.STATUS_ERROR,
		"FROZEN":                  rds.STATUS_ISOLATIONED,
		"REBOOTING":               rds.STATUS_REBOOTING,
		"RESTORING":               rds.STATUS_RESTORING,
		"MIGRATING":               rds.STATUS_TRANSING,
		"MODIFYING":               rds.STATUS_MODIFYING,
		"SWITCHOVER":              rds.STATUS_SWITCHOVER,
		"STORAGE FULL":            rds.STATUS_STORAGE_FULL,
		"BACKING UP":              rds.STATUS_BACKING_UP,
		"INS_CLONING":             rds.STATUS_CLONING,
		"MODIFYING INSTANCE TYPE": rds.STATUS_MODIFYING,
		"MODIFYING DATABASE PORT": rds.STATUS_MODIFYING,
	}
)

func praseStatus(s string) string {
	s = strings.ToUpper(s)

	if v, ok := STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}
