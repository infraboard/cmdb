package dcs

import "github.com/infraboard/cmdb/apps/redis"

// CREATING: 申请缓存实例后，在缓存实例状态进入运行中之前的状态。
// RUNNING: 缓存实例正常运行状态。在这个状态的实例可以运行您的业务。
// CREATEFAILED: 缓存实例处于创建失败的状态。
// ERROR: 缓存实例处于故障的状态。
// FLUSHING: 缓存实例数据清空中的状态
// RESTARTING: 缓存实例正在进行重启操作。
// FROZEN: 缓存实例处于已冻结状态，用户可以在“我的订单”中续费开启冻结的缓存实例。
// RESTORING: 缓存实例数据恢复中的状态。
// EXTENDING: 缓存实例处于正在扩容的状态。
var (
	STATUS_MAP = map[string]redis.STATUS{
		"RUNNING":      redis.STATUS_RUNNING,
		"CREATING":     redis.STATUS_PENDING,
		"ERROR":        redis.STATUS_ERROR,
		"CREATEFAILED": redis.STATUS_ERROR,
		"FLUSHING":     redis.STATUS_FLUSHING,
		"RESTARTING":   redis.STATUS_REBOOTING,
		"FROZEN":       redis.STATUS_ISOLATIONED,
		"RESTORING":    redis.STATUS_RESTORING,
		"EXTENDING":    redis.STATUS_EXTENDING,
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
