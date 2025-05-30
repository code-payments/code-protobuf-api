// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: micropayment/v1/micro_payment_service.proto

package micropayment

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

// MicroPaymentClient is the client API for MicroPayment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroPaymentClient interface {
	// GetStatus gets basic request status
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	// RegisterWebhook registers a webhook for a request
	//
	// todo: Once Kik codes can encode the entire payment request details, we can
	//
	//	remove the messaging service component and have a Create RPC that
	//	reserves the intent ID with payment details, plus registers the webhook
	//	at the same time. Until that's possible, we're stuck with two RPC calls.
	RegisterWebhook(ctx context.Context, in *RegisterWebhookRequest, opts ...grpc.CallOption) (*RegisterWebhookResponse, error)
}

type microPaymentClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroPaymentClient(cc grpc.ClientConnInterface) MicroPaymentClient {
	return &microPaymentClient{cc}
}

func (c *microPaymentClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/code.micropayment.v1.MicroPayment/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microPaymentClient) RegisterWebhook(ctx context.Context, in *RegisterWebhookRequest, opts ...grpc.CallOption) (*RegisterWebhookResponse, error) {
	out := new(RegisterWebhookResponse)
	err := c.cc.Invoke(ctx, "/code.micropayment.v1.MicroPayment/RegisterWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroPaymentServer is the server API for MicroPayment service.
// All implementations must embed UnimplementedMicroPaymentServer
// for forward compatibility
type MicroPaymentServer interface {
	// GetStatus gets basic request status
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	// RegisterWebhook registers a webhook for a request
	//
	// todo: Once Kik codes can encode the entire payment request details, we can
	//
	//	remove the messaging service component and have a Create RPC that
	//	reserves the intent ID with payment details, plus registers the webhook
	//	at the same time. Until that's possible, we're stuck with two RPC calls.
	RegisterWebhook(context.Context, *RegisterWebhookRequest) (*RegisterWebhookResponse, error)
	mustEmbedUnimplementedMicroPaymentServer()
}

// UnimplementedMicroPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedMicroPaymentServer struct {
}

func (UnimplementedMicroPaymentServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedMicroPaymentServer) RegisterWebhook(context.Context, *RegisterWebhookRequest) (*RegisterWebhookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWebhook not implemented")
}
func (UnimplementedMicroPaymentServer) mustEmbedUnimplementedMicroPaymentServer() {}

// UnsafeMicroPaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroPaymentServer will
// result in compilation errors.
type UnsafeMicroPaymentServer interface {
	mustEmbedUnimplementedMicroPaymentServer()
}

func RegisterMicroPaymentServer(s grpc.ServiceRegistrar, srv MicroPaymentServer) {
	s.RegisterService(&MicroPayment_ServiceDesc, srv)
}

func _MicroPayment_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroPaymentServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.micropayment.v1.MicroPayment/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroPaymentServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroPayment_RegisterWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroPaymentServer).RegisterWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.micropayment.v1.MicroPayment/RegisterWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroPaymentServer).RegisterWebhook(ctx, req.(*RegisterWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MicroPayment_ServiceDesc is the grpc.ServiceDesc for MicroPayment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroPayment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.micropayment.v1.MicroPayment",
	HandlerType: (*MicroPaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _MicroPayment_GetStatus_Handler,
		},
		{
			MethodName: "RegisterWebhook",
			Handler:    _MicroPayment_RegisterWebhook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "micropayment/v1/micro_payment_service.proto",
}
