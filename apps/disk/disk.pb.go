// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/disk/pb/disk.proto

package disk

import (
	resource "github.com/infraboard/cmdb/apps/resource"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type STATUS int32

const (
	// 未知状态
	STATUS_UNKNOW STATUS = 0
	// 表示创建中
	STATUS_PENDING STATUS = 5
	// 挂载中
	STATUS_ATTACHING STATUS = 6
	// 未挂载
	STATUS_UNATTACHED STATUS = 7
	// 已挂载
	STATUS_ATTACHED STATUS = 11
	// 解挂中
	STATUS_DETACHING STATUS = 30
	// 回滚中
	STATUS_ROLLBACKING STATUS = 31
	// 扩容中
	STATUS_EXPANDING STATUS = 32
	// 实例克隆中
	STATUS_DUMPING STATUS = 34
	// 表示停止待销毁
	STATUS_RECYCLE STATUS = 80
	// 已销毁
	STATUS_DESTROYED STATUS = 90
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0:  "UNKNOW",
		5:  "PENDING",
		6:  "ATTACHING",
		7:  "UNATTACHED",
		11: "ATTACHED",
		30: "DETACHING",
		31: "ROLLBACKING",
		32: "EXPANDING",
		34: "DUMPING",
		80: "RECYCLE",
		90: "DESTROYED",
	}
	STATUS_value = map[string]int32{
		"UNKNOW":      0,
		"PENDING":     5,
		"ATTACHING":   6,
		"UNATTACHED":  7,
		"ATTACHED":    11,
		"DETACHING":   30,
		"ROLLBACKING": 31,
		"EXPANDING":   32,
		"DUMPING":     34,
		"RECYCLE":     80,
		"DESTROYED":   90,
	}
)

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}

func (x STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_disk_pb_disk_proto_enumTypes[0].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_apps_disk_pb_disk_proto_enumTypes[0]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_apps_disk_pb_disk_proto_rawDescGZIP(), []int{0}
}

type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"resource"
	Resource *resource.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource"`
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,2,opt,name=describe,proto3" json:"describe"`
}

func (x *Disk) Reset() {
	*x = Disk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_disk_pb_disk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disk) ProtoMessage() {}

func (x *Disk) ProtoReflect() protoreflect.Message {
	mi := &file_apps_disk_pb_disk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disk.ProtoReflect.Descriptor instead.
func (*Disk) Descriptor() ([]byte, []int) {
	return file_apps_disk_pb_disk_proto_rawDescGZIP(), []int{0}
}

func (x *Disk) GetResource() *resource.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Disk) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 类型 (系统盘, 数据盘) 阿里: system：系统盘; data：数据盘
	// @gotags: json:"type"
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	// 关联的实例
	// @gotags: json:"instance_id"
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id"`
	// 容量大小
	// @gotags: json:"size"
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size"`
	// IOPS
	// @gotags: json:"iops"
	Iops int32 `protobuf:"varint,11,opt,name=iops,proto3" json:"iops"`
	// 卸载时间
	// @gotags: json:"detached_time"
	DetachedTime int64 `protobuf:"varint,4,opt,name=detached_time,json=detachedTime,proto3" json:"detached_time"`
	// 挂载时间
	// @gotags: json:"attached_time"
	AttachedTime int64 `protobuf:"varint,5,opt,name=attached_time,json=attachedTime,proto3" json:"attached_time"`
	// 是否为弹性云盘
	// @gotags: json:"portable"
	Portable bool `protobuf:"varint,6,opt,name=portable,proto3" json:"portable"`
	// 比如: /dev/xvdc
	// @gotags: json:"device"
	Device string `protobuf:"bytes,7,opt,name=device,proto3" json:"device"`
	// 是否加密
	// @gotags: json:"encrypted"
	Encrypted bool `protobuf:"varint,8,opt,name=encrypted,proto3" json:"encrypted"`
	// 自动快照
	// @gotags: json:"enable_auto_snapshot"
	EnableAutoSnapshot bool `protobuf:"varint,9,opt,name=enable_auto_snapshot,json=enableAutoSnapshot,proto3" json:"enable_auto_snapshot"`
	// 随实例销毁
	// @gotags: json:"delete_with_instance"
	DeleteWithInstance bool `protobuf:"varint,10,opt,name=delete_with_instance,json=deleteWithInstance,proto3" json:"delete_with_instance"`
	// 是否是共享盘
	// @gotags: json:"multi_attach"
	MultiAttach bool `protobuf:"varint,12,opt,name=multi_attach,json=multiAttach,proto3" json:"multi_attach"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_disk_pb_disk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_disk_pb_disk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Describe.ProtoReflect.Descriptor instead.
func (*Describe) Descriptor() ([]byte, []int) {
	return file_apps_disk_pb_disk_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Describe) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Describe) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Describe) GetIops() int32 {
	if x != nil {
		return x.Iops
	}
	return 0
}

func (x *Describe) GetDetachedTime() int64 {
	if x != nil {
		return x.DetachedTime
	}
	return 0
}

func (x *Describe) GetAttachedTime() int64 {
	if x != nil {
		return x.AttachedTime
	}
	return 0
}

func (x *Describe) GetPortable() bool {
	if x != nil {
		return x.Portable
	}
	return false
}

func (x *Describe) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Describe) GetEncrypted() bool {
	if x != nil {
		return x.Encrypted
	}
	return false
}

func (x *Describe) GetEnableAutoSnapshot() bool {
	if x != nil {
		return x.EnableAutoSnapshot
	}
	return false
}

func (x *Describe) GetDeleteWithInstance() bool {
	if x != nil {
		return x.DeleteWithInstance
	}
	return false
}

func (x *Describe) GetMultiAttach() bool {
	if x != nil {
		return x.MultiAttach
	}
	return false
}

type DiskSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表项
	// @gotags: json:"items"
	Items []*Disk `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *DiskSet) Reset() {
	*x = DiskSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_disk_pb_disk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskSet) ProtoMessage() {}

func (x *DiskSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_disk_pb_disk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskSet.ProtoReflect.Descriptor instead.
func (*DiskSet) Descriptor() ([]byte, []int) {
	return file_apps_disk_pb_disk_proto_rawDescGZIP(), []int{2}
}

func (x *DiskSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DiskSet) GetItems() []*Disk {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_disk_pb_disk_proto protoreflect.FileDescriptor

var file_apps_disk_pb_disk_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x64,
	0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x1a,
	0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69,
	0x73, 0x6b, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x8a, 0x03, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x6f, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x69, 0x6f, 0x70, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x30, 0x0a,
	0x14, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x22, 0x51, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0xa6, 0x01, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x54, 0x54,
	0x41, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x41, 0x54,
	0x54, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x54, 0x54, 0x41,
	0x43, 0x48, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x54, 0x41, 0x43, 0x48,
	0x49, 0x4e, 0x47, 0x10, 0x1e, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x4c, 0x42, 0x41, 0x43,
	0x4b, 0x49, 0x4e, 0x47, 0x10, 0x1f, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x50, 0x41, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x20, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x55, 0x4d, 0x50, 0x49, 0x4e, 0x47,
	0x10, 0x22, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x10, 0x50, 0x12,
	0x0d, 0x0a, 0x09, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x5a, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_disk_pb_disk_proto_rawDescOnce sync.Once
	file_apps_disk_pb_disk_proto_rawDescData = file_apps_disk_pb_disk_proto_rawDesc
)

func file_apps_disk_pb_disk_proto_rawDescGZIP() []byte {
	file_apps_disk_pb_disk_proto_rawDescOnce.Do(func() {
		file_apps_disk_pb_disk_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_disk_pb_disk_proto_rawDescData)
	})
	return file_apps_disk_pb_disk_proto_rawDescData
}

var file_apps_disk_pb_disk_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_disk_pb_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_disk_pb_disk_proto_goTypes = []interface{}{
	(STATUS)(0),               // 0: infraboard.cmdb.disk.STATUS
	(*Disk)(nil),              // 1: infraboard.cmdb.disk.Disk
	(*Describe)(nil),          // 2: infraboard.cmdb.disk.Describe
	(*DiskSet)(nil),           // 3: infraboard.cmdb.disk.DiskSet
	(*resource.Resource)(nil), // 4: infraboard.cmdb.resource.Resource
}
var file_apps_disk_pb_disk_proto_depIdxs = []int32{
	4, // 0: infraboard.cmdb.disk.Disk.resource:type_name -> infraboard.cmdb.resource.Resource
	2, // 1: infraboard.cmdb.disk.Disk.describe:type_name -> infraboard.cmdb.disk.Describe
	1, // 2: infraboard.cmdb.disk.DiskSet.items:type_name -> infraboard.cmdb.disk.Disk
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_disk_pb_disk_proto_init() }
func file_apps_disk_pb_disk_proto_init() {
	if File_apps_disk_pb_disk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_disk_pb_disk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_disk_pb_disk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Describe); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_disk_pb_disk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_disk_pb_disk_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_disk_pb_disk_proto_goTypes,
		DependencyIndexes: file_apps_disk_pb_disk_proto_depIdxs,
		EnumInfos:         file_apps_disk_pb_disk_proto_enumTypes,
		MessageInfos:      file_apps_disk_pb_disk_proto_msgTypes,
	}.Build()
	File_apps_disk_pb_disk_proto = out.File
	file_apps_disk_pb_disk_proto_rawDesc = nil
	file_apps_disk_pb_disk_proto_goTypes = nil
	file_apps_disk_pb_disk_proto_depIdxs = nil
}
