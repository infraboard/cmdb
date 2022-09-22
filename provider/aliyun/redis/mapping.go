package redis

import "github.com/infraboard/cmdb/apps/redis"

var (
	STATUS_MAP = map[string]redis.STATUS{
		"Normal":                redis.STATUS_RUNNING,
		"Creating":              redis.STATUS_PENDING,
		"Error":                 redis.STATUS_ERROR,
		"Flushing":              redis.STATUS_FLUSHING,
		"Released":              redis.STATUS_DESTROYED,
		"BackupRecovering":      redis.STATUS_RESTORING,
		"MinorVersionUpgrading": redis.STATUS_UPGRADING,
		"MajorVersionUpgrading": redis.STATUS_UPGRADING,
		"NetworkModifying":      redis.STATUS_NET_CHANGING,
		"Inactive":              redis.STATUS_ISOLATIONED,
		"Changing":              redis.STATUS_MODIFYING,
		"Transforming":          redis.STATUS_TRANSING,
		"Migrating":             redis.STATUS_TRANSING,
		"Unavailable":           redis.STATUS_ERROR,
		"SSLModifying":          redis.STATUS_MODIFYING,
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
