// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/mongodb/pb/mongodb.proto

package mongodb

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
	// 表示创建失败
	STATUS_CREATE_FAILED STATUS = 6
	// 表示运行中
	STATUS_RUNNING STATUS = 11
	// 实例配置变更生效中
	STATUS_MODIFYING STATUS = 20
	// 表示重启中
	STATUS_REBOOTING STATUS = 30
	// 备份恢复中
	STATUS_RESTORING STATUS = 31
	// 迁移中
	STATUS_TRANSING STATUS = 32
	// 数据导入中
	STATUS_IMPORTING STATUS = 33
	// 实例克隆中
	STATUS_CLONING STATUS = 34
	// 迁移版本中
	STATUS_UPGRADING STATUS = 35
	// 表示实例正在进行备份
	STATUS_BACKING_UP STATUS = 36
	// 内外网切换中
	STATUS_NET_CHANGING STATUS = 37
	// 表示实例正在主备切换
	STATUS_SWITCHOVER STATUS = 40
	// 灾备实例创建同步中
	STATUS_GUARD_CREATING STATUS = 41
	// 状态异常
	STATUS_ERROR STATUS = 50
	// 表示实例磁盘空间满
	STATUS_STORAGE_FULL STATUS = 51
	// 表示实例已经锁定
	STATUS_LOCKED STATUS = 70
	// 隔离中
	STATUS_ISOLATIONING STATUS = 71
	// 已隔中
	STATUS_ISOLATIONED STATUS = 72
	// 表示停止待销毁
	STATUS_SHUTDOWN STATUS = 80
	// 表示销毁中
	STATUS_DELETING STATUS = 81
	// 已销毁
	STATUS_DESTROYED STATUS = 90
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0:  "UNKNOW",
		5:  "PENDING",
		6:  "CREATE_FAILED",
		11: "RUNNING",
		20: "MODIFYING",
		30: "REBOOTING",
		31: "RESTORING",
		32: "TRANSING",
		33: "IMPORTING",
		34: "CLONING",
		35: "UPGRADING",
		36: "BACKING_UP",
		37: "NET_CHANGING",
		40: "SWITCHOVER",
		41: "GUARD_CREATING",
		50: "ERROR",
		51: "STORAGE_FULL",
		70: "LOCKED",
		71: "ISOLATIONING",
		72: "ISOLATIONED",
		80: "SHUTDOWN",
		81: "DELETING",
		90: "DESTROYED",
	}
	STATUS_value = map[string]int32{
		"UNKNOW":         0,
		"PENDING":        5,
		"CREATE_FAILED":  6,
		"RUNNING":        11,
		"MODIFYING":      20,
		"REBOOTING":      30,
		"RESTORING":      31,
		"TRANSING":       32,
		"IMPORTING":      33,
		"CLONING":        34,
		"UPGRADING":      35,
		"BACKING_UP":     36,
		"NET_CHANGING":   37,
		"SWITCHOVER":     40,
		"GUARD_CREATING": 41,
		"ERROR":          50,
		"STORAGE_FULL":   51,
		"LOCKED":         70,
		"ISOLATIONING":   71,
		"ISOLATIONED":    72,
		"SHUTDOWN":       80,
		"DELETING":       81,
		"DESTROYED":      90,
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
	return file_apps_mongodb_pb_mongodb_proto_enumTypes[0].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_apps_mongodb_pb_mongodb_proto_enumTypes[0]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_apps_mongodb_pb_mongodb_proto_rawDescGZIP(), []int{0}
}

type MongoDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"resource"
	Resource *resource.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource"`
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe"`
}

func (x *MongoDB) Reset() {
	*x = MongoDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_mongodb_pb_mongodb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoDB) ProtoMessage() {}

func (x *MongoDB) ProtoReflect() protoreflect.Message {
	mi := &file_apps_mongodb_pb_mongodb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoDB.ProtoReflect.Descriptor instead.
func (*MongoDB) Descriptor() ([]byte, []int) {
	return file_apps_mongodb_pb_mongodb_proto_rawDescGZIP(), []int{0}
}

func (x *MongoDB) GetResource() *resource.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *MongoDB) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据库引擎
	// @gotags: json:"engine"
	Engine string `protobuf:"bytes,1,opt,name=engine,proto3" json:"engine"`
	// 存储类型
	// @gotags: json:"storage_type"
	StorageType string `protobuf:"bytes,2,opt,name=storage_type,json=storageType,proto3" json:"storage_type"`
	// 数据库版本号
	// @gotags: json:"engine_version"
	EngineVersion string `protobuf:"bytes,3,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version"`
	// Mongos节点的最大连接数
	// @gotags: json:"max_connections"
	MaxConnections int32 `protobuf:"varint,4,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections"`
	// 实例最大IOPS
	// @gotags: json:"max_iops"
	MaxIops int32 `protobuf:"varint,5,opt,name=max_iops,json=maxIops,proto3" json:"max_iops"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_mongodb_pb_mongodb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_mongodb_pb_mongodb_proto_msgTypes[1]
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
	return file_apps_mongodb_pb_mongodb_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *Describe) GetStorageType() string {
	if x != nil {
		return x.StorageType
	}
	return ""
}

func (x *Describe) GetEngineVersion() string {
	if x != nil {
		return x.EngineVersion
	}
	return ""
}

func (x *Describe) GetMaxConnections() int32 {
	if x != nil {
		return x.MaxConnections
	}
	return 0
}

func (x *Describe) GetMaxIops() int32 {
	if x != nil {
		return x.MaxIops
	}
	return 0
}

type MongoDBSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表项
	// @gotags: json:"items"
	Items []*MongoDB `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *MongoDBSet) Reset() {
	*x = MongoDBSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_mongodb_pb_mongodb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoDBSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoDBSet) ProtoMessage() {}

func (x *MongoDBSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_mongodb_pb_mongodb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoDBSet.ProtoReflect.Descriptor instead.
func (*MongoDBSet) Descriptor() ([]byte, []int) {
	return file_apps_mongodb_pb_mongodb_proto_rawDescGZIP(), []int{2}
}

func (x *MongoDBSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MongoDBSet) GetItems() []*MongoDB {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_mongodb_pb_mongodb_proto protoreflect.FileDescriptor

var file_apps_mongodb_pb_mongodb_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x70,
	0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x4d, 0x6f,
	0x6e, 0x67, 0x6f, 0x44, 0x42, 0x12, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64,
	0x62, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61,
	0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6f, 0x70, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6d, 0x61, 0x78, 0x49, 0x6f, 0x70, 0x73, 0x22, 0x5a, 0x0a, 0x0a, 0x4d, 0x6f, 0x6e, 0x67, 0x6f,
	0x44, 0x42, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x64, 0x62, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x2a, 0xe4, 0x02, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x59, 0x49, 0x4e, 0x47, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x42, 0x4f, 0x4f, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x1e, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x1f, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x4e, 0x47,
	0x10, 0x20, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x21, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x22, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x23, 0x12, 0x0e, 0x0a,
	0x0a, 0x42, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x50, 0x10, 0x24, 0x12, 0x10, 0x0a,
	0x0c, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x25, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x28, 0x12,
	0x12, 0x0a, 0x0e, 0x47, 0x55, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x29, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x32, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x33,
	0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x46, 0x12, 0x10, 0x0a, 0x0c,
	0x49, 0x53, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x47, 0x12, 0x0f,
	0x0a, 0x0b, 0x49, 0x53, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x48, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x50, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x51, 0x12, 0x0d, 0x0a, 0x09, 0x44,
	0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x5a, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_mongodb_pb_mongodb_proto_rawDescOnce sync.Once
	file_apps_mongodb_pb_mongodb_proto_rawDescData = file_apps_mongodb_pb_mongodb_proto_rawDesc
)

func file_apps_mongodb_pb_mongodb_proto_rawDescGZIP() []byte {
	file_apps_mongodb_pb_mongodb_proto_rawDescOnce.Do(func() {
		file_apps_mongodb_pb_mongodb_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_mongodb_pb_mongodb_proto_rawDescData)
	})
	return file_apps_mongodb_pb_mongodb_proto_rawDescData
}

var file_apps_mongodb_pb_mongodb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_mongodb_pb_mongodb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_mongodb_pb_mongodb_proto_goTypes = []interface{}{
	(STATUS)(0),               // 0: infraboard.cmdb.mongodb.STATUS
	(*MongoDB)(nil),           // 1: infraboard.cmdb.mongodb.MongoDB
	(*Describe)(nil),          // 2: infraboard.cmdb.mongodb.Describe
	(*MongoDBSet)(nil),        // 3: infraboard.cmdb.mongodb.MongoDBSet
	(*resource.Resource)(nil), // 4: infraboard.cmdb.resource.Resource
}
var file_apps_mongodb_pb_mongodb_proto_depIdxs = []int32{
	4, // 0: infraboard.cmdb.mongodb.MongoDB.resource:type_name -> infraboard.cmdb.resource.Resource
	2, // 1: infraboard.cmdb.mongodb.MongoDB.describe:type_name -> infraboard.cmdb.mongodb.Describe
	1, // 2: infraboard.cmdb.mongodb.MongoDBSet.items:type_name -> infraboard.cmdb.mongodb.MongoDB
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_mongodb_pb_mongodb_proto_init() }
func file_apps_mongodb_pb_mongodb_proto_init() {
	if File_apps_mongodb_pb_mongodb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_mongodb_pb_mongodb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoDB); i {
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
		file_apps_mongodb_pb_mongodb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_mongodb_pb_mongodb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoDBSet); i {
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
			RawDescriptor: file_apps_mongodb_pb_mongodb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_mongodb_pb_mongodb_proto_goTypes,
		DependencyIndexes: file_apps_mongodb_pb_mongodb_proto_depIdxs,
		EnumInfos:         file_apps_mongodb_pb_mongodb_proto_enumTypes,
		MessageInfos:      file_apps_mongodb_pb_mongodb_proto_msgTypes,
	}.Build()
	File_apps_mongodb_pb_mongodb_proto = out.File
	file_apps_mongodb_pb_mongodb_proto_rawDesc = nil
	file_apps_mongodb_pb_mongodb_proto_goTypes = nil
	file_apps_mongodb_pb_mongodb_proto_depIdxs = nil
}
