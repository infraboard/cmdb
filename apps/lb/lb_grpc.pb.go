// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: apps/lb/pb/lb.proto

package lb

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
	SyncSLB(ctx context.Context, in *LoadBalancer, opts ...grpc.CallOption) (*LoadBalancer, error)
	QuerySLB(ctx context.Context, in *QuerySLBRequest, opts ...grpc.CallOption) (*LoadBalancer, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SyncSLB(ctx context.Context, in *LoadBalancer, opts ...grpc.CallOption) (*LoadBalancer, error) {
	out := new(LoadBalancer)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.slb.Service/SyncSLB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QuerySLB(ctx context.Context, in *QuerySLBRequest, opts ...grpc.CallOption) (*LoadBalancer, error) {
	out := new(LoadBalancer)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.slb.Service/QuerySLB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SyncSLB(context.Context, *LoadBalancer) (*LoadBalancer, error)
	QuerySLB(context.Context, *QuerySLBRequest) (*LoadBalancer, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SyncSLB(context.Context, *LoadBalancer) (*LoadBalancer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncSLB not implemented")
}
func (UnimplementedServiceServer) QuerySLB(context.Context, *QuerySLBRequest) (*LoadBalancer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySLB not implemented")
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

func _Service_SyncSLB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadBalancer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SyncSLB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.slb.Service/SyncSLB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SyncSLB(ctx, req.(*LoadBalancer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QuerySLB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySLBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QuerySLB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.slb.Service/QuerySLB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QuerySLB(ctx, req.(*QuerySLBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.cmdb.slb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncSLB",
			Handler:    _Service_SyncSLB_Handler,
		},
		{
			MethodName: "QuerySLB",
			Handler:    _Service_QuerySLB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/lb/pb/lb.proto",
}
