package cvm

import (
	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/host"
)

var (
	CVM_STATUS_MAP = map[string]host.STATUS{
		"PENDING":               host.STATUS_PENDING,
		"LAUNCH_FAILED":         host.STATUS_CREATE_FAILED,
		"RUNNING":               host.STATUS_RUNNING,
		"STOPPED":               host.STATUS_STOPPED,
		"STARTING":              host.STATUS_STARTING,
		"SHUTDOWN":              host.STATUS_SHUTDOWN,
		"STOPPING":              host.STATUS_STOPPING,
		"REBOOTING":             host.STATUS_REBOOTING,
		"TERMINATING":           host.STATUS_TERMINATING,
		"EXPIRED":               host.STATUS_SHUTDOWN,
		"PROTECTIVELY_ISOLATED": host.STATUS_ERROR,
	}
)

func praseCvmStatus(s *string) string {
	if s == nil {
		return ""
	}
	if v, ok := CVM_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return *s
}

var (
	// UNATTACHED：未挂载
	// ATTACHING：挂载中
	// ATTACHED：已挂载
	// DETACHING：解挂中
	// EXPANDING：扩容中
	// ROLLBACKING：回滚中
	// TORECYCLE：待回收
	// DUMPING：拷贝硬盘中
	DISK_STATUS_MAP = map[string]disk.STATUS{
		"UNATTACHED":  disk.STATUS_UNATTACHED,
		"ATTACHING":   disk.STATUS_ATTACHING,
		"ATTACHED":    disk.STATUS_ATTACHED,
		"DETACHING":   disk.STATUS_DETACHING,
		"EXPANDING":   disk.STATUS_EXPANDING,
		"ROLLBACKING": disk.STATUS_ROLLBACKING,
		"TORECYCLE":   disk.STATUS_RECYCLE,
		"DUMPING":     disk.STATUS_DUMPING,
	}
)

func praseDiskStatus(s *string) string {
	if s == nil {
		return ""
	}
	if v, ok := DISK_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return *s
}
