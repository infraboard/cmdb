// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/disk/pb/disk.proto

package disk

import (
	resource "github.com/infraboard/cmdb/apps/resource"
	request "github.com/infraboard/mcube/http/request"
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

type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        *resource.Base        `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Information *resource.Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information,omitempty"`
	Describe    *Describe             `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe,omitempty"`
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

func (x *Disk) GetBase() *resource.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Disk) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
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

type Set struct {
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

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_disk_pb_disk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_apps_disk_pb_disk_proto_rawDescGZIP(), []int{2}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*Disk {
	if x != nil {
		return x.Items
	}
	return nil
}

type QueryDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryDiskRequest) Reset() {
	*x = QueryDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_disk_pb_disk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDiskRequest) ProtoMessage() {}

func (x *QueryDiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_disk_pb_disk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDiskRequest.ProtoReflect.Descriptor instead.
func (*QueryDiskRequest) Descriptor() ([]byte, []int) {
	return file_apps_disk_pb_disk_proto_rawDescGZIP(), []int{3}
}

func (x *QueryDiskRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
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
	0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbf, 0x01, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x22, 0xd3, 0x02, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x64, 0x65, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x44, 0x69, 0x73, 0x6b,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4a, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x32, 0xa1, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1a, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69,
	0x73, 0x6b, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x44,
	0x69, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69,
	0x73, 0x6b, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44,
	0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x64, 0x69, 0x73,
	0x6b, 0x2e, 0x53, 0x65, 0x74, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_apps_disk_pb_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_disk_pb_disk_proto_goTypes = []interface{}{
	(*Disk)(nil),                 // 0: infraboard.cmdb.disk.Disk
	(*Describe)(nil),             // 1: infraboard.cmdb.disk.Describe
	(*Set)(nil),                  // 2: infraboard.cmdb.disk.Set
	(*QueryDiskRequest)(nil),     // 3: infraboard.cmdb.disk.QueryDiskRequest
	(*resource.Base)(nil),        // 4: infraboard.cmdb.resource.Base
	(*resource.Information)(nil), // 5: infraboard.cmdb.resource.Information
	(*request.PageRequest)(nil),  // 6: infraboard.mcube.page.PageRequest
}
var file_apps_disk_pb_disk_proto_depIdxs = []int32{
	4, // 0: infraboard.cmdb.disk.Disk.base:type_name -> infraboard.cmdb.resource.Base
	5, // 1: infraboard.cmdb.disk.Disk.information:type_name -> infraboard.cmdb.resource.Information
	1, // 2: infraboard.cmdb.disk.Disk.describe:type_name -> infraboard.cmdb.disk.Describe
	0, // 3: infraboard.cmdb.disk.Set.items:type_name -> infraboard.cmdb.disk.Disk
	6, // 4: infraboard.cmdb.disk.QueryDiskRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 5: infraboard.cmdb.disk.Service.SyncDisk:input_type -> infraboard.cmdb.disk.Disk
	3, // 6: infraboard.cmdb.disk.Service.QueryDisk:input_type -> infraboard.cmdb.disk.QueryDiskRequest
	0, // 7: infraboard.cmdb.disk.Service.SyncDisk:output_type -> infraboard.cmdb.disk.Disk
	2, // 8: infraboard.cmdb.disk.Service.QueryDisk:output_type -> infraboard.cmdb.disk.Set
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
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
			switch v := v.(*Set); i {
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
		file_apps_disk_pb_disk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDiskRequest); i {
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
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_disk_pb_disk_proto_goTypes,
		DependencyIndexes: file_apps_disk_pb_disk_proto_depIdxs,
		MessageInfos:      file_apps_disk_pb_disk_proto_msgTypes,
	}.Build()
	File_apps_disk_pb_disk_proto = out.File
	file_apps_disk_pb_disk_proto_rawDesc = nil
	file_apps_disk_pb_disk_proto_goTypes = nil
	file_apps_disk_pb_disk_proto_depIdxs = nil
}
