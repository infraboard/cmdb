// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: apps/oss/pb/oss.proto

package oss

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

type OSS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源元数据
	// @gotags: json:"base"
	Base *resource.Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base"`
	// 资源基础信息
	// @gotags: json:"information"
	Information *resource.Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information"`
	// Oss详细信息
	// @gotags: json:"describe"
	Describe *Describe `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe"`
}

func (x *OSS) Reset() {
	*x = OSS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_oss_pb_oss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSS) ProtoMessage() {}

func (x *OSS) ProtoReflect() protoreflect.Message {
	mi := &file_apps_oss_pb_oss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSS.ProtoReflect.Descriptor instead.
func (*OSS) Descriptor() ([]byte, []int) {
	return file_apps_oss_pb_oss_proto_rawDescGZIP(), []int{0}
}

func (x *OSS) GetBase() *resource.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OSS) GetInformation() *resource.Information {
	if x != nil {
		return x.Information
	}
	return nil
}

func (x *OSS) GetDescribe() *Describe {
	if x != nil {
		return x.Describe
	}
	return nil
}

type Describe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Describe) Reset() {
	*x = Describe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_oss_pb_oss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Describe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Describe) ProtoMessage() {}

func (x *Describe) ProtoReflect() protoreflect.Message {
	mi := &file_apps_oss_pb_oss_proto_msgTypes[1]
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
	return file_apps_oss_pb_oss_proto_rawDescGZIP(), []int{1}
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
	Items []*OSS `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_oss_pb_oss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_apps_oss_pb_oss_proto_msgTypes[2]
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
	return file_apps_oss_pb_oss_proto_rawDescGZIP(), []int{2}
}

func (x *Set) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Set) GetItems() []*OSS {
	if x != nil {
		return x.Items
	}
	return nil
}

type QueryOSSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryOSSRequest) Reset() {
	*x = QueryOSSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_oss_pb_oss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOSSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOSSRequest) ProtoMessage() {}

func (x *QueryOSSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_oss_pb_oss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOSSRequest.ProtoReflect.Descriptor instead.
func (*QueryOSSRequest) Descriptor() ([]byte, []int) {
	return file_apps_oss_pb_oss_proto_rawDescGZIP(), []int{3}
}

func (x *QueryOSSRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_apps_oss_pb_oss_proto protoreflect.FileDescriptor

var file_apps_oss_pb_oss_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x1a, 0x1f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01,
	0x0a, 0x03, 0x4f, 0x53, 0x53, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x0a, 0x0a,
	0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x4b, 0x0a, 0x03, 0x53, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x4f, 0x53, 0x53, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x49, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f,
	0x53, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x32, 0x98, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a,
	0x07, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x53, 0x53, 0x12, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x4f,
	0x53, 0x53, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x4f, 0x53, 0x53, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x53, 0x53, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x53, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x6f, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_oss_pb_oss_proto_rawDescOnce sync.Once
	file_apps_oss_pb_oss_proto_rawDescData = file_apps_oss_pb_oss_proto_rawDesc
)

func file_apps_oss_pb_oss_proto_rawDescGZIP() []byte {
	file_apps_oss_pb_oss_proto_rawDescOnce.Do(func() {
		file_apps_oss_pb_oss_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_oss_pb_oss_proto_rawDescData)
	})
	return file_apps_oss_pb_oss_proto_rawDescData
}

var file_apps_oss_pb_oss_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_oss_pb_oss_proto_goTypes = []interface{}{
	(*OSS)(nil),                  // 0: infraboard.cmdb.oss.OSS
	(*Describe)(nil),             // 1: infraboard.cmdb.oss.Describe
	(*Set)(nil),                  // 2: infraboard.cmdb.oss.Set
	(*QueryOSSRequest)(nil),      // 3: infraboard.cmdb.oss.QueryOSSRequest
	(*resource.Base)(nil),        // 4: infraboard.cmdb.resource.Base
	(*resource.Information)(nil), // 5: infraboard.cmdb.resource.Information
	(*request.PageRequest)(nil),  // 6: infraboard.mcube.page.PageRequest
}
var file_apps_oss_pb_oss_proto_depIdxs = []int32{
	4, // 0: infraboard.cmdb.oss.OSS.base:type_name -> infraboard.cmdb.resource.Base
	5, // 1: infraboard.cmdb.oss.OSS.information:type_name -> infraboard.cmdb.resource.Information
	1, // 2: infraboard.cmdb.oss.OSS.describe:type_name -> infraboard.cmdb.oss.Describe
	0, // 3: infraboard.cmdb.oss.Set.items:type_name -> infraboard.cmdb.oss.OSS
	6, // 4: infraboard.cmdb.oss.QueryOSSRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 5: infraboard.cmdb.oss.Service.SyncOSS:input_type -> infraboard.cmdb.oss.OSS
	3, // 6: infraboard.cmdb.oss.Service.QueryOSS:input_type -> infraboard.cmdb.oss.QueryOSSRequest
	0, // 7: infraboard.cmdb.oss.Service.SyncOSS:output_type -> infraboard.cmdb.oss.OSS
	2, // 8: infraboard.cmdb.oss.Service.QueryOSS:output_type -> infraboard.cmdb.oss.Set
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_oss_pb_oss_proto_init() }
func file_apps_oss_pb_oss_proto_init() {
	if File_apps_oss_pb_oss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_oss_pb_oss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSS); i {
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
		file_apps_oss_pb_oss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_oss_pb_oss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_oss_pb_oss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOSSRequest); i {
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
			RawDescriptor: file_apps_oss_pb_oss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_oss_pb_oss_proto_goTypes,
		DependencyIndexes: file_apps_oss_pb_oss_proto_depIdxs,
		MessageInfos:      file_apps_oss_pb_oss_proto_msgTypes,
	}.Build()
	File_apps_oss_pb_oss_proto = out.File
	file_apps_oss_pb_oss_proto_rawDesc = nil
	file_apps_oss_pb_oss_proto_goTypes = nil
	file_apps_oss_pb_oss_proto_depIdxs = nil
}
