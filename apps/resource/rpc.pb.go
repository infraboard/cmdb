// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/resource/pb/rpc.proto

package resource

import (
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

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源Id列表
	// @gotags: json:"resource_ids"
	ResourceIds []string `protobuf:"bytes,1,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteRequest) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 删除记录
	// @gotags: json:"delete_records"
	DeleteRecords []*DeleteRecord `protobuf:"bytes,1,rep,name=delete_records,json=deleteRecords,proto3" json:"delete_records"`
	// 删除失败统计
	// @gotags: json:"failed_count"
	FailedCount int64 `protobuf:"varint,2,opt,name=failed_count,json=failedCount,proto3" json:"failed_count"`
	// 删除成功统计
	// @gotags: json:"success_count"
	SuccessCount int64 `protobuf:"varint,3,opt,name=success_count,json=successCount,proto3" json:"success_count"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteResponse) GetDeleteRecords() []*DeleteRecord {
	if x != nil {
		return x.DeleteRecords
	}
	return nil
}

func (x *DeleteResponse) GetFailedCount() int64 {
	if x != nil {
		return x.FailedCount
	}
	return 0
}

func (x *DeleteResponse) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

type DeleteRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 资源描述
	// @gotags: json:"describe"
	Describe string `protobuf:"bytes,2,opt,name=describe,proto3" json:"describe"`
	// 资源是否删除成功
	// @gotags: json:"success"
	Success bool `protobuf:"varint,3,opt,name=success,proto3" json:"success"`
	// 删除失败的原因
	// @gotags: json:"reason"
	Reason string `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason"`
}

func (x *DeleteRecord) Reset() {
	*x = DeleteRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecord) ProtoMessage() {}

func (x *DeleteRecord) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecord.ProtoReflect.Descriptor instead.
func (*DeleteRecord) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRecord) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *DeleteRecord) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteRecord) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 资源所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 资源所属空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace"`
	// 资源所属环境
	// @gotags: json:"env"
	Env string `protobuf:"bytes,4,opt,name=env,proto3" json:"env"`
	// 使用方式
	// @gotags: json:"usage_mode"
	UsageMode *USAGE_MODE `protobuf:"varint,5,opt,name=usage_mode,json=usageMode,proto3,enum=infraboard.cmdb.resource.USAGE_MODE,oneof" json:"usage_mode"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor *VENDOR `protobuf:"varint,6,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.VENDOR,oneof" json:"vendor"`
	// 资源所属账号
	// @gotags: json:"owner"
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner"`
	// 资源类型
	// @gotags: json:"type"
	Type *TYPE `protobuf:"varint,8,opt,name=type,proto3,enum=infraboard.cmdb.resource.TYPE,oneof" json:"type"`
	// 服务商中的状态
	// @gotags: json:"status"
	Status string `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	// 资源标签
	// @gotags: json:"tags"
	Tags []*TagSelector `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags"`
	// 是否返回资源的标签
	// @gotags: json:"with_tags"
	WithTags bool `protobuf:"varint,11,opt,name=with_tags,json=withTags,proto3" json:"with_tags"`
	// 关联资源查询深度, -1表示查询所有, 默认为0,表示不查询关联资源
	// @gotags: json:"relation_deep"
	RelationDeep int32 `protobuf:"varint,12,opt,name=relation_deep,json=relationDeep,proto3" json:"relation_deep"`
	// 关键字参数
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,14,opt,name=keywords,proto3" json:"keywords"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *SearchRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *SearchRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SearchRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SearchRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *SearchRequest) GetUsageMode() USAGE_MODE {
	if x != nil && x.UsageMode != nil {
		return *x.UsageMode
	}
	return USAGE_MODE_SHARED
}

func (x *SearchRequest) GetVendor() VENDOR {
	if x != nil && x.Vendor != nil {
		return *x.Vendor
	}
	return VENDOR_ALIYUN
}

func (x *SearchRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SearchRequest) GetType() TYPE {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return TYPE_HOST
}

func (x *SearchRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SearchRequest) GetTags() []*TagSelector {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *SearchRequest) GetWithTags() bool {
	if x != nil {
		return x.WithTags
	}
	return false
}

func (x *SearchRequest) GetRelationDeep() int32 {
	if x != nil {
		return x.RelationDeep
	}
	return 0
}

func (x *SearchRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type TagSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 匹配的key, 支持LIKE匹配
	// @gotags: json:"key" validate:"required"
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key" validate:"required"`
	// 匹配符
	// @gotags: json:"opertor"
	Opertor string `protobuf:"bytes,2,opt,name=opertor,proto3" json:"opertor"`
	// 匹配的值
	// @gotags: json:"values"
	Values []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values"`
}

func (x *TagSelector) Reset() {
	*x = TagSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagSelector) ProtoMessage() {}

func (x *TagSelector) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagSelector.ProtoReflect.Descriptor instead.
func (*TagSelector) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *TagSelector) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TagSelector) GetOpertor() string {
	if x != nil {
		return x.Opertor
	}
	return ""
}

func (x *TagSelector) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_apps_resource_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_resource_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0d,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0xbb, 0x04, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x48, 0x0a, 0x0a, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x48,
	0x00, 0x52, 0x09, 0x75, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x3d, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x45, 0x4e, 0x44, 0x4f,
	0x52, 0x48, 0x01, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x59,
	0x50, 0x45, 0x48, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54,
	0x61, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x54, 0x61, 0x67, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x65, 0x70, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x65, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x51, 0x0a, 0x0b, 0x54, 0x61, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x32, 0x5f, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x58, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63,
	0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_resource_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_resource_pb_rpc_proto_rawDescData = file_apps_resource_pb_rpc_proto_rawDesc
)

func file_apps_resource_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_resource_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_resource_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_resource_pb_rpc_proto_rawDescData)
	})
	return file_apps_resource_pb_rpc_proto_rawDescData
}

var file_apps_resource_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_resource_pb_rpc_proto_goTypes = []interface{}{
	(*DeleteRequest)(nil),       // 0: infraboard.cmdb.resource.DeleteRequest
	(*DeleteResponse)(nil),      // 1: infraboard.cmdb.resource.DeleteResponse
	(*DeleteRecord)(nil),        // 2: infraboard.cmdb.resource.DeleteRecord
	(*SearchRequest)(nil),       // 3: infraboard.cmdb.resource.SearchRequest
	(*TagSelector)(nil),         // 4: infraboard.cmdb.resource.TagSelector
	(*request.PageRequest)(nil), // 5: infraboard.mcube.page.PageRequest
	(USAGE_MODE)(0),             // 6: infraboard.cmdb.resource.USAGE_MODE
	(VENDOR)(0),                 // 7: infraboard.cmdb.resource.VENDOR
	(TYPE)(0),                   // 8: infraboard.cmdb.resource.TYPE
	(*ResourceSet)(nil),         // 9: infraboard.cmdb.resource.ResourceSet
}
var file_apps_resource_pb_rpc_proto_depIdxs = []int32{
	2, // 0: infraboard.cmdb.resource.DeleteResponse.delete_records:type_name -> infraboard.cmdb.resource.DeleteRecord
	5, // 1: infraboard.cmdb.resource.SearchRequest.page:type_name -> infraboard.mcube.page.PageRequest
	6, // 2: infraboard.cmdb.resource.SearchRequest.usage_mode:type_name -> infraboard.cmdb.resource.USAGE_MODE
	7, // 3: infraboard.cmdb.resource.SearchRequest.vendor:type_name -> infraboard.cmdb.resource.VENDOR
	8, // 4: infraboard.cmdb.resource.SearchRequest.type:type_name -> infraboard.cmdb.resource.TYPE
	4, // 5: infraboard.cmdb.resource.SearchRequest.tags:type_name -> infraboard.cmdb.resource.TagSelector
	3, // 6: infraboard.cmdb.resource.RPC.Search:input_type -> infraboard.cmdb.resource.SearchRequest
	9, // 7: infraboard.cmdb.resource.RPC.Search:output_type -> infraboard.cmdb.resource.ResourceSet
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_resource_pb_rpc_proto_init() }
func file_apps_resource_pb_rpc_proto_init() {
	if File_apps_resource_pb_rpc_proto != nil {
		return
	}
	file_apps_resource_pb_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_resource_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_apps_resource_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_apps_resource_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRecord); i {
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
		file_apps_resource_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_apps_resource_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagSelector); i {
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
	file_apps_resource_pb_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_resource_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_resource_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_resource_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_resource_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_resource_pb_rpc_proto = out.File
	file_apps_resource_pb_rpc_proto_rawDesc = nil
	file_apps_resource_pb_rpc_proto_goTypes = nil
	file_apps_resource_pb_rpc_proto_depIdxs = nil
}