// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: model.proto

package __

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

// GuessTheNumberClient is the client API for GuessTheNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuessTheNumberClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Response, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Response, error)
	Play(ctx context.Context, in *GuessRequest, opts ...grpc.CallOption) (*GuessResponse, error)
}

type guessTheNumberClient struct {
	cc grpc.ClientConnInterface
}

func NewGuessTheNumberClient(cc grpc.ClientConnInterface) GuessTheNumberClient {
	return &guessTheNumberClient{cc}
}

func (c *guessTheNumberClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/api.GuessTheNumber/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guessTheNumberClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.GuessTheNumber/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guessTheNumberClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/api.GuessTheNumber/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guessTheNumberClient) GetAllUsers(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/api.GuessTheNumber/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guessTheNumberClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.GuessTheNumber/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guessTheNumberClient) Play(ctx context.Context, in *GuessRequest, opts ...grpc.CallOption) (*GuessResponse, error) {
	out := new(GuessResponse)
	err := c.cc.Invoke(ctx, "/api.GuessTheNumber/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuessTheNumberServer is the server API for GuessTheNumber service.
// All implementations must embed UnimplementedGuessTheNumberServer
// for forward compatibility
type GuessTheNumberServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*Response, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetAllUsers(context.Context, *GetAllRequest) (*GetAllResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*Response, error)
	Play(context.Context, *GuessRequest) (*GuessResponse, error)
	mustEmbedUnimplementedGuessTheNumberServer()
}

// UnimplementedGuessTheNumberServer must be embedded to have forward compatible implementations.
type UnimplementedGuessTheNumberServer struct {
}

func (UnimplementedGuessTheNumberServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGuessTheNumberServer) UpdateUser(context.Context, *UpdateUserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedGuessTheNumberServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedGuessTheNumberServer) GetAllUsers(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedGuessTheNumberServer) DeleteUser(context.Context, *DeleteUserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedGuessTheNumberServer) Play(context.Context, *GuessRequest) (*GuessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedGuessTheNumberServer) mustEmbedUnimplementedGuessTheNumberServer() {}

// UnsafeGuessTheNumberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuessTheNumberServer will
// result in compilation errors.
type UnsafeGuessTheNumberServer interface {
	mustEmbedUnimplementedGuessTheNumberServer()
}

func RegisterGuessTheNumberServer(s grpc.ServiceRegistrar, srv GuessTheNumberServer) {
	s.RegisterService(&GuessTheNumber_ServiceDesc, srv)
}

func _GuessTheNumber_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuessTheNumberServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GuessTheNumber/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuessTheNumberServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuessTheNumber_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuessTheNumberServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GuessTheNumber/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuessTheNumberServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuessTheNumber_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuessTheNumberServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GuessTheNumber/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuessTheNumberServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuessTheNumber_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuessTheNumberServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GuessTheNumber/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuessTheNumberServer).GetAllUsers(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuessTheNumber_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuessTheNumberServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GuessTheNumber/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuessTheNumberServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuessTheNumber_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuessTheNumberServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GuessTheNumber/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuessTheNumberServer).Play(ctx, req.(*GuessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GuessTheNumber_ServiceDesc is the grpc.ServiceDesc for GuessTheNumber service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuessTheNumber_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.GuessTheNumber",
	HandlerType: (*GuessTheNumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _GuessTheNumber_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _GuessTheNumber_UpdateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _GuessTheNumber_GetUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _GuessTheNumber_GetAllUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _GuessTheNumber_DeleteUser_Handler,
		},
		{
			MethodName: "Play",
			Handler:    _GuessTheNumber_Play_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}
