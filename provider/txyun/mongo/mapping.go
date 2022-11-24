package mongo

import (
	"fmt"

	"github.com/infraboard/cmdb/apps/mongodb"
)

var (
	// 0-待初始化，1-流程执行中，2-实例有效，-2-已隔离（包年包月实例），-3-已隔离（按量计费实例）
	// 参考文档: https://console.cloud.tencent.com/api/explorer?Product=mongodb&Version=2019-07-25&Action=DescribeDBInstances&SignVersion=
	STATUS_MAP = map[int64]mongodb.STATUS{
		0:  mongodb.STATUS_PENDING,
		1:  mongodb.STATUS_PENDING,
		2:  mongodb.STATUS_RUNNING,
		-2: mongodb.STATUS_ISOLATIONED,
		-3: mongodb.STATUS_ISOLATIONED,
	}
)

func praseStatus(s int64) string {
	if v, ok := STATUS_MAP[s]; ok {
		return v.String()
	}

	return fmt.Sprintf("%d", s)
}
