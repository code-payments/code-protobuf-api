// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: messaging/v1/messaging_service.proto

package messaging

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
	Messaging_OpenMessageStream_FullMethodName              = "/code.messaging.v1.Messaging/OpenMessageStream"
	Messaging_OpenMessageStreamWithKeepAlive_FullMethodName = "/code.messaging.v1.Messaging/OpenMessageStreamWithKeepAlive"
	Messaging_PollMessages_FullMethodName                   = "/code.messaging.v1.Messaging/PollMessages"
	Messaging_AckMessages_FullMethodName                    = "/code.messaging.v1.Messaging/AckMessages"
	Messaging_SendMessage_FullMethodName                    = "/code.messaging.v1.Messaging/SendMessage"
)

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	// OpenMessageStream opens a stream of messages. Messages are routed using the
	// public key of a rendezvous keypair derived by both the sender and the
	// recipient of the messages. The sender may be a client or server.
	//
	// Messages are expected to be acked once they have been processed by the client.
	// Ack'd messages will no longer be delivered on future OpenMessageStream calls,
	// and are eligible for deletion from the service. Clients should, however, handle
	// duplicate delivery of messages.
	//
	// For grabbing a bill, the expected flow is as follows:
	//  1. The payment sender creates a cash scan code
	//  2. The payment sender calls OpenMessageStream on the rendezvous public key, which is
	//     derived by using sha256(scan payload) as the keypair seed.
	//  3. The payment recipient scans the code and uses SendMessage to send their account ID
	//     back to the sender via the rendezvous public key.
	//  4. The payment sender receives the message, submits the intent, and closes the stream.
	//
	// For receiving a bill of requested value, the expected flow is as follows:
	//  1. The payment recipient uses SendMessage to send their account ID and payment amount to
	//     the sender via the rendezvous public key, which is derived by using sha256(scan payload)
	//     as the keypair seed.
	//  2. The payment recipient calls OpenMessageStream on the rendezvous public key to listen
	//     for status messages generated by client/server. It must ignore the original message it sent
	//     as part of step 1.
	//  3. The payment recipient creates a payment request scan code
	//  4. The payment sender calls PollMessages on the rendezvous public key. This is ok because
	//     we know the message exists per step 1, and doesn't actually incur a long poll. This is a
	//     required hack because we don't have the infrastructure in place to allow multiple listens
	//     on the same stream, and the recipient needs real-time status updates.
	//  5. The payment sender receives the message (any status messages are ignored), and submits the
	//     intent.
	//  6. The payment recipient observes status message (eg. IntentSubmitted, ClientRejectedPayment,
	//     WebhookCalled) for payment state.
	//  7. The payment recipient closes the stream once the payment hits a terminal state, or times out.
	OpenMessageStream(ctx context.Context, in *OpenMessageStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OpenMessageStreamResponse], error)
	// OpenMessageStreamWithKeepAlive is like OpenMessageStream, but enables a ping/pong
	// keepalive to determine the health of the stream at both the client and server.
	//
	// The keepalive protocol is as follows:
	//  1. Client initiates a stream by sending an OpenMessageStreamRequest.
	//  2. Upon stream initialization, server begins the keepalive protocol.
	//  3. Server sends a ping to the client.
	//  4. Client responds with a pong as fast as possible, making note of
	//     the delay for when to expect the next ping.
	//  5. Steps 3 and 4 are repeated until the stream is explicitly terminated
	//     or is deemed to be unhealthy.
	//
	// Client notes:
	//   - Client should be careful to process messages async, so any responses to pings are
	//     not delayed.
	//   - Clients should implement a reasonable backoff strategy upon continued timeout failures.
	//   - Clients that abuse pong messages may have their streams terminated by server.
	//
	// At any point in the stream, server will respond with messages in real time as
	// they are observed. Messages sent over the stream should not affect the ping/pong
	// protocol timings. Individual protocols for payment flows remain the same, and are
	// documented in OpenMessageStream.
	//
	// Note: This API will enforce OpenMessageStreamRequest.signature is set as part of migration
	//
	//	to this newer protocol
	OpenMessageStreamWithKeepAlive(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse], error)
	// PollMessages is like OpenMessageStream, but uses a polling flow for receiving
	// messages. Updates are not real-time and depedent on the polling interval.
	// This RPC supports all message types.
	//
	// This is a temporary RPC until OpenMessageStream can be built out generically on
	// both client and server, while supporting things like multiple listeners.
	PollMessages(ctx context.Context, in *PollMessagesRequest, opts ...grpc.CallOption) (*PollMessagesResponse, error)
	// AckMessages acks one or more messages that have been successfully delivered to
	// the client.
	AckMessages(ctx context.Context, in *AckMessagesRequest, opts ...grpc.CallOption) (*AckMesssagesResponse, error)
	// SendMessage sends a message.
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) OpenMessageStream(ctx context.Context, in *OpenMessageStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OpenMessageStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[0], Messaging_OpenMessageStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OpenMessageStreamRequest, OpenMessageStreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messaging_OpenMessageStreamClient = grpc.ServerStreamingClient[OpenMessageStreamResponse]

func (c *messagingClient) OpenMessageStreamWithKeepAlive(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[1], Messaging_OpenMessageStreamWithKeepAlive_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messaging_OpenMessageStreamWithKeepAliveClient = grpc.BidiStreamingClient[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse]

func (c *messagingClient) PollMessages(ctx context.Context, in *PollMessagesRequest, opts ...grpc.CallOption) (*PollMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PollMessagesResponse)
	err := c.cc.Invoke(ctx, Messaging_PollMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) AckMessages(ctx context.Context, in *AckMessagesRequest, opts ...grpc.CallOption) (*AckMesssagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AckMesssagesResponse)
	err := c.cc.Invoke(ctx, Messaging_AckMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, Messaging_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility.
type MessagingServer interface {
	// OpenMessageStream opens a stream of messages. Messages are routed using the
	// public key of a rendezvous keypair derived by both the sender and the
	// recipient of the messages. The sender may be a client or server.
	//
	// Messages are expected to be acked once they have been processed by the client.
	// Ack'd messages will no longer be delivered on future OpenMessageStream calls,
	// and are eligible for deletion from the service. Clients should, however, handle
	// duplicate delivery of messages.
	//
	// For grabbing a bill, the expected flow is as follows:
	//  1. The payment sender creates a cash scan code
	//  2. The payment sender calls OpenMessageStream on the rendezvous public key, which is
	//     derived by using sha256(scan payload) as the keypair seed.
	//  3. The payment recipient scans the code and uses SendMessage to send their account ID
	//     back to the sender via the rendezvous public key.
	//  4. The payment sender receives the message, submits the intent, and closes the stream.
	//
	// For receiving a bill of requested value, the expected flow is as follows:
	//  1. The payment recipient uses SendMessage to send their account ID and payment amount to
	//     the sender via the rendezvous public key, which is derived by using sha256(scan payload)
	//     as the keypair seed.
	//  2. The payment recipient calls OpenMessageStream on the rendezvous public key to listen
	//     for status messages generated by client/server. It must ignore the original message it sent
	//     as part of step 1.
	//  3. The payment recipient creates a payment request scan code
	//  4. The payment sender calls PollMessages on the rendezvous public key. This is ok because
	//     we know the message exists per step 1, and doesn't actually incur a long poll. This is a
	//     required hack because we don't have the infrastructure in place to allow multiple listens
	//     on the same stream, and the recipient needs real-time status updates.
	//  5. The payment sender receives the message (any status messages are ignored), and submits the
	//     intent.
	//  6. The payment recipient observes status message (eg. IntentSubmitted, ClientRejectedPayment,
	//     WebhookCalled) for payment state.
	//  7. The payment recipient closes the stream once the payment hits a terminal state, or times out.
	OpenMessageStream(*OpenMessageStreamRequest, grpc.ServerStreamingServer[OpenMessageStreamResponse]) error
	// OpenMessageStreamWithKeepAlive is like OpenMessageStream, but enables a ping/pong
	// keepalive to determine the health of the stream at both the client and server.
	//
	// The keepalive protocol is as follows:
	//  1. Client initiates a stream by sending an OpenMessageStreamRequest.
	//  2. Upon stream initialization, server begins the keepalive protocol.
	//  3. Server sends a ping to the client.
	//  4. Client responds with a pong as fast as possible, making note of
	//     the delay for when to expect the next ping.
	//  5. Steps 3 and 4 are repeated until the stream is explicitly terminated
	//     or is deemed to be unhealthy.
	//
	// Client notes:
	//   - Client should be careful to process messages async, so any responses to pings are
	//     not delayed.
	//   - Clients should implement a reasonable backoff strategy upon continued timeout failures.
	//   - Clients that abuse pong messages may have their streams terminated by server.
	//
	// At any point in the stream, server will respond with messages in real time as
	// they are observed. Messages sent over the stream should not affect the ping/pong
	// protocol timings. Individual protocols for payment flows remain the same, and are
	// documented in OpenMessageStream.
	//
	// Note: This API will enforce OpenMessageStreamRequest.signature is set as part of migration
	//
	//	to this newer protocol
	OpenMessageStreamWithKeepAlive(grpc.BidiStreamingServer[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse]) error
	// PollMessages is like OpenMessageStream, but uses a polling flow for receiving
	// messages. Updates are not real-time and depedent on the polling interval.
	// This RPC supports all message types.
	//
	// This is a temporary RPC until OpenMessageStream can be built out generically on
	// both client and server, while supporting things like multiple listeners.
	PollMessages(context.Context, *PollMessagesRequest) (*PollMessagesResponse, error)
	// AckMessages acks one or more messages that have been successfully delivered to
	// the client.
	AckMessages(context.Context, *AckMessagesRequest) (*AckMesssagesResponse, error)
	// SendMessage sends a message.
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessagingServer struct{}

func (UnimplementedMessagingServer) OpenMessageStream(*OpenMessageStreamRequest, grpc.ServerStreamingServer[OpenMessageStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method OpenMessageStream not implemented")
}
func (UnimplementedMessagingServer) OpenMessageStreamWithKeepAlive(grpc.BidiStreamingServer[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse]) error {
	return status.Errorf(codes.Unimplemented, "method OpenMessageStreamWithKeepAlive not implemented")
}
func (UnimplementedMessagingServer) PollMessages(context.Context, *PollMessagesRequest) (*PollMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PollMessages not implemented")
}
func (UnimplementedMessagingServer) AckMessages(context.Context, *AckMessagesRequest) (*AckMesssagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckMessages not implemented")
}
func (UnimplementedMessagingServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagingServer) mustEmbedUnimplementedMessagingServer() {}
func (UnimplementedMessagingServer) testEmbeddedByValue()                   {}

// UnsafeMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServer will
// result in compilation errors.
type UnsafeMessagingServer interface {
	mustEmbedUnimplementedMessagingServer()
}

func RegisterMessagingServer(s grpc.ServiceRegistrar, srv MessagingServer) {
	// If the following call pancis, it indicates UnimplementedMessagingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Messaging_ServiceDesc, srv)
}

func _Messaging_OpenMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OpenMessageStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServer).OpenMessageStream(m, &grpc.GenericServerStream[OpenMessageStreamRequest, OpenMessageStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messaging_OpenMessageStreamServer = grpc.ServerStreamingServer[OpenMessageStreamResponse]

func _Messaging_OpenMessageStreamWithKeepAlive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServer).OpenMessageStreamWithKeepAlive(&grpc.GenericServerStream[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messaging_OpenMessageStreamWithKeepAliveServer = grpc.BidiStreamingServer[OpenMessageStreamWithKeepAliveRequest, OpenMessageStreamWithKeepAliveResponse]

func _Messaging_PollMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).PollMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_PollMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).PollMessages(ctx, req.(*PollMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_AckMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).AckMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_AckMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).AckMessages(ctx, req.(*AckMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.messaging.v1.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PollMessages",
			Handler:    _Messaging_PollMessages_Handler,
		},
		{
			MethodName: "AckMessages",
			Handler:    _Messaging_AckMessages_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Messaging_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenMessageStream",
			Handler:       _Messaging_OpenMessageStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OpenMessageStreamWithKeepAlive",
			Handler:       _Messaging_OpenMessageStreamWithKeepAlive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messaging/v1/messaging_service.proto",
}
