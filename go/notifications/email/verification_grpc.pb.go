// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package email

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

// VerificationServiceClient is the client API for VerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerificationServiceClient interface {
	// untuk mengirim email verifikasi
	SendEmailVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error)
}

type verificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificationServiceClient(cc grpc.ClientConnInterface) VerificationServiceClient {
	return &verificationServiceClient{cc}
}

func (c *verificationServiceClient) SendEmailVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error) {
	out := new(VerificationResponse)
	err := c.cc.Invoke(ctx, "/VerificationService/SendEmailVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationServiceServer is the server API for VerificationService service.
// All implementations must embed UnimplementedVerificationServiceServer
// for forward compatibility
type VerificationServiceServer interface {
	// untuk mengirim email verifikasi
	SendEmailVerification(context.Context, *VerificationRequest) (*VerificationResponse, error)
	mustEmbedUnimplementedVerificationServiceServer()
}

// UnimplementedVerificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVerificationServiceServer struct {
}

func (UnimplementedVerificationServiceServer) SendEmailVerification(context.Context, *VerificationRequest) (*VerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailVerification not implemented")
}
func (UnimplementedVerificationServiceServer) mustEmbedUnimplementedVerificationServiceServer() {}

// UnsafeVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerificationServiceServer will
// result in compilation errors.
type UnsafeVerificationServiceServer interface {
	mustEmbedUnimplementedVerificationServiceServer()
}

func RegisterVerificationServiceServer(s grpc.ServiceRegistrar, srv VerificationServiceServer) {
	s.RegisterService(&VerificationService_ServiceDesc, srv)
}

func _VerificationService_SendEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).SendEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VerificationService/SendEmailVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).SendEmailVerification(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerificationService_ServiceDesc is the grpc.ServiceDesc for VerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VerificationService",
	HandlerType: (*VerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmailVerification",
			Handler:    _VerificationService_SendEmailVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/verification.proto",
}
