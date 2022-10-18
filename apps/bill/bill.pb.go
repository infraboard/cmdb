// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/bill/pb/bill.proto

package bill

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

// 账单数据结构参数:
// 阿里云参考: https://www.alibabacloud.com/help/zh/doc-detail/100400.htm?spm=a2c63.p38356.b99.39.22b56c48uGGnNz
// 腾讯云参考: https://cloud.tencent.com/document/api/555/19182
// 华为云参考:
type Bill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账单项目Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 本次账单同步关联的任务Id
	// @gotags: json:"task_id"
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor resource.VENDOR `protobuf:"varint,3,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.VENDOR" json:"vendor"`
	// 资源类型
	// @gotags: json:"resource_type"
	ResourceType resource.TYPE `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=infraboard.cmdb.resource.TYPE" json:"resource_type"`
	// 账单的年份
	// @gotags: json:"year"
	Year string `protobuf:"bytes,20,opt,name=year,proto3" json:"year"`
	// 账单的月份
	// @gotags: json:"month"
	Month string `protobuf:"bytes,5,opt,name=month,proto3" json:"month"`
	// 账单日
	// @gotags: json:"day"
	Day string `protobuf:"bytes,19,opt,name=day,proto3" json:"day"`
	// 账号Id
	// @gotags: json:"owner_id"
	OwnerId string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id"`
	// 账号名称
	// @gotags: json:"owner_name"
	OwnerName string `protobuf:"bytes,7,opt,name=owner_name,json=ownerName,proto3" json:"owner_name"`
	// 产品类型
	// @gotags: json:"product_type"
	ProductType string `protobuf:"bytes,8,opt,name=product_type,json=productType,proto3" json:"product_type"`
	// 产品编码
	// @gotags: json:"product_code"
	ProductCode string `protobuf:"bytes,9,opt,name=product_code,json=productCode,proto3" json:"product_code"`
	// 产品明细
	// @gotags: json:"product_detail"
	ProductDetail string `protobuf:"bytes,10,opt,name=product_detail,json=productDetail,proto3" json:"product_detail"`
	// 计费方式，比如：月95，日均峰值
	// @gotags: json:"pay_mode"
	PayMode resource.PayMode `protobuf:"varint,11,opt,name=pay_mode,json=payMode,proto3,enum=infraboard.cmdb.resource.PayMode" json:"pay_mode"`
	// 计费方式详情  云服务器ECS-按量付费
	// @gotags: json:"pay_mode_detail"
	PayModeDetail string `protobuf:"bytes,12,opt,name=pay_mode_detail,json=payModeDetail,proto3" json:"pay_mode_detail"`
	// 订单/账单ID
	// @gotags: json:"order_id"
	OrderId string `protobuf:"bytes,13,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	// 实例ID
	// @gotags: json:"instance_id"
	InstanceId string `protobuf:"bytes,14,opt,name=instance_id,json=instanceId,proto3" json:"instance_id"`
	// 实例名称
	// @gotags: json:"instance_name"
	InstanceName string `protobuf:"bytes,15,opt,name=instance_name,json=instanceName,proto3" json:"instance_name"`
	// 实例配置信息
	// @gotags: json:"instance_config"
	InstanceConfig string `protobuf:"bytes,16,opt,name=instance_config,json=instanceConfig,proto3" json:"instance_config"`
	// 地域Id
	// @gotags: json:"region_code"
	RegionCode string `protobuf:"bytes,17,opt,name=region_code,json=regionCode,proto3" json:"region_code"`
	// 地域名称
	// @gotags: json:"region_name"
	RegionName string `protobuf:"bytes,18,opt,name=region_name,json=regionName,proto3" json:"region_name"`
	// 实例所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,21,opt,name=domain,proto3" json:"domain"`
	// 实例所属空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,22,opt,name=namespace,proto3" json:"namespace"`
	// 资源所属环境
	// @gotags: json:"env"
	Env string `protobuf:"bytes,23,opt,name=env,proto3" json:"env"`
	// 该资产的基于Tag的分摊
	// @gotags: json:"share_units"
	ShareUnits []*ShareUnit `protobuf:"bytes,24,rep,name=share_units,json=shareUnits,proto3" json:"share_units"`
	// 具体费用
	// @gotags: json:"cost"
	Cost *Cost `protobuf:"bytes,25,opt,name=cost,proto3" json:"cost"`
	// 是否已经按照实例聚合
	// @gotags: json:"is_merged"
	IsMerged bool `protobuf:"varint,26,opt,name=is_merged,json=isMerged,proto3" json:"is_merged"`
	// 额外的无法通用的一些属性, 比如只有腾讯云独有的一些属性
	// @gotags: json:"extra"
	Extra map[string]string `protobuf:"bytes,35,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Bill) Reset() {
	*x = Bill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bill) ProtoMessage() {}

func (x *Bill) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bill.ProtoReflect.Descriptor instead.
func (*Bill) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{0}
}

func (x *Bill) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bill) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Bill) GetVendor() resource.VENDOR {
	if x != nil {
		return x.Vendor
	}
	return resource.VENDOR(0)
}

func (x *Bill) GetResourceType() resource.TYPE {
	if x != nil {
		return x.ResourceType
	}
	return resource.TYPE(0)
}

func (x *Bill) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Bill) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *Bill) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *Bill) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Bill) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Bill) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

func (x *Bill) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Bill) GetProductDetail() string {
	if x != nil {
		return x.ProductDetail
	}
	return ""
}

func (x *Bill) GetPayMode() resource.PayMode {
	if x != nil {
		return x.PayMode
	}
	return resource.PayMode(0)
}

func (x *Bill) GetPayModeDetail() string {
	if x != nil {
		return x.PayModeDetail
	}
	return ""
}

func (x *Bill) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Bill) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Bill) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *Bill) GetInstanceConfig() string {
	if x != nil {
		return x.InstanceConfig
	}
	return ""
}

func (x *Bill) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *Bill) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *Bill) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Bill) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Bill) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Bill) GetShareUnits() []*ShareUnit {
	if x != nil {
		return x.ShareUnits
	}
	return nil
}

func (x *Bill) GetCost() *Cost {
	if x != nil {
		return x.Cost
	}
	return nil
}

func (x *Bill) GetIsMerged() bool {
	if x != nil {
		return x.IsMerged
	}
	return false
}

func (x *Bill) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

// 资源月账单分摊单元
type ShareUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账单的月份
	// @gotags: json:"month"
	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month"`
	// 账单月当时标签
	// @gotags: json:"tag"
	Tag *resource.Tag `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag"`
	// 该成本单元的总权重, 用于计算分摊比例
	// @gotags: json:"total_weight"
	TotalWeight int64 `protobuf:"varint,3,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	// 分摊逻辑描述, 比如 instance_cost * percent
	// @gotags: json:"share_describe"
	ShareDescribe string `protobuf:"bytes,4,opt,name=share_describe,json=shareDescribe,proto3" json:"share_describe"`
	// 具体分摊额度, 具体金额根据标签权重进行计算
	// @gotags: json:"cost"
	Cost *Cost `protobuf:"bytes,5,opt,name=cost,proto3" json:"cost"`
}

func (x *ShareUnit) Reset() {
	*x = ShareUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareUnit) ProtoMessage() {}

func (x *ShareUnit) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_bill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareUnit.ProtoReflect.Descriptor instead.
func (*ShareUnit) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{1}
}

func (x *ShareUnit) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *ShareUnit) GetTag() *resource.Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *ShareUnit) GetTotalWeight() int64 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *ShareUnit) GetShareDescribe() string {
	if x != nil {
		return x.ShareDescribe
	}
	return ""
}

func (x *ShareUnit) GetCost() *Cost {
	if x != nil {
		return x.Cost
	}
	return nil
}

type BillSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总条数
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 费用统计
	// @gotags: json:"sum"
	Sum *Cost `protobuf:"bytes,2,opt,name=sum,proto3" json:"sum"`
	// 具体条目
	// @gotags: json:"items"
	Items []*Bill `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
}

func (x *BillSet) Reset() {
	*x = BillSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillSet) ProtoMessage() {}

func (x *BillSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_bill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillSet.ProtoReflect.Descriptor instead.
func (*BillSet) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{2}
}

func (x *BillSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BillSet) GetSum() *Cost {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *BillSet) GetItems() []*Bill {
	if x != nil {
		return x.Items
	}
	return nil
}

type Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 官网价
	// @gotags: json:"sale_price"
	SalePrice float64 `protobuf:"fixed64,1,opt,name=sale_price,json=salePrice,proto3" json:"sale_price"`
	// 优惠金额
	// @gotags: json:"save_cost"
	SaveCost float64 `protobuf:"fixed64,2,opt,name=save_cost,json=saveCost,proto3" json:"save_cost"`
	// 应付金额
	// @gotags: json:"real_cost"
	RealCost float64 `protobuf:"fixed64,3,opt,name=real_cost,json=realCost,proto3" json:"real_cost"`
	// 信用额度支付金额
	// @gotags: json:"credit_pay"
	CreditPay float64 `protobuf:"fixed64,4,opt,name=credit_pay,json=creditPay,proto3" json:"credit_pay"`
	// 代金券抵扣
	// 腾讯对应: VoucherPayAmount
	// 阿里云: DeductedByCashCoupons
	VoucherPay float64 `protobuf:"fixed64,5,opt,name=voucher_pay,json=voucherPay,proto3" json:"voucher_pay,omitempty"`
	// 现金抵扣
	// @gotags: json:"cash_pay"
	CashPay float64 `protobuf:"fixed64,6,opt,name=cash_pay,json=cashPay,proto3" json:"cash_pay"`
	// 储值卡抵扣
	// @gotags: json:"storedcard_pay"
	StoredcardPay float64 `protobuf:"fixed64,7,opt,name=storedcard_pay,json=storedcardPay,proto3" json:"storedcard_pay"`
	// 欠费金额
	// @gotags: json:"outstanding_amount"
	OutstandingAmount float64 `protobuf:"fixed64,8,opt,name=outstanding_amount,json=outstandingAmount,proto3" json:"outstanding_amount"`
}

func (x *Cost) Reset() {
	*x = Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_bill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cost.ProtoReflect.Descriptor instead.
func (*Cost) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{3}
}

func (x *Cost) GetSalePrice() float64 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *Cost) GetSaveCost() float64 {
	if x != nil {
		return x.SaveCost
	}
	return 0
}

func (x *Cost) GetRealCost() float64 {
	if x != nil {
		return x.RealCost
	}
	return 0
}

func (x *Cost) GetCreditPay() float64 {
	if x != nil {
		return x.CreditPay
	}
	return 0
}

func (x *Cost) GetVoucherPay() float64 {
	if x != nil {
		return x.VoucherPay
	}
	return 0
}

func (x *Cost) GetCashPay() float64 {
	if x != nil {
		return x.CashPay
	}
	return 0
}

func (x *Cost) GetStoredcardPay() float64 {
	if x != nil {
		return x.StoredcardPay
	}
	return 0
}

func (x *Cost) GetOutstandingAmount() float64 {
	if x != nil {
		return x.OutstandingAmount
	}
	return 0
}

type SummaryRecordSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总条数
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 费用统计
	// @gotags: json:"sum"
	Sum *Cost `protobuf:"bytes,2,opt,name=sum,proto3" json:"sum"`
	// 具体条目
	// @gotags: json:"items"
	Items []*SummaryRecord `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
}

func (x *SummaryRecordSet) Reset() {
	*x = SummaryRecordSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryRecordSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryRecordSet) ProtoMessage() {}

func (x *SummaryRecordSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_bill_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryRecordSet.ProtoReflect.Descriptor instead.
func (*SummaryRecordSet) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{4}
}

func (x *SummaryRecordSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SummaryRecordSet) GetSum() *Cost {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *SummaryRecordSet) GetItems() []*SummaryRecord {
	if x != nil {
		return x.Items
	}
	return nil
}

// 账单总览项
type SummaryRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账号Id
	// @gotags: json:"owner_id"
	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id"`
	// 账号名称
	// @gotags: json:"owner_name"
	OwnerName string `protobuf:"bytes,2,opt,name=owner_name,json=ownerName,proto3" json:"owner_name"`
	// 产品类型
	// @gotags: json:"product_type"
	ProductType string `protobuf:"bytes,3,opt,name=product_type,json=productType,proto3" json:"product_type"`
	// 产品编码
	// @gotags: json:"product_code"
	ProductCode string `protobuf:"bytes,4,opt,name=product_code,json=productCode,proto3" json:"product_code"`
	// 产品明细
	// @gotags: json:"product_detail"
	ProductDetail string `protobuf:"bytes,5,opt,name=product_detail,json=productDetail,proto3" json:"product_detail"`
	// 费用统计
	// @gotags: json:"sum"
	Sum *Cost `protobuf:"bytes,6,opt,name=sum,proto3" json:"sum"`
}

func (x *SummaryRecord) Reset() {
	*x = SummaryRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummaryRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryRecord) ProtoMessage() {}

func (x *SummaryRecord) ProtoReflect() protoreflect.Message {
	mi := &file_apps_bill_pb_bill_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryRecord.ProtoReflect.Descriptor instead.
func (*SummaryRecord) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{5}
}

func (x *SummaryRecord) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *SummaryRecord) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *SummaryRecord) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

func (x *SummaryRecord) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *SummaryRecord) GetProductDetail() string {
	if x != nil {
		return x.ProductDetail
	}
	return ""
}

func (x *SummaryRecord) GetSum() *Cost {
	if x != nil {
		return x.Sum
	}
	return nil
}

var File_apps_bill_pb_bill_proto protoreflect.FileDescriptor

var file_apps_bill_pb_bill_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x1a,
	0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x91, 0x08, 0x0a, 0x04, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x45,
	0x4e, 0x44, 0x4f, 0x52, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54,
	0x59, 0x50, 0x45, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x61, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x50, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x6e, 0x76, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x40,
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x18, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x12, 0x3b, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xcc, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x61, 0x72, 0x65, 0x55, 0x6e,
	0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x63,
	0x6f, 0x73, 0x74, 0x22, 0x7f, 0x0a, 0x07, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x03, 0x73,
	0x75, 0x6d, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x90, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x61, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x73, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x5f, 0x70, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x50, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x76, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x50, 0x61, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x73, 0x68, 0x5f, 0x70,
	0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x63, 0x61, 0x73, 0x68, 0x50, 0x61,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x63, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x03, 0x73, 0x75, 0x6d,
	0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x0d,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x03, 0x73,
	0x75, 0x6d, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x62,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_bill_pb_bill_proto_rawDescOnce sync.Once
	file_apps_bill_pb_bill_proto_rawDescData = file_apps_bill_pb_bill_proto_rawDesc
)

func file_apps_bill_pb_bill_proto_rawDescGZIP() []byte {
	file_apps_bill_pb_bill_proto_rawDescOnce.Do(func() {
		file_apps_bill_pb_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_bill_pb_bill_proto_rawDescData)
	})
	return file_apps_bill_pb_bill_proto_rawDescData
}

var file_apps_bill_pb_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_bill_pb_bill_proto_goTypes = []interface{}{
	(*Bill)(nil),             // 0: infraboard.cmdb.bill.Bill
	(*ShareUnit)(nil),        // 1: infraboard.cmdb.bill.ShareUnit
	(*BillSet)(nil),          // 2: infraboard.cmdb.bill.BillSet
	(*Cost)(nil),             // 3: infraboard.cmdb.bill.Cost
	(*SummaryRecordSet)(nil), // 4: infraboard.cmdb.bill.SummaryRecordSet
	(*SummaryRecord)(nil),    // 5: infraboard.cmdb.bill.SummaryRecord
	nil,                      // 6: infraboard.cmdb.bill.Bill.ExtraEntry
	(resource.VENDOR)(0),     // 7: infraboard.cmdb.resource.VENDOR
	(resource.TYPE)(0),       // 8: infraboard.cmdb.resource.TYPE
	(resource.PayMode)(0),    // 9: infraboard.cmdb.resource.PayMode
	(*resource.Tag)(nil),     // 10: infraboard.cmdb.resource.Tag
}
var file_apps_bill_pb_bill_proto_depIdxs = []int32{
	7,  // 0: infraboard.cmdb.bill.Bill.vendor:type_name -> infraboard.cmdb.resource.VENDOR
	8,  // 1: infraboard.cmdb.bill.Bill.resource_type:type_name -> infraboard.cmdb.resource.TYPE
	9,  // 2: infraboard.cmdb.bill.Bill.pay_mode:type_name -> infraboard.cmdb.resource.PayMode
	1,  // 3: infraboard.cmdb.bill.Bill.share_units:type_name -> infraboard.cmdb.bill.ShareUnit
	3,  // 4: infraboard.cmdb.bill.Bill.cost:type_name -> infraboard.cmdb.bill.Cost
	6,  // 5: infraboard.cmdb.bill.Bill.extra:type_name -> infraboard.cmdb.bill.Bill.ExtraEntry
	10, // 6: infraboard.cmdb.bill.ShareUnit.tag:type_name -> infraboard.cmdb.resource.Tag
	3,  // 7: infraboard.cmdb.bill.ShareUnit.cost:type_name -> infraboard.cmdb.bill.Cost
	3,  // 8: infraboard.cmdb.bill.BillSet.sum:type_name -> infraboard.cmdb.bill.Cost
	0,  // 9: infraboard.cmdb.bill.BillSet.items:type_name -> infraboard.cmdb.bill.Bill
	3,  // 10: infraboard.cmdb.bill.SummaryRecordSet.sum:type_name -> infraboard.cmdb.bill.Cost
	5,  // 11: infraboard.cmdb.bill.SummaryRecordSet.items:type_name -> infraboard.cmdb.bill.SummaryRecord
	3,  // 12: infraboard.cmdb.bill.SummaryRecord.sum:type_name -> infraboard.cmdb.bill.Cost
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_apps_bill_pb_bill_proto_init() }
func file_apps_bill_pb_bill_proto_init() {
	if File_apps_bill_pb_bill_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_bill_pb_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bill); i {
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
		file_apps_bill_pb_bill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareUnit); i {
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
		file_apps_bill_pb_bill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillSet); i {
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
		file_apps_bill_pb_bill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cost); i {
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
		file_apps_bill_pb_bill_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryRecordSet); i {
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
		file_apps_bill_pb_bill_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummaryRecord); i {
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
			RawDescriptor: file_apps_bill_pb_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_bill_pb_bill_proto_goTypes,
		DependencyIndexes: file_apps_bill_pb_bill_proto_depIdxs,
		MessageInfos:      file_apps_bill_pb_bill_proto_msgTypes,
	}.Build()
	File_apps_bill_pb_bill_proto = out.File
	file_apps_bill_pb_bill_proto_rawDesc = nil
	file_apps_bill_pb_bill_proto_goTypes = nil
	file_apps_bill_pb_bill_proto_depIdxs = nil
}
