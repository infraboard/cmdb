package sqlserver

import (
	"fmt"

	"github.com/infraboard/cmdb/apps/rds"
)

// 实例状态，可取值：
// 1：申请中
// 2：运行中
// 3：受限运行中 (主备切换中)
// 4：已隔离
// 5：回收中
// 6：已回收
// 7：任务执行中 (实例做备份、回档等操作)
// 8：已下线
// 9：实例扩容中
// 10：实例迁移中
// 11：只读
// 12：重启中
var (
	STATUS_MAP = map[int64]rds.STATUS{
		1:  rds.STATUS_PENDING,
		2:  rds.STATUS_RUNNING,
		3:  rds.STATUS_SWITCHOVER,
		4:  rds.STATUS_ISOLATIONED,
		5:  rds.STATUS_DELETING,
		6:  rds.STATUS_DESTROYED,
		7:  rds.STATUS_RUNNING,
		8:  rds.STATUS_SHUTDOWN,
		9:  rds.STATUS_RUNNING,
		10: rds.STATUS_TRANSING,
		11: rds.STATUS_RUNNING,
		12: rds.STATUS_REBOOTING,
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
