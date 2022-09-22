package dds

import "github.com/infraboard/cmdb/apps/mongodb"

var (
	// 参考: https://help.aliyun.com/document_detail/63870.html?spm=api-workbench.API%20Explorer.0.0.1d601e0fQ8GR2n
	STATUS_MAP = map[string]mongodb.STATUS{
		"Creating":                  mongodb.STATUS_PENDING,
		"Running":                   mongodb.STATUS_RUNNING,
		"Deleting":                  mongodb.STATUS_DELETING,
		"Rebooting":                 mongodb.STATUS_REBOOTING,
		"Restoring":                 mongodb.STATUS_RESTORING,
		"Transing":                  mongodb.STATUS_TRANSING,
		"DBInstanceClassChanging":   mongodb.STATUS_MODIFYING,
		"GuardSwitching":            mongodb.STATUS_SWITCHOVER,
		"Importing":                 mongodb.STATUS_IMPORTING,
		"EngineVersionUpgrading":    mongodb.STATUS_UPGRADING,
		"DBInstanceNetTypeChanging": mongodb.STATUS_NET_CHANGING,
		"TransingToOthers":          mongodb.STATUS_TRANSING,
		"ImportingFromOthers":       mongodb.STATUS_IMPORTING,
	}
)

func praseStatus(s string) string {
	if v, ok := STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}
