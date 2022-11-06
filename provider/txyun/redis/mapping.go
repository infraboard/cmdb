package redis

import (
	"fmt"
	"time"

	"github.com/infraboard/cmdb/apps/redis"
	"github.com/infraboard/cmdb/apps/resource"
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
		-3: redis.STATUS_DESTROYED,
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

var (
	typeMap = map[int64]string{
		1: "Redis2.8内存版（集群架构）",
		2: "Redis2.8内存版（标准架构）",
		3: "CKV 3.2内存版(标准架构)",
		4: "CKV 3.2内存版(集群架构)",
		5: "Redis2.8内存版（单机）",
		6: "Redis4.0内存版（标准架构）",
		7: "Redis4.0内存版（集群架构）",
		8: "Redis5.0内存版（标准架构）",
		9: "Redis5.0内存版（集群架构）",
	}
)

func (o *RedisOperator) ParseType(t *int64) string {
	if t == nil {
		return ""
	}
	return typeMap[*t]
}

func (o *RedisOperator) parsePAY_MODE(t *int64) resource.PAY_MODE {
	if t == nil {
		return resource.PAY_MODE_NULL
	}

	switch *t {
	case 0:
		return resource.PAY_MODE_POST_PAY
	case 1:
		return resource.PAY_MODE_PRE_PAY
	default:

	}

	return resource.PAY_MODE_NULL
}

func (o *RedisOperator) parseTime(t string) int64 {
	ts, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixMilli()
}
