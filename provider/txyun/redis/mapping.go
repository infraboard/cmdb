package redis

import (
	"fmt"

	"github.com/infraboard/cmdb/apps/redis"
)

// + 0-待初始化
// + 1-流程中
// + 2-运行中
// + -2-已隔离
// + -3-待删除
var (
	STATUS_MAP = map[int64]redis.STATUS{
		0:  redis.STATUS_PENDING,
		1:  redis.STATUS_PENDING,
		2:  redis.STATUS_RUNNING,
		-2: redis.STATUS_ISOLATIONED,
		-3: redis.STATUS_ARCHIVED,
	}
)

func praseStatus(s *int64) string {
	if s == nil {
		return ""
	}

	if v, ok := STATUS_MAP[*s]; ok {
		return v.String()
	}

	return fmt.Sprintf("%d", *s)
}
