// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package users

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

// UserServicesClient is the client API for UserServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServicesClient interface {
	// untuk mendapatkan single data user berdasarkan id
	GetUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	// untuk pencarian user berdasarkan email
	GetUserByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*User, error)
	// untuk pencarian user berdasarkan phone
	GetUserByPhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*User, error)
	// untuk membuat user baru
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	// untuk melakukan update user
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	// untuk delete user
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type userServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServicesClient(cc grpc.ClientConnInterface) UserServicesClient {
	return &userServicesClient{cc}
}

func (c *userServicesClient) GetUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserServices/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) GetUserByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserServices/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) GetUserByPhone(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserServices/GetUserByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserServices/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserServices/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserServices/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServicesServer is the server API for UserServices service.
// All implementations must embed UnimplementedUserServicesServer
// for forward compatibility
type UserServicesServer interface {
	// untuk mendapatkan single data user berdasarkan id
	GetUserByID(context.Context, *User) (*User, error)
	// untuk pencarian user berdasarkan email
	GetUserByEmail(context.Context, *Email) (*User, error)
	// untuk pencarian user berdasarkan phone
	GetUserByPhone(context.Context, *Phone) (*User, error)
	// untuk membuat user baru
	CreateUser(context.Context, *User) (*User, error)
	// untuk melakukan update user
	UpdateUser(context.Context, *User) (*User, error)
	// untuk delete user
	DeleteUser(context.Context, *User) (*User, error)
	mustEmbedUnimplementedUserServicesServer()
}

// UnimplementedUserServicesServer must be embedded to have forward compatible implementations.
type UnimplementedUserServicesServer struct {
}

func (UnimplementedUserServicesServer) GetUserByID(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedUserServicesServer) GetUserByEmail(context.Context, *Email) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserServicesServer) GetUserByPhone(context.Context, *Phone) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByPhone not implemented")
}
func (UnimplementedUserServicesServer) CreateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServicesServer) UpdateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServicesServer) DeleteUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServicesServer) mustEmbedUnimplementedUserServicesServer() {}

// UnsafeUserServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServicesServer will
// result in compilation errors.
type UnsafeUserServicesServer interface {
	mustEmbedUnimplementedUserServicesServer()
}

func RegisterUserServicesServer(s grpc.ServiceRegistrar, srv UserServicesServer) {
	s.RegisterService(&UserServices_ServiceDesc, srv)
}

func _UserServices_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserServices/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).GetUserByID(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserServices/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).GetUserByEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_GetUserByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).GetUserByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserServices/GetUserByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).GetUserByPhone(ctx, req.(*Phone))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserServices/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserServices/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServices_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicesServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserServices/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicesServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServices_ServiceDesc is the grpc.ServiceDesc for UserServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserServices",
	HandlerType: (*UserServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByID",
			Handler:    _UserServices_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserServices_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetUserByPhone",
			Handler:    _UserServices_GetUserByPhone_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserServices_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserServices_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserServices_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users.proto",
}
