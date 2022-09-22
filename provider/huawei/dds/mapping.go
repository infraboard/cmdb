package dds

import "github.com/infraboard/cmdb/apps/mongodb"

var (
	// 参考: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=DDS&api=ListInstances
	// normal,表示实例正常。
	// abnormal,表示实例异常。
	// creating,表示实例创建中。
	// frozen,表示实例被冻结。
	// data_disk_full,表示实例磁盘已满。
	// createfail,表示实例创建失败。
	// enlargefail,表示实例扩容节点个数失败。
	STATUS_MAP = map[string]mongodb.STATUS{
		"creating":       mongodb.STATUS_PENDING,
		"normal":         mongodb.STATUS_RUNNING,
		"abnormal":       mongodb.STATUS_ERROR,
		"frozen":         mongodb.STATUS_LOCKED,
		"data_disk_full": mongodb.STATUS_STORAGE_FULL,
		"createfail":     mongodb.STATUS_CREATE_FAILED,
		"enlargefail":    mongodb.STATUS_ERROR,
	}
)

func praseStatus(s string) string {
	if v, ok := STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}
