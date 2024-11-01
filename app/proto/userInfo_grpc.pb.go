// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: userInfo.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Usr_GetUser_FullMethodName = "/proto.Usr/GetUser"
)

// UsrClient is the client API for Usr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsrClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type usrClient struct {
	cc grpc.ClientConnInterface
}

func NewUsrClient(cc grpc.ClientConnInterface) UsrClient {
	return &usrClient{cc}
}

func (c *usrClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, Usr_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsrServer is the server API for Usr service.
// All implementations must embed UnimplementedUsrServer
// for forward compatibility.
type UsrServer interface {
	GetUser(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUsrServer()
}

// UnimplementedUsrServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsrServer struct{}

func (UnimplementedUsrServer) GetUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUsrServer) mustEmbedUnimplementedUsrServer() {}
func (UnimplementedUsrServer) testEmbeddedByValue()             {}

// UnsafeUsrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsrServer will
// result in compilation errors.
type UnsafeUsrServer interface {
	mustEmbedUnimplementedUsrServer()
}

func RegisterUsrServer(s grpc.ServiceRegistrar, srv UsrServer) {
	// If the following call pancis, it indicates UnimplementedUsrServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Usr_ServiceDesc, srv)
}

func _Usr_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsrServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usr_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsrServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Usr_ServiceDesc is the grpc.ServiceDesc for Usr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Usr",
	HandlerType: (*UsrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Usr_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userInfo.proto",
}
