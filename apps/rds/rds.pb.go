// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/rds/pb/rds.proto

package rds

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

type RDS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源元数据信息
	// @gotags: json:"base"
	Base *resource.Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base"`
	// 资源基础信息
	// @gotags: json:"information"
	Information *resource.Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information"`
	// Rds描述信息
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe"`
}

func (x *RDS) Reset() {
	*x = RDS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rds_pb_rds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RDS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RDS) ProtoMessage() {}

func (x *RDS) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rds_pb_rds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RDS.ProtoReflect.Descriptor instead.
func (*RDS) Descriptor() ([]byte, []int) {
	return file_apps_rds_pb_rds_proto_rawDescGZIP(), []int{0}
}

func (x *RDS) GetBase() *resource.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *RDS) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *RDS) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 引擎 比如 MYSQL, SQLServer, PGSQL
	// @gotags: json:"engine_type"
	EngineType string `protobuf:"bytes,2,opt,name=engine_type,json=engineType,proto3" json:"engine_type"`
	// 引擎版本
	// @gotags: json:"engine_version"
	EngineVersion string `protobuf:"bytes,3,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version"`
	// 实例规格: 对应ALI(DBInstanceClass)
	// @gotags: json:"instance_class"
	InstanceClass string `protobuf:"bytes,4,opt,name=instance_class,json=instanceClass,proto3" json:"instance_class"`
	// 实例规格族，取值：s：共享型；x：通用型；d：独享套餐；h：独占物理机。
	// @gotags: json:"class_type"
	ClassType string `protobuf:"bytes,5,opt,name=class_type,json=classType,proto3" json:"class_type"`
	// 实例是内网或外网 (Internet：外网/Intranet：内网)
	// @gotags: json:"export_type"
	ExportType string `protobuf:"bytes,6,opt,name=export_type,json=exportType,proto3" json:"export_type"`
	// 实例的网络类型 (Classic：经典网络/VPC：专有网络。)
	// @gotags: json:"network_type"
	NetworkType string `protobuf:"bytes,7,opt,name=network_type,json=networkType,proto3" json:"network_type"`
	// 实例类型 Primary：主实例, Readonly：只读实例, Guard：灾备实例, Temp：临时实例
	// @gotags: json:"type"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type"`
	// CPU 核数
	// @gotags: json:"cpu"
	Cpu int32 `protobuf:"varint,9,opt,name=cpu,proto3" json:"cpu"`
	// 实例内存，单位：M。
	// @gotags: json:"memory"
	Memory int64 `protobuf:"varint,10,opt,name=memory,proto3" json:"memory"`
	// 一个实例下可创建最大数据库数量
	// @gotags: json:"db_max_quantity"
	DbMaxQuantity int64 `protobuf:"varint,11,opt,name=db_max_quantity,json=dbMaxQuantity,proto3" json:"db_max_quantity"`
	// 可创建账号的最大数量
	// @gotags: json:"account_max_quantity"
	AccountMaxQuantity int64 `protobuf:"varint,12,opt,name=account_max_quantity,json=accountMaxQuantity,proto3" json:"account_max_quantity"`
	// 最大并发连接数
	// @gotags: json:"max_connections"
	MaxConnections int64 `protobuf:"varint,13,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections"`
	// 最大每秒IO请求次数
	// @gotags: json:"max_iops"
	MaxIops int64 `protobuf:"varint,14,opt,name=max_iops,json=maxIops,proto3" json:"max_iops"`
	// 系统字符集排序规则
	// @gotags: json:"collation"
	Collation string `protobuf:"bytes,15,opt,name=collation,proto3" json:"collation"`
	// 时区
	// @gotags: json:"time_zone"
	TimeZone string `protobuf:"bytes,16,opt,name=time_zone,json=timeZone,proto3" json:"time_zone"`
	// 实例存储空间，单位：GB。
	// @gotags: json:"storage_capacity"
	StorageCapacity int64 `protobuf:"varint,17,opt,name=storage_capacity,json=storageCapacity,proto3" json:"storage_capacity"`
	// 实例储存类型 local_ssd/ephemeral_ssd：本地SSD盘, cloud_ssd：SSD云盘；cloud_essd：ESSD云盘
	// @gotags: json:"storage_type"
	StorageType string `protobuf:"bytes,18,opt,name=storage_type,json=storageType,proto3" json:"storage_type"`
	// 安全名单模式, 默认白名单
	// @gotags: json:"security_ip_mode"
	SecurityIpMode string `protobuf:"bytes,19,opt,name=security_ip_mode,json=securityIpMode,proto3" json:"security_ip_mode"`
	// IP白名单
	// @gotags: json:"security_ip_list"
	SecurityIpList []string `protobuf:"bytes,20,rep,name=security_ip_list,json=securityIpList,proto3" json:"security_ip_list"`
	// 实例的访问模式，取值：Standard：标准访问模式；Safe：数据库代理模式。
	// @gotags: json:"connection_mode"
	ConnectionMode string `protobuf:"bytes,21,opt,name=connection_mode,json=connectionMode,proto3" json:"connection_mode"`
	// IP类型
	// @gotags: json:"ip_type"
	IpType string `protobuf:"bytes,22,opt,name=ip_type,json=ipType,proto3" json:"ip_type"`
	// 实例锁定模式; Unlock：正常；ManualLock：手动触发锁定；LockByExpiration：实例过期自动锁定；LockByRestoration：实例回滚前的自动锁定；LockByDiskQuota：实例空间满自动锁定
	// @gotags: json:"lock_mode"
	LockMode string `protobuf:"bytes,23,opt,name=lock_mode,json=lockMode,proto3" json:"lock_mode"`
	// 锁定原因
	// @gotags: json:"lock_reason"
	LockReason string `protobuf:"bytes,24,opt,name=lock_reason,json=lockReason,proto3" json:"lock_reason"`
	// 部署模式(腾讯云独有)
	// @gotags: json:"deploy_mode"
	DeployMode string `protobuf:"bytes,25,opt,name=deploy_mode,json=deployMode,proto3" json:"deploy_mode"`
	// 端口
	// @gotags: json:"port"
	Port int64 `protobuf:"varint,26,opt,name=port,proto3" json:"port"`
	// 额外的无法通用的一些属性, 比如只有腾讯云独有的一些属性
	// @gotags: json:"extra"
	Extra map[string]string `protobuf:"bytes,27,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rds_pb_rds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rds_pb_rds_proto_msgTypes[1]
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
	return file_apps_rds_pb_rds_proto_rawDescGZIP(), []int{1}
}

func (x *Describe) GetEngineType() string {
	if x != nil {
		return x.EngineType
	}
	return ""
}

func (x *Describe) GetEngineVersion() string {
	if x != nil {
		return x.EngineVersion
	}
	return ""
}

func (x *Describe) GetInstanceClass() string {
	if x != nil {
		return x.InstanceClass
	}
	return ""
}

func (x *Describe) GetClassType() string {
	if x != nil {
		return x.ClassType
	}
	return ""
}

func (x *Describe) GetExportType() string {
	if x != nil {
		return x.ExportType
	}
	return ""
}

func (x *Describe) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

func (x *Describe) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Describe) GetCpu() int32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Describe) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Describe) GetDbMaxQuantity() int64 {
	if x != nil {
		return x.DbMaxQuantity
	}
	return 0
}

func (x *Describe) GetAccountMaxQuantity() int64 {
	if x != nil {
		return x.AccountMaxQuantity
	}
	return 0
}

func (x *Describe) GetMaxConnections() int64 {
	if x != nil {
		return x.MaxConnections
	}
	return 0
}

func (x *Describe) GetMaxIops() int64 {
	if x != nil {
		return x.MaxIops
	}
	return 0
}

func (x *Describe) GetCollation() string {
	if x != nil {
		return x.Collation
	}
	return ""
}

func (x *Describe) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *Describe) GetStorageCapacity() int64 {
	if x != nil {
		return x.StorageCapacity
	}
	return 0
}

func (x *Describe) GetStorageType() string {
	if x != nil {
		return x.StorageType
	}
	return ""
}

func (x *Describe) GetSecurityIpMode() string {
	if x != nil {
		return x.SecurityIpMode
	}
	return ""
}

func (x *Describe) GetSecurityIpList() []string {
	if x != nil {
		return x.SecurityIpList
	}
	return nil
}

func (x *Describe) GetConnectionMode() string {
	if x != nil {
		return x.ConnectionMode
	}
	return ""
}

func (x *Describe) GetIpType() string {
	if x != nil {
		return x.IpType
	}
	return ""
}

func (x *Describe) GetLockMode() string {
	if x != nil {
		return x.LockMode
	}
	return ""
}

func (x *Describe) GetLockReason() string {
	if x != nil {
		return x.LockReason
	}
	return ""
}

func (x *Describe) GetDeployMode() string {
	if x != nil {
		return x.DeployMode
	}
	return ""
}

func (x *Describe) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Describe) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
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
	Items []*RDS `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rds_pb_rds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rds_pb_rds_proto_msgTypes[2]
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
	return file_apps_rds_pb_rds_proto_rawDescGZIP(), []int{2}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*RDS {
	if x != nil {
		return x.Items
	}
	return nil
}

type QueryRDSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryRDSRequest) Reset() {
	*x = QueryRDSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rds_pb_rds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRDSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRDSRequest) ProtoMessage() {}

func (x *QueryRDSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rds_pb_rds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRDSRequest.ProtoReflect.Descriptor instead.
func (*QueryRDSRequest) Descriptor() ([]byte, []int) {
	return file_apps_rds_pb_rds_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRDSRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_apps_rds_pb_rds_proto protoreflect.FileDescriptor

var file_apps_rds_pb_rds_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x1a, 0x1f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01,
	0x0a, 0x03, 0x52, 0x44, 0x53, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0xc4, 0x07,
	0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63,
	0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x62,
	0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x62, 0x4d, 0x61, 0x78, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x61,
	0x78, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x78, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6f, 0x70, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6d, 0x61, 0x78, 0x49, 0x6f, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a,
	0x6f, 0x6e, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x70,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x44, 0x53, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x49, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x44, 0x53, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x98, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x44, 0x53, 0x12, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x44, 0x53, 0x1a, 0x18, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x72, 0x64, 0x73, 0x2e, 0x52, 0x44, 0x53, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x08, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x44, 0x53, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x44, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x64,
	0x73, 0x2e, 0x53, 0x65, 0x74, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_rds_pb_rds_proto_rawDescOnce sync.Once
	file_apps_rds_pb_rds_proto_rawDescData = file_apps_rds_pb_rds_proto_rawDesc
)

func file_apps_rds_pb_rds_proto_rawDescGZIP() []byte {
	file_apps_rds_pb_rds_proto_rawDescOnce.Do(func() {
		file_apps_rds_pb_rds_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_rds_pb_rds_proto_rawDescData)
	})
	return file_apps_rds_pb_rds_proto_rawDescData
}

var file_apps_rds_pb_rds_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_rds_pb_rds_proto_goTypes = []interface{}{
	(*RDS)(nil),                  // 0: infraboard.cmdb.rds.RDS
	(*Describe)(nil),             // 1: infraboard.cmdb.rds.Describe
	(*Set)(nil),                  // 2: infraboard.cmdb.rds.Set
	(*QueryRDSRequest)(nil),      // 3: infraboard.cmdb.rds.QueryRDSRequest
	nil,                          // 4: infraboard.cmdb.rds.Describe.ExtraEntry
	(*resource.Base)(nil),        // 5: infraboard.cmdb.resource.Base
	(*resource.Information)(nil), // 6: infraboard.cmdb.resource.Information
	(*request.PageRequest)(nil),  // 7: infraboard.mcube.page.PageRequest
}
var file_apps_rds_pb_rds_proto_depIdxs = []int32{
	5, // 0: infraboard.cmdb.rds.RDS.base:type_name -> infraboard.cmdb.resource.Base
	6, // 1: infraboard.cmdb.rds.RDS.information:type_name -> infraboard.cmdb.resource.Information
	1, // 2: infraboard.cmdb.rds.RDS.describe:type_name -> infraboard.cmdb.rds.Describe
	4, // 3: infraboard.cmdb.rds.Describe.extra:type_name -> infraboard.cmdb.rds.Describe.ExtraEntry
	0, // 4: infraboard.cmdb.rds.Set.items:type_name -> infraboard.cmdb.rds.RDS
	7, // 5: infraboard.cmdb.rds.QueryRDSRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 6: infraboard.cmdb.rds.Service.SyncRDS:input_type -> infraboard.cmdb.rds.RDS
	3, // 7: infraboard.cmdb.rds.Service.QueryRDS:input_type -> infraboard.cmdb.rds.QueryRDSRequest
	0, // 8: infraboard.cmdb.rds.Service.SyncRDS:output_type -> infraboard.cmdb.rds.RDS
	2, // 9: infraboard.cmdb.rds.Service.QueryRDS:output_type -> infraboard.cmdb.rds.Set
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_rds_pb_rds_proto_init() }
func file_apps_rds_pb_rds_proto_init() {
	if File_apps_rds_pb_rds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_rds_pb_rds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RDS); i {
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
		file_apps_rds_pb_rds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_rds_pb_rds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_rds_pb_rds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRDSRequest); i {
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
			RawDescriptor: file_apps_rds_pb_rds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_rds_pb_rds_proto_goTypes,
		DependencyIndexes: file_apps_rds_pb_rds_proto_depIdxs,
		MessageInfos:      file_apps_rds_pb_rds_proto_msgTypes,
	}.Build()
	File_apps_rds_pb_rds_proto = out.File
	file_apps_rds_pb_rds_proto_rawDesc = nil
	file_apps_rds_pb_rds_proto_goTypes = nil
	file_apps_rds_pb_rds_proto_depIdxs = nil
}
