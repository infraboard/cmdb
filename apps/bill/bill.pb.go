// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: apps/bill/pb/bill.proto

package bill

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

// 账单数据结构参数:
// 阿里云参考: https://www.alibabacloud.com/help/zh/doc-detail/100400.htm?spm=a2c63.p38356.b99.39.22b56c48uGGnNz
// 腾讯云参考: https://cloud.tencent.com/document/api/555/19182
type Bill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 本次账单同步关联的任务Id
	// @gotags: json:"task_id"
	TaskId string `protobuf:"bytes,26,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor resource.Vendor `protobuf:"varint,1,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.Vendor" json:"vendor"`
	// 账单的月份
	// @gotags: json:"month"
	Month string `protobuf:"bytes,2,opt,name=month,proto3" json:"month"`
	// 账号Id
	// @gotags: json:"owner_id"
	OwnerId string `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id"`
	// 账号名称
	// @gotags: json:"owner_name"
	OwnerName string `protobuf:"bytes,4,opt,name=owner_name,json=ownerName,proto3" json:"owner_name"`
	// 产品类型
	// @gotags: json:"product_type"
	ProductType string `protobuf:"bytes,5,opt,name=product_type,json=productType,proto3" json:"product_type"`
	// 产品编码
	// @gotags: json:"product_code"
	ProductCode string `protobuf:"bytes,6,opt,name=product_code,json=productCode,proto3" json:"product_code"`
	// 产品明细
	// @gotags: json:"product_detail"
	ProductDetail string `protobuf:"bytes,7,opt,name=product_detail,json=productDetail,proto3" json:"product_detail"`
	// 计费方式，比如：月95，日均峰值
	// @gotags: json:"pay_mode"
	PayMode string `protobuf:"bytes,8,opt,name=pay_mode,json=payMode,proto3" json:"pay_mode"`
	// 计费方式详情  云服务器ECS-按量付费
	// @gotags: json:"pay_mode_detail"
	PayModeDetail string `protobuf:"bytes,9,opt,name=pay_mode_detail,json=payModeDetail,proto3" json:"pay_mode_detail"`
	// 订单/账单ID
	// @gotags: json:"order_id"
	OrderId string `protobuf:"bytes,10,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	// 实例ID
	// @gotags: json:"instance_id"
	InstanceId string `protobuf:"bytes,11,opt,name=instance_id,json=instanceId,proto3" json:"instance_id"`
	// 实例名称
	// @gotags: json:"instance_name"
	InstanceName string `protobuf:"bytes,12,opt,name=instance_name,json=instanceName,proto3" json:"instance_name"`
	// 公网Ip
	// @gotags: json:"public_ip"
	PublicIp string `protobuf:"bytes,13,opt,name=public_ip,json=publicIp,proto3" json:"public_ip"`
	// 内网Ip
	// @gotags: json:"private_ip"
	PrivateIp string `protobuf:"bytes,14,opt,name=private_ip,json=privateIp,proto3" json:"private_ip"`
	// 实例配置信息
	// @gotags: json:"instance_config"
	InstanceConfig string `protobuf:"bytes,15,opt,name=instance_config,json=instanceConfig,proto3" json:"instance_config"`
	// 地域Id
	// @gotags: json:"region_code"
	RegionCode string `protobuf:"bytes,16,opt,name=region_code,json=regionCode,proto3" json:"region_code"`
	// 地域名称
	// @gotags: json:"region_name"
	RegionName string `protobuf:"bytes,17,opt,name=region_name,json=regionName,proto3" json:"region_name"`
	// 实例所属空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,28,opt,name=namespace,proto3" json:"namespace"`
	// 资源所属环境
	// @gotags: json:"env"
	Env string `protobuf:"bytes,29,opt,name=env,proto3" json:"env"`
	// 账单月当时标签
	// @gotags: json:"tags"
	Tags []*resource.Tag `protobuf:"bytes,27,rep,name=tags,proto3" json:"tags"`
	// 官网价
	// @gotags: json:"sale_price"
	SalePrice float64 `protobuf:"fixed64,18,opt,name=sale_price,json=salePrice,proto3" json:"sale_price"`
	// 优惠金额
	// @gotags: json:"save_cost"
	SaveCost float64 `protobuf:"fixed64,19,opt,name=save_cost,json=saveCost,proto3" json:"save_cost"`
	// 应付金额
	// @gotags: json:"real_cost"
	RealCost float64 `protobuf:"fixed64,20,opt,name=real_cost,json=realCost,proto3" json:"real_cost"`
	// 信用额度支付金额
	// @gotags: json:"credit_pay"
	CreditPay float64 `protobuf:"fixed64,21,opt,name=credit_pay,json=creditPay,proto3" json:"credit_pay"`
	// 代金券抵扣
	// 腾讯对应: VoucherPayAmount
	// 阿里云: DeductedByCashCoupons
	VoucherPay float64 `protobuf:"fixed64,22,opt,name=voucher_pay,json=voucherPay,proto3" json:"voucher_pay,omitempty"`
	// 现金抵扣
	// @gotags: json:"cash_pay"
	CashPay float64 `protobuf:"fixed64,23,opt,name=cash_pay,json=cashPay,proto3" json:"cash_pay"`
	// 储值卡抵扣
	// @gotags: json:"storedcard_pay"
	StoredcardPay float64 `protobuf:"fixed64,24,opt,name=storedcard_pay,json=storedcardPay,proto3" json:"storedcard_pay"`
	// 欠费金额
	// @gotags: json:"outstanding_amount"
	OutstandingAmount float64 `protobuf:"fixed64,25,opt,name=outstanding_amount,json=outstandingAmount,proto3" json:"outstanding_amount"`
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

func (x *Bill) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Bill) GetVendor() resource.Vendor {
	if x != nil {
		return x.Vendor
	}
	return resource.Vendor(0)
}

func (x *Bill) GetMonth() string {
	if x != nil {
		return x.Month
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

func (x *Bill) GetPayMode() string {
	if x != nil {
		return x.PayMode
	}
	return ""
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

func (x *Bill) GetPublicIp() string {
	if x != nil {
		return x.PublicIp
	}
	return ""
}

func (x *Bill) GetPrivateIp() string {
	if x != nil {
		return x.PrivateIp
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

func (x *Bill) GetTags() []*resource.Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Bill) GetSalePrice() float64 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *Bill) GetSaveCost() float64 {
	if x != nil {
		return x.SaveCost
	}
	return 0
}

func (x *Bill) GetRealCost() float64 {
	if x != nil {
		return x.RealCost
	}
	return 0
}

func (x *Bill) GetCreditPay() float64 {
	if x != nil {
		return x.CreditPay
	}
	return 0
}

func (x *Bill) GetVoucherPay() float64 {
	if x != nil {
		return x.VoucherPay
	}
	return 0
}

func (x *Bill) GetCashPay() float64 {
	if x != nil {
		return x.CashPay
	}
	return 0
}

func (x *Bill) GetStoredcardPay() float64 {
	if x != nil {
		return x.StoredcardPay
	}
	return 0
}

func (x *Bill) GetOutstandingAmount() float64 {
	if x != nil {
		return x.OutstandingAmount
	}
	return 0
}

func (x *Bill) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
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
	// 具体条目
	// @gotags: json:"items"
	Items []*Bill `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *BillSet) Reset() {
	*x = BillSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillSet) ProtoMessage() {}

func (x *BillSet) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BillSet.ProtoReflect.Descriptor instead.
func (*BillSet) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{1}
}

func (x *BillSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BillSet) GetItems() []*Bill {
	if x != nil {
		return x.Items
	}
	return nil
}

type QueryBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 实例所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 实例所属空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace"`
	// 资源所属环境
	// @gotags: json:"env"
	Env string `protobuf:"bytes,4,opt,name=env,proto3" json:"env"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor resource.Vendor `protobuf:"varint,5,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.Vendor" json:"vendor"`
	// 账单的月份
	// @gotags: json:"month"
	Month string `protobuf:"bytes,6,opt,name=month,proto3" json:"month"`
	// 账号Id
	// @gotags: json:"account_id"
	AccountId string `protobuf:"bytes,7,opt,name=account_id,json=accountId,proto3" json:"account_id"`
	// 地域Id
	// @gotags: json:"region_code"
	RegionCode string `protobuf:"bytes,8,opt,name=region_code,json=regionCode,proto3" json:"region_code"`
	// 产品编码
	// @gotags: json:"product_code"
	ProductCode string `protobuf:"bytes,9,opt,name=product_code,json=productCode,proto3" json:"product_code"`
	// 账单月当时标签
	// @gotags: json:"tags"
	Tags []*resource.Tag `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags"`
	// 实例ID
	// @gotags: json:"instance_id"
	InstanceId string `protobuf:"bytes,11,opt,name=instance_id,json=instanceId,proto3" json:"instance_id"`
}

func (x *QueryBillRequest) Reset() {
	*x = QueryBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBillRequest) ProtoMessage() {}

func (x *QueryBillRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueryBillRequest.ProtoReflect.Descriptor instead.
func (*QueryBillRequest) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{2}
}

func (x *QueryBillRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryBillRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryBillRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryBillRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *QueryBillRequest) GetVendor() resource.Vendor {
	if x != nil {
		return x.Vendor
	}
	return resource.Vendor(0)
}

func (x *QueryBillRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *QueryBillRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *QueryBillRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *QueryBillRequest) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *QueryBillRequest) GetTags() []*resource.Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *QueryBillRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type DeleteBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账单的月份
	// @gotags: json:"task_id" validate:"required"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id" validate:"required"`
}

func (x *DeleteBillRequest) Reset() {
	*x = DeleteBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBillRequest) ProtoMessage() {}

func (x *DeleteBillRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteBillRequest.ProtoReflect.Descriptor instead.
func (*DeleteBillRequest) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteBillRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type ConfirmBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账单的月份
	// @gotags: json:"task_id" validate:"required"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id" validate:"required"`
}

func (x *ConfirmBillRequest) Reset() {
	*x = ConfirmBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_bill_pb_bill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmBillRequest) ProtoMessage() {}

func (x *ConfirmBillRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ConfirmBillRequest.ProtoReflect.Descriptor instead.
func (*ConfirmBillRequest) Descriptor() ([]byte, []int) {
	return file_apps_bill_pb_bill_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmBillRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_apps_bill_pb_bill_proto protoreflect.FileDescriptor

var file_apps_bill_pb_bill_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x1a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc5, 0x08, 0x0a, 0x04, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x6e, 0x76, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x1b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x61,
	0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x61, 0x6c,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x61, 0x76, 0x65, 0x43,
	0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x50, 0x61, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x50, 0x61, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x63, 0x61, 0x73, 0x68, 0x50, 0x61, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x63, 0x61, 0x72, 0x64, 0x50,
	0x61, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11,
	0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38,
	0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x51, 0x0a, 0x07, 0x42, 0x69, 0x6c, 0x6c,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e,
	0x42, 0x69, 0x6c, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x99, 0x03, 0x0a, 0x10,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76,
	0x12, 0x38, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x67,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x32, 0xcf, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e,
	0x42, 0x69, 0x6c, 0x6c, 0x12, 0x52, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x69, 0x6c,
	0x6c, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x56, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63,
	0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74,
	0x12, 0x54, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x27,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x53, 0x65, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_apps_bill_pb_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_bill_pb_bill_proto_goTypes = []interface{}{
	(*Bill)(nil),                // 0: infraboard.cmdb.bill.Bill
	(*BillSet)(nil),             // 1: infraboard.cmdb.bill.BillSet
	(*QueryBillRequest)(nil),    // 2: infraboard.cmdb.bill.QueryBillRequest
	(*DeleteBillRequest)(nil),   // 3: infraboard.cmdb.bill.DeleteBillRequest
	(*ConfirmBillRequest)(nil),  // 4: infraboard.cmdb.bill.ConfirmBillRequest
	nil,                         // 5: infraboard.cmdb.bill.Bill.ExtraEntry
	(resource.Vendor)(0),        // 6: infraboard.cmdb.resource.Vendor
	(*resource.Tag)(nil),        // 7: infraboard.cmdb.resource.Tag
	(*request.PageRequest)(nil), // 8: infraboard.mcube.page.PageRequest
}
var file_apps_bill_pb_bill_proto_depIdxs = []int32{
	6,  // 0: infraboard.cmdb.bill.Bill.vendor:type_name -> infraboard.cmdb.resource.Vendor
	7,  // 1: infraboard.cmdb.bill.Bill.tags:type_name -> infraboard.cmdb.resource.Tag
	5,  // 2: infraboard.cmdb.bill.Bill.extra:type_name -> infraboard.cmdb.bill.Bill.ExtraEntry
	0,  // 3: infraboard.cmdb.bill.BillSet.items:type_name -> infraboard.cmdb.bill.Bill
	8,  // 4: infraboard.cmdb.bill.QueryBillRequest.page:type_name -> infraboard.mcube.page.PageRequest
	6,  // 5: infraboard.cmdb.bill.QueryBillRequest.vendor:type_name -> infraboard.cmdb.resource.Vendor
	7,  // 6: infraboard.cmdb.bill.QueryBillRequest.tags:type_name -> infraboard.cmdb.resource.Tag
	0,  // 7: infraboard.cmdb.bill.Service.SyncBill:input_type -> infraboard.cmdb.bill.Bill
	2,  // 8: infraboard.cmdb.bill.Service.QueryBill:input_type -> infraboard.cmdb.bill.QueryBillRequest
	4,  // 9: infraboard.cmdb.bill.Service.ConfirmBill:input_type -> infraboard.cmdb.bill.ConfirmBillRequest
	3,  // 10: infraboard.cmdb.bill.Service.DeleteBill:input_type -> infraboard.cmdb.bill.DeleteBillRequest
	0,  // 11: infraboard.cmdb.bill.Service.SyncBill:output_type -> infraboard.cmdb.bill.Bill
	1,  // 12: infraboard.cmdb.bill.Service.QueryBill:output_type -> infraboard.cmdb.bill.BillSet
	1,  // 13: infraboard.cmdb.bill.Service.ConfirmBill:output_type -> infraboard.cmdb.bill.BillSet
	1,  // 14: infraboard.cmdb.bill.Service.DeleteBill:output_type -> infraboard.cmdb.bill.BillSet
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
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
		file_apps_bill_pb_bill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBillRequest); i {
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
			switch v := v.(*DeleteBillRequest); i {
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
			switch v := v.(*ConfirmBillRequest); i {
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
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
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
