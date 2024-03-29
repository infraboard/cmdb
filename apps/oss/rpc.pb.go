// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/oss/pb/rpc.proto

package oss

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

type QueryBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryBucketRequest) Reset() {
	*x = QueryBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_oss_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBucketRequest) ProtoMessage() {}

func (x *QueryBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_oss_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBucketRequest.ProtoReflect.Descriptor instead.
func (*QueryBucketRequest) Descriptor() ([]byte, []int) {
	return file_apps_oss_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryBucketRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_apps_oss_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_oss_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x1a, 0x15, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x32, 0xad, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x0a, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73,
	0x73, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x6f, 0x73, 0x73, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x74, 0x22,
	0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x6f, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_oss_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_oss_pb_rpc_proto_rawDescData = file_apps_oss_pb_rpc_proto_rawDesc
)

func file_apps_oss_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_oss_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_oss_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_oss_pb_rpc_proto_rawDescData)
	})
	return file_apps_oss_pb_rpc_proto_rawDescData
}

var file_apps_oss_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apps_oss_pb_rpc_proto_goTypes = []interface{}{
	(*QueryBucketRequest)(nil),  // 0: infraboard.cmdb.oss.QueryBucketRequest
	(*request.PageRequest)(nil), // 1: infraboard.mcube.page.PageRequest
	(*Bucket)(nil),              // 2: infraboard.cmdb.oss.Bucket
	(*BucketSet)(nil),           // 3: infraboard.cmdb.oss.BucketSet
}
var file_apps_oss_pb_rpc_proto_depIdxs = []int32{
	1, // 0: infraboard.cmdb.oss.QueryBucketRequest.page:type_name -> infraboard.mcube.page.PageRequest
	2, // 1: infraboard.cmdb.oss.Service.SyncBucket:input_type -> infraboard.cmdb.oss.Bucket
	0, // 2: infraboard.cmdb.oss.Service.QueryBucket:input_type -> infraboard.cmdb.oss.QueryBucketRequest
	2, // 3: infraboard.cmdb.oss.Service.SyncBucket:output_type -> infraboard.cmdb.oss.Bucket
	3, // 4: infraboard.cmdb.oss.Service.QueryBucket:output_type -> infraboard.cmdb.oss.BucketSet
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_oss_pb_rpc_proto_init() }
func file_apps_oss_pb_rpc_proto_init() {
	if File_apps_oss_pb_rpc_proto != nil {
		return
	}
	file_apps_oss_pb_oss_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_oss_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBucketRequest); i {
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
			RawDescriptor: file_apps_oss_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_oss_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_oss_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_oss_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_oss_pb_rpc_proto = out.File
	file_apps_oss_pb_rpc_proto_rawDesc = nil
	file_apps_oss_pb_rpc_proto_goTypes = nil
	file_apps_oss_pb_rpc_proto_depIdxs = nil
}
