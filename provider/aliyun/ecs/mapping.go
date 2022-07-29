package ecs

import (
	"strings"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/host"
)

var (
	ECS_STATUS_MAP = map[string]host.STATUS{
		"Pending":  host.STATUS_PENDING,
		"Running":  host.STATUS_RUNNING,
		"Starting": host.STATUS_STARTING,
		"Stopping": host.STATUS_STOPPING,
		"Stopped":  host.STATUS_STOPPED,
	}
)

func praseEcsStatus(s *string) string {
	if s == nil {
		return ""
	}
	if v, ok := ECS_STATUS_MAP[*s]; ok {
		return v.String()
	}

	return *s
}

// Creating	创建中。 通过RunInstances、CreateInstance或CreateDisk创建了云盘后，云盘进入短暂的创建中状态。
// Available	待挂载。 通过CreateDisk成功创建一块按量付费云盘或通过DetachDisk卸载一块按量付费数据盘后，云盘进入稳定的待挂载状态。
// In_Use	使用中。 云盘的稳定状态，
// ReIniting	初始化中。通过ReInitDisk重新初始化一块系统盘或者数据盘后，云盘进入短暂的初始化中状态。
// Detaching	卸载中。 通过DetachDisk卸载一块按量付费数据盘后，云盘进入短暂的卸载中状态。
// Deleting*	删除中。 通过DeleteDisk释放一块按量付费数据盘后，云盘进入短暂的删除中状态。
// Deleted*	已删除。 通过DeleteDisk释放一块按量付费数据盘后，云盘进入短暂的已删除状态。
var (
	DISK_STATUS_MAP = map[string]disk.STATUS{
		"creating":  disk.STATUS_PENDING,
		"available": disk.STATUS_UNATTACHED,
		"in_use":    disk.STATUS_ATTACHED,
		"reiniting": disk.STATUS_PENDING,
		"detaching": disk.STATUS_DETACHING,
		"deleting*": disk.STATUS_RECYCLE,
		"deleted*":  disk.STATUS_ARCHIVED,
	}
)

func praseDiskStatus(s *string) string {
	if s == nil {
		return ""
	}

	t := strings.ToLower(*s)
	if v, ok := DISK_STATUS_MAP[t]; ok {
		return v.String()
	}

	return *s
}
