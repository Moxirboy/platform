// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: backend/api/proto/admin.proto

package proto

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

const (
	CreateTeacher_CreateUser_FullMethodName = "/admin.CreateTeacher/CreateUser"
)

// CreateTeacherClient is the client API for CreateTeacher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateTeacherClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type createTeacherClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateTeacherClient(cc grpc.ClientConnInterface) CreateTeacherClient {
	return &createTeacherClient{cc}
}

func (c *createTeacherClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, CreateTeacher_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateTeacherServer is the server API for CreateTeacher service.
// All implementations must embed UnimplementedCreateTeacherServer
// for forward compatibility
type CreateTeacherServer interface {
	CreateUser(context.Context, *User) (*ErrorResponse, error)
	mustEmbedUnimplementedCreateTeacherServer()
}

// UnimplementedCreateTeacherServer must be embedded to have forward compatible implementations.
type UnimplementedCreateTeacherServer struct {
}

func (UnimplementedCreateTeacherServer) CreateUser(context.Context, *User) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedCreateTeacherServer) mustEmbedUnimplementedCreateTeacherServer() {}

// UnsafeCreateTeacherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateTeacherServer will
// result in compilation errors.
type UnsafeCreateTeacherServer interface {
	mustEmbedUnimplementedCreateTeacherServer()
}

func RegisterCreateTeacherServer(s grpc.ServiceRegistrar, srv CreateTeacherServer) {
	s.RegisterService(&CreateTeacher_ServiceDesc, srv)
}

func _CreateTeacher_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateTeacherServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateTeacher_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateTeacherServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateTeacher_ServiceDesc is the grpc.ServiceDesc for CreateTeacher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateTeacher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.CreateTeacher",
	HandlerType: (*CreateTeacherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _CreateTeacher_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/proto/admin.proto",
}