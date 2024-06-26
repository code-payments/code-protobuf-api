// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: chat/v1/chat_service.proto

package chat

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	// GetChats gets the set of chats for an owner account
	GetChats(ctx context.Context, in *GetChatsRequest, opts ...grpc.CallOption) (*GetChatsResponse, error)
	// GetMessages gets the set of messages for a chat
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	// AdvancePointer advances a pointer in chat history
	AdvancePointer(ctx context.Context, in *AdvancePointerRequest, opts ...grpc.CallOption) (*AdvancePointerResponse, error)
	// SetMuteState configures the mute state of a chat
	SetMuteState(ctx context.Context, in *SetMuteStateRequest, opts ...grpc.CallOption) (*SetMuteStateResponse, error)
	// SetSubscriptionState configures the susbscription state of a chat
	SetSubscriptionState(ctx context.Context, in *SetSubscriptionStateRequest, opts ...grpc.CallOption) (*SetSubscriptionStateResponse, error)
	StreamChatEvents(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamChatEventsClient, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) GetChats(ctx context.Context, in *GetChatsRequest, opts ...grpc.CallOption) (*GetChatsResponse, error) {
	out := new(GetChatsResponse)
	err := c.cc.Invoke(ctx, "/code.chat.v1.Chat/GetChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, "/code.chat.v1.Chat/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) AdvancePointer(ctx context.Context, in *AdvancePointerRequest, opts ...grpc.CallOption) (*AdvancePointerResponse, error) {
	out := new(AdvancePointerResponse)
	err := c.cc.Invoke(ctx, "/code.chat.v1.Chat/AdvancePointer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SetMuteState(ctx context.Context, in *SetMuteStateRequest, opts ...grpc.CallOption) (*SetMuteStateResponse, error) {
	out := new(SetMuteStateResponse)
	err := c.cc.Invoke(ctx, "/code.chat.v1.Chat/SetMuteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SetSubscriptionState(ctx context.Context, in *SetSubscriptionStateRequest, opts ...grpc.CallOption) (*SetSubscriptionStateResponse, error) {
	out := new(SetSubscriptionStateResponse)
	err := c.cc.Invoke(ctx, "/code.chat.v1.Chat/SetSubscriptionState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) StreamChatEvents(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamChatEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/code.chat.v1.Chat/StreamChatEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamChatEventsClient{stream}
	return x, nil
}

type Chat_StreamChatEventsClient interface {
	Send(*StreamChatEventsRequest) error
	Recv() (*StreamChatEventsResponse, error)
	grpc.ClientStream
}

type chatStreamChatEventsClient struct {
	grpc.ClientStream
}

func (x *chatStreamChatEventsClient) Send(m *StreamChatEventsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatStreamChatEventsClient) Recv() (*StreamChatEventsResponse, error) {
	m := new(StreamChatEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/code.chat.v1.Chat/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	// GetChats gets the set of chats for an owner account
	GetChats(context.Context, *GetChatsRequest) (*GetChatsResponse, error)
	// GetMessages gets the set of messages for a chat
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	// AdvancePointer advances a pointer in chat history
	AdvancePointer(context.Context, *AdvancePointerRequest) (*AdvancePointerResponse, error)
	// SetMuteState configures the mute state of a chat
	SetMuteState(context.Context, *SetMuteStateRequest) (*SetMuteStateResponse, error)
	// SetSubscriptionState configures the susbscription state of a chat
	SetSubscriptionState(context.Context, *SetSubscriptionStateRequest) (*SetSubscriptionStateResponse, error)
	StreamChatEvents(Chat_StreamChatEventsServer) error
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) GetChats(context.Context, *GetChatsRequest) (*GetChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedChatServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatServer) AdvancePointer(context.Context, *AdvancePointerRequest) (*AdvancePointerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdvancePointer not implemented")
}
func (UnimplementedChatServer) SetMuteState(context.Context, *SetMuteStateRequest) (*SetMuteStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMuteState not implemented")
}
func (UnimplementedChatServer) SetSubscriptionState(context.Context, *SetSubscriptionStateRequest) (*SetSubscriptionStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSubscriptionState not implemented")
}
func (UnimplementedChatServer) StreamChatEvents(Chat_StreamChatEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamChatEvents not implemented")
}
func (UnimplementedChatServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.chat.v1.Chat/GetChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetChats(ctx, req.(*GetChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.chat.v1.Chat/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_AdvancePointer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvancePointerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AdvancePointer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.chat.v1.Chat/AdvancePointer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AdvancePointer(ctx, req.(*AdvancePointerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SetMuteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMuteStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SetMuteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.chat.v1.Chat/SetMuteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SetMuteState(ctx, req.(*SetMuteStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SetSubscriptionState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSubscriptionStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SetSubscriptionState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.chat.v1.Chat/SetSubscriptionState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SetSubscriptionState(ctx, req.(*SetSubscriptionStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_StreamChatEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).StreamChatEvents(&chatStreamChatEventsServer{stream})
}

type Chat_StreamChatEventsServer interface {
	Send(*StreamChatEventsResponse) error
	Recv() (*StreamChatEventsRequest, error)
	grpc.ServerStream
}

type chatStreamChatEventsServer struct {
	grpc.ServerStream
}

func (x *chatStreamChatEventsServer) Send(m *StreamChatEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatStreamChatEventsServer) Recv() (*StreamChatEventsRequest, error) {
	m := new(StreamChatEventsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.chat.v1.Chat/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.chat.v1.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChats",
			Handler:    _Chat_GetChats_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Chat_GetMessages_Handler,
		},
		{
			MethodName: "AdvancePointer",
			Handler:    _Chat_AdvancePointer_Handler,
		},
		{
			MethodName: "SetMuteState",
			Handler:    _Chat_SetMuteState_Handler,
		},
		{
			MethodName: "SetSubscriptionState",
			Handler:    _Chat_SetSubscriptionState_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Chat_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChatEvents",
			Handler:       _Chat_StreamChatEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat/v1/chat_service.proto",
}
