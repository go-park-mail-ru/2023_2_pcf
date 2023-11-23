// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: select.proto

package api

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

// SelectClient is the client API for Select service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SelectClient interface {
	Get(ctx context.Context, in *SelectRequests, opts ...grpc.CallOption) (*SelectResponse, error)
}

type selectClient struct {
	cc grpc.ClientConnInterface
}

func NewSelectClient(cc grpc.ClientConnInterface) SelectClient {
	return &selectClient{cc}
}

func (c *selectClient) Get(ctx context.Context, in *SelectRequests, opts ...grpc.CallOption) (*SelectResponse, error) {
	out := new(SelectResponse)
	err := c.cc.Invoke(ctx, "/proto.Select/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SelectServer is the server API for Select service.
// All implementations must embed UnimplementedSelectServer
// for forward compatibility
type SelectServer interface {
	Get(context.Context, *SelectRequests) (*SelectResponse, error)
	mustEmbedUnimplementedSelectServer()
}

// UnimplementedSelectServer must be embedded to have forward compatible implementations.
type UnimplementedSelectServer struct {
}

func (UnimplementedSelectServer) Get(context.Context, *SelectRequests) (*SelectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSelectServer) mustEmbedUnimplementedSelectServer() {}

// UnsafeSelectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SelectServer will
// result in compilation errors.
type UnsafeSelectServer interface {
	mustEmbedUnimplementedSelectServer()
}

func RegisterSelectServer(s grpc.ServiceRegistrar, srv SelectServer) {
	s.RegisterService(&Select_ServiceDesc, srv)
}

func _Select_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelectServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Select/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelectServer).Get(ctx, req.(*SelectRequests))
	}
	return interceptor(ctx, in, info, handler)
}

// Select_ServiceDesc is the grpc.ServiceDesc for Select service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Select_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Select",
	HandlerType: (*SelectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Select_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "select.proto",
}