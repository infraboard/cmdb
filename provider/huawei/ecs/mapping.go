package ecs

import (
	"strings"

	"github.com/infraboard/cmdb/apps/disk"
	"github.com/infraboard/cmdb/apps/eip"
	"github.com/infraboard/cmdb/apps/host"
)

var (
	ECS_STATUS_MAP = map[string]host.STATUS{
		"ERROR":         host.STATUS_ERROR,
		"RESIZE":        host.STATUS_PENDING,
		"REVERT_RESIZE": host.STATUS_PENDING,
		"VERIFY_RESIZE": host.STATUS_PENDING,
		"MIGRATING":     host.STATUS_PENDING,
		"ACTIVE":        host.STATUS_RUNNING,
		"SHUTOFF":       host.STATUS_STOPPED,
		"BUILD":         host.STATUS_PENDING,
		"REBUILD":       host.STATUS_PENDING,
		"REBOOT":        host.STATUS_REBOOTING,
		"HARD_REBOOT":   host.STATUS_REBOOTING,
		"DELETING":      host.STATUS_DELETING,
		"DELETED":       host.STATUS_DESTROYED,
	}
)

func praseEcsStatus(s string) string {
	if v, ok := ECS_STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}

var (
	// creating 云硬盘处于正在创建的过程中。
	// available 云硬盘创建成功，还未挂载给任何云服务器，可以进行挂载。
	// in-use 云硬盘已挂载给云服务器，正在使用中。
	// error 云硬盘在创建过程中出现错误。
	// attaching 云硬盘处于正在挂载的过程中。
	// detaching 云硬盘处于正在卸载的过程中。
	// restoring-backup 云硬盘处于正在从备份恢复的过程中。
	// backing-up 云硬盘处于通过备份创建的过程中。
	// error_restoring 云硬盘从备份恢复过程中出现错误。
	// uploading 云硬盘数据正在被上传到镜像中。此状态出现在从云服务器创建镜像的操作过程中。
	// downloading 正在从镜像下载数据到云硬盘。此状态出现在创建云服务器的操作过程中。
	// extending 云硬盘处于正在扩容的过程中。
	// error_extending 云硬盘在扩容过程中出现错误。
	// deleting 云硬盘处于正在删除的过程中。
	// error_deleting云硬盘在删除过程中出现错误。
	// rollbacking 云硬盘处于正在从快照回滚数据的过程中。
	// error_rollbacking 云硬盘在从快照回滚数据的过程中出现错误。
	// awaiting-transfer 云硬盘处于等待过户的过程中。
	DISK_STATUS_MAP = map[string]disk.STATUS{
		"creating":         disk.STATUS_PENDING,
		"available":        disk.STATUS_UNATTACHED,
		"in-use":           disk.STATUS_ATTACHED,
		"attaching":        disk.STATUS_ATTACHING,
		"detaching":        disk.STATUS_DETACHING,
		"restoring-backup": disk.STATUS_ROLLBACKING,
		"backing-up":       disk.STATUS_ROLLBACKING,
		"rollbacking":      disk.STATUS_ROLLBACKING,
		"extending":        disk.STATUS_EXPANDING,
		"deleting":         disk.STATUS_RECYCLE,
	}
)

func praseDiskStatus(s string) string {
	s = strings.ToLower(s)

	if v, ok := DISK_STATUS_MAP[s]; ok {
		return v.String()
	}

	return s
}

var (
	// FREEZED 冻结
	// BIND_ERROR 绑定失败
	// BINDING  绑定中
	// PENDING_DELETE 释放中
	// PENDING_CREATE 创建中
	// NOTIFYING 创建中
	// NOTIFY_DELETE 释放中
	// PENDING_UPDATE 更新中
	// DOWN 未绑定
	// ACTIVE 绑定
	// ELB 绑定ELB
	// ERROR 失败
	// VPN 绑定VPN
	EIP_STATUS_MAP = map[string]eip.STATUS{
		"BINDING":        eip.STATUS_BINDING,
		"PENDING_CREATE": eip.STATUS_PENDING,
		"PENDING_DELETE": eip.STATUS_OFFLINING,
		"NOTIFYING":      eip.STATUS_PENDING,
		"NOTIFY_DELETE":  eip.STATUS_OFFLINING,
		"PENDING_UPDATE": eip.STATUS_PENDING,
		"DOWN":           eip.STATUS_UNBIND,
		"ACTIVE":         eip.STATUS_BIND,
		"ELB":            eip.STATUS_BIND,
		"VPN":            eip.STATUS_BIND,
		"BIND_ERROR":     eip.STATUS_ERROR,
		"ERROR":          eip.STATUS_ERROR,
	}
)

func praseEIPStatus(s string) string {
	s = strings.ToLower(s)

	if v, ok := EIP_STATUS_MAP[strings.ToUpper(s)]; ok {
		return v.String()
	}

	return s
}
