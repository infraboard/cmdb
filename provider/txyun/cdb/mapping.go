package cdb

import (
	"fmt"

	"github.com/infraboard/cmdb/apps/rds"
)

// 实例状态，可取值：
// + 0 - 创建中
// + 1 - 运行中
// + 4 - 正在进行隔离操作
// + 5 - 隔离中（可在回收站恢复开机）
var (
	STATUS_MAP = map[int64]rds.STATUS{
		0: rds.STATUS_PENDING,
		1: rds.STATUS_RUNNING,
		4: rds.STATUS_ISOLATIONING,
		5: rds.STATUS_ISOLATIONED,
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

// 实例任务状态，可能取值：
// + 0 - 没有任务

// + 1 - 升级中
// + 2 - 数据导入中
// + 6 - 回档中
// + 10 - 重启中
// + 12 - 自建迁移中
// + 14 - 灾备实例创建同步中
// + 15 - 升级待切换
// + 16 - 升级切换中
// + 17 - 升级切换完成
// + 4 - 外网访问开通中
// + 7 - 外网访问关闭中

// + 3 - 开放Slave中
// + 5 - 批量操作执行中
// + 8 - 密码修改中
// + 9 - 实例名修改中
// + 13 - 删除库表中
// + 19 - 参数设置待执行
var (
	TASK_STATUS_MAP = map[int64]rds.STATUS{
		1:  rds.STATUS_UPGRADING,
		2:  rds.STATUS_IMPORTING,
		6:  rds.STATUS_RESTORING,
		10: rds.STATUS_REBOOTING,
		12: rds.STATUS_TRANSING,
		14: rds.STATUS_GUARD_CREATING,
		15: rds.STATUS_UPGRADING,
		16: rds.STATUS_UPGRADING,
		17: rds.STATUS_UPGRADING,
		4:  rds.STATUS_NET_CHANGING,
		7:  rds.STATUS_NET_CHANGING,
		3:  rds.STATUS_MODIFYING,
		5:  rds.STATUS_MODIFYING,
		8:  rds.STATUS_MODIFYING,
		9:  rds.STATUS_MODIFYING,
		13: rds.STATUS_MODIFYING,
		19: rds.STATUS_MODIFYING,
	}
)

func praseTaskStatus(s *int64) string {
	if s == nil {
		return ""
	}

	if v, ok := STATUS_MAP[*s]; ok {
		return v.String()
	}

	return fmt.Sprintf("%d", *s)
}
