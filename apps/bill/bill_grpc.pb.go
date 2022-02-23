// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: apps/bill/pb/bill.proto

package bill

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	SyncBill(ctx context.Context, in *Bill, opts ...grpc.CallOption) (*Bill, error)
	QueryBill(ctx context.Context, in *QueryBillRequest, opts ...grpc.CallOption) (*BillSet, error)
	ConfirmBill(ctx context.Context, in *ConfirmBillRequest, opts ...grpc.CallOption) (*BillSet, error)
	DeleteBill(ctx context.Context, in *DeleteBillRequest, opts ...grpc.CallOption) (*BillSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SyncBill(ctx context.Context, in *Bill, opts ...grpc.CallOption) (*Bill, error) {
	out := new(Bill)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.bill.Service/SyncBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryBill(ctx context.Context, in *QueryBillRequest, opts ...grpc.CallOption) (*BillSet, error) {
	out := new(BillSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.bill.Service/QueryBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ConfirmBill(ctx context.Context, in *ConfirmBillRequest, opts ...grpc.CallOption) (*BillSet, error) {
	out := new(BillSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.bill.Service/ConfirmBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteBill(ctx context.Context, in *DeleteBillRequest, opts ...grpc.CallOption) (*BillSet, error) {
	out := new(BillSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.bill.Service/DeleteBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SyncBill(context.Context, *Bill) (*Bill, error)
	QueryBill(context.Context, *QueryBillRequest) (*BillSet, error)
	ConfirmBill(context.Context, *ConfirmBillRequest) (*BillSet, error)
	DeleteBill(context.Context, *DeleteBillRequest) (*BillSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SyncBill(context.Context, *Bill) (*Bill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncBill not implemented")
}
func (UnimplementedServiceServer) QueryBill(context.Context, *QueryBillRequest) (*BillSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBill not implemented")
}
func (UnimplementedServiceServer) ConfirmBill(context.Context, *ConfirmBillRequest) (*BillSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmBill not implemented")
}
func (UnimplementedServiceServer) DeleteBill(context.Context, *DeleteBillRequest) (*BillSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBill not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_SyncBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bill)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SyncBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.bill.Service/SyncBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SyncBill(ctx, req.(*Bill))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.bill.Service/QueryBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryBill(ctx, req.(*QueryBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ConfirmBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ConfirmBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.bill.Service/ConfirmBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ConfirmBill(ctx, req.(*ConfirmBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.bill.Service/DeleteBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteBill(ctx, req.(*DeleteBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.cmdb.bill.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncBill",
			Handler:    _Service_SyncBill_Handler,
		},
		{
			MethodName: "QueryBill",
			Handler:    _Service_QueryBill_Handler,
		},
		{
			MethodName: "ConfirmBill",
			Handler:    _Service_ConfirmBill_Handler,
		},
		{
			MethodName: "DeleteBill",
			Handler:    _Service_DeleteBill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/bill/pb/bill.proto",
}
