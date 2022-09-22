package rds

import "github.com/infraboard/cmdb/apps/rds"

var (
	STATUS_MAP = map[string]rds.STATUS{
		"Creating":                  rds.STATUS_PENDING,
		"Running":                   rds.STATUS_RUNNING,
		"Deleting":                  rds.STATUS_DELETING,
		"Released":                  rds.STATUS_DESTROYED,
		"Rebooting":                 rds.STATUS_REBOOTING,
		"Restoring":                 rds.STATUS_RESTORING,
		"TRANSING":                  rds.STATUS_TRANSING,
		"DBInstanceClassChanging":   rds.STATUS_MODIFYING,
		"GuardSwitching":            rds.STATUS_SWITCHOVER,
		"GuardDBInstanceCreating":   rds.STATUS_GUARD_CREATING,
		"Importing":                 rds.STATUS_IMPORTING,
		"INS_CLONING":               rds.STATUS_CLONING,
		"EngineVersionUpgrading":    rds.STATUS_UPGRADING,
		"DBInstanceNetTypeChanging": rds.STATUS_NET_CHANGING,
		"TransingToOthers":          rds.STATUS_TRANSING,
		"ImportingFromOthers":       rds.STATUS_IMPORTING,
	}
)

func praseStatus(s string) string {
	if v, ok := STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}
