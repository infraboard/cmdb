// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: apps/oss/pb/rpc.proto

package oss

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
	SyncBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*Bucket, error)
	QueryBucket(ctx context.Context, in *QueryBucketRequest, opts ...grpc.CallOption) (*BucketSet, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SyncBucket(ctx context.Context, in *Bucket, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.oss.Service/SyncBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryBucket(ctx context.Context, in *QueryBucketRequest, opts ...grpc.CallOption) (*BucketSet, error) {
	out := new(BucketSet)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.oss.Service/QueryBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SyncBucket(context.Context, *Bucket) (*Bucket, error)
	QueryBucket(context.Context, *QueryBucketRequest) (*BucketSet, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SyncBucket(context.Context, *Bucket) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncBucket not implemented")
}
func (UnimplementedServiceServer) QueryBucket(context.Context, *QueryBucketRequest) (*BucketSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBucket not implemented")
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

func _Service_SyncBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SyncBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.oss.Service/SyncBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SyncBucket(ctx, req.(*Bucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.oss.Service/QueryBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryBucket(ctx, req.(*QueryBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.cmdb.oss.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncBucket",
			Handler:    _Service_SyncBucket_Handler,
		},
		{
			MethodName: "QueryBucket",
			Handler:    _Service_QueryBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/oss/pb/rpc.proto",
}