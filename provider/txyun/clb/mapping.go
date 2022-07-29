package clb

import (
	"fmt"

	"github.com/infraboard/cmdb/apps/lb"
)

var (
	// 0：创建中
	// 1：正常运行。
	CLB_STATUS_MAP = map[uint64]lb.STATUS{
		0: lb.STATUS_PENDING,
		1: lb.STATUS_RUNNING,
	}
)

func praseClbStatus(s *uint64) string {
	if s == nil {
		return ""
	}

	if v, ok := CLB_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return fmt.Sprintf("%d", *s)
}
