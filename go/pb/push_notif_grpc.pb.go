// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// FcmPersonalClient is the client API for FcmPersonal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FcmPersonalClient interface {
	Personal(ctx context.Context, in *DataFcmPersonal, opts ...grpc.CallOption) (*FcmResponse, error)
	Multiple(ctx context.Context, in *DataFcmMultiple, opts ...grpc.CallOption) (*FcmResponse, error)
}

type fcmPersonalClient struct {
	cc grpc.ClientConnInterface
}

func NewFcmPersonalClient(cc grpc.ClientConnInterface) FcmPersonalClient {
	return &fcmPersonalClient{cc}
}

func (c *fcmPersonalClient) Personal(ctx context.Context, in *DataFcmPersonal, opts ...grpc.CallOption) (*FcmResponse, error) {
	out := new(FcmResponse)
	err := c.cc.Invoke(ctx, "/push.notification.FcmPersonal/Personal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fcmPersonalClient) Multiple(ctx context.Context, in *DataFcmMultiple, opts ...grpc.CallOption) (*FcmResponse, error) {
	out := new(FcmResponse)
	err := c.cc.Invoke(ctx, "/push.notification.FcmPersonal/Multiple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FcmPersonalServer is the server API for FcmPersonal service.
// All implementations must embed UnimplementedFcmPersonalServer
// for forward compatibility
type FcmPersonalServer interface {
	Personal(context.Context, *DataFcmPersonal) (*FcmResponse, error)
	Multiple(context.Context, *DataFcmMultiple) (*FcmResponse, error)
	mustEmbedUnimplementedFcmPersonalServer()
}

// UnimplementedFcmPersonalServer must be embedded to have forward compatible implementations.
type UnimplementedFcmPersonalServer struct {
}

func (UnimplementedFcmPersonalServer) Personal(context.Context, *DataFcmPersonal) (*FcmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Personal not implemented")
}
func (UnimplementedFcmPersonalServer) Multiple(context.Context, *DataFcmMultiple) (*FcmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiple not implemented")
}
func (UnimplementedFcmPersonalServer) mustEmbedUnimplementedFcmPersonalServer() {}

// UnsafeFcmPersonalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FcmPersonalServer will
// result in compilation errors.
type UnsafeFcmPersonalServer interface {
	mustEmbedUnimplementedFcmPersonalServer()
}

func RegisterFcmPersonalServer(s grpc.ServiceRegistrar, srv FcmPersonalServer) {
	s.RegisterService(&FcmPersonal_ServiceDesc, srv)
}

func _FcmPersonal_Personal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFcmPersonal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FcmPersonalServer).Personal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.notification.FcmPersonal/Personal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FcmPersonalServer).Personal(ctx, req.(*DataFcmPersonal))
	}
	return interceptor(ctx, in, info, handler)
}

func _FcmPersonal_Multiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFcmMultiple)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FcmPersonalServer).Multiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.notification.FcmPersonal/Multiple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FcmPersonalServer).Multiple(ctx, req.(*DataFcmMultiple))
	}
	return interceptor(ctx, in, info, handler)
}

// FcmPersonal_ServiceDesc is the grpc.ServiceDesc for FcmPersonal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FcmPersonal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "push.notification.FcmPersonal",
	HandlerType: (*FcmPersonalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Personal",
			Handler:    _FcmPersonal_Personal_Handler,
		},
		{
			MethodName: "Multiple",
			Handler:    _FcmPersonal_Multiple_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/push_notif.proto",
}
