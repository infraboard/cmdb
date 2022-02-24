// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/cost/pb/cost.proto

package cost

import (
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

// 成本单元类型
type Type int32

const (
	// 基于域统计的总账单
	Type_DOMAIN Type = 0
	// 基于空间的总账单
	Type_NAMESPACE Type = 1
	// 基于云商的总账单
	Type_VENDOR Type = 2
	// 基于云商账号的总账单
	Type_ACCOUNT Type = 3
	// 基于环境的总账单
	Type_ENV Type = 4
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "DOMAIN",
		1: "NAMESPACE",
		2: "VENDOR",
		3: "ACCOUNT",
		4: "ENV",
	}
	Type_value = map[string]int32{
		"DOMAIN":    0,
		"NAMESPACE": 1,
		"VENDOR":    2,
		"ACCOUNT":   3,
		"ENV":       4,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_cost_pb_cost_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_apps_cost_pb_cost_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_apps_cost_pb_cost_proto_rawDescGZIP(), []int{0}
}

// 用于成本聚合统计的单元
type Unit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 财务单元类型
	// @gotags: json:"type"
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=infraboard.cmdb.cost.Type" json:"type"`
	// 财务单元具体类型的值
	// @gotags: json:"value"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	// 年份
	// @gotags: json:"year"
	Year int32 `protobuf:"varint,3,opt,name=year,proto3" json:"year"`
	// 月份
	// @gotags: json:"month"
	Month int32 `protobuf:"varint,4,opt,name=month,proto3" json:"month"`
	// 月账单
	// @gotags: json:"real_cost"
	RealCost float64 `protobuf:"fixed64,5,opt,name=real_cost,json=realCost,proto3" json:"real_cost"`
	// 相比上月上涨或下降的具体金额
	// @gotags: json:"delta_cost"
	DeltaCost float64 `protobuf:"fixed64,6,opt,name=delta_cost,json=deltaCost,proto3" json:"delta_cost"`
	// 相比上月上涨或下降的百分比
	// @gotags: json:"delta_percent"
	DeltaPercent float64 `protobuf:"fixed64,7,opt,name=delta_percent,json=deltaPercent,proto3" json:"delta_percent"`
	// 关联的具体详情Task
	// @gotags: json:"task_id"
	TaskId float64 `protobuf:"fixed64,15,opt,name=task_id,json=taskId,proto3" json:"task_id"`
}

func (x *Unit) Reset() {
	*x = Unit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cost_pb_cost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unit) ProtoMessage() {}

func (x *Unit) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cost_pb_cost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unit.ProtoReflect.Descriptor instead.
func (*Unit) Descriptor() ([]byte, []int) {
	return file_apps_cost_pb_cost_proto_rawDescGZIP(), []int{0}
}

func (x *Unit) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_DOMAIN
}

func (x *Unit) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Unit) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Unit) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Unit) GetRealCost() float64 {
	if x != nil {
		return x.RealCost
	}
	return 0
}

func (x *Unit) GetDeltaCost() float64 {
	if x != nil {
		return x.DeltaCost
	}
	return 0
}

func (x *Unit) GetDeltaPercent() float64 {
	if x != nil {
		return x.DeltaPercent
	}
	return 0
}

func (x *Unit) GetTaskId() float64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

var File_apps_cost_pb_cost_proto protoreflect.FileDescriptor

var file_apps_cost_pb_cost_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x22,
	0xf0, 0x01, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x6c,
	0x74, 0x61, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x2a, 0x43, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f,
	0x4d, 0x41, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50,
	0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x45, 0x4e, 0x56, 0x10, 0x04, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_cost_pb_cost_proto_rawDescOnce sync.Once
	file_apps_cost_pb_cost_proto_rawDescData = file_apps_cost_pb_cost_proto_rawDesc
)

func file_apps_cost_pb_cost_proto_rawDescGZIP() []byte {
	file_apps_cost_pb_cost_proto_rawDescOnce.Do(func() {
		file_apps_cost_pb_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_cost_pb_cost_proto_rawDescData)
	})
	return file_apps_cost_pb_cost_proto_rawDescData
}

var file_apps_cost_pb_cost_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_cost_pb_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apps_cost_pb_cost_proto_goTypes = []interface{}{
	(Type)(0),    // 0: infraboard.cmdb.cost.Type
	(*Unit)(nil), // 1: infraboard.cmdb.cost.Unit
}
var file_apps_cost_pb_cost_proto_depIdxs = []int32{
	0, // 0: infraboard.cmdb.cost.Unit.type:type_name -> infraboard.cmdb.cost.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_cost_pb_cost_proto_init() }
func file_apps_cost_pb_cost_proto_init() {
	if File_apps_cost_pb_cost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_cost_pb_cost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unit); i {
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
			RawDescriptor: file_apps_cost_pb_cost_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_cost_pb_cost_proto_goTypes,
		DependencyIndexes: file_apps_cost_pb_cost_proto_depIdxs,
		EnumInfos:         file_apps_cost_pb_cost_proto_enumTypes,
		MessageInfos:      file_apps_cost_pb_cost_proto_msgTypes,
	}.Build()
	File_apps_cost_pb_cost_proto = out.File
	file_apps_cost_pb_cost_proto_rawDesc = nil
	file_apps_cost_pb_cost_proto_goTypes = nil
	file_apps_cost_pb_cost_proto_depIdxs = nil
}
