// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user/v1/identity_service.proto

package user

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

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityClient interface {
	// LinkAccount links an owner account to the user identified and authenticated
	// by a one-time use token.
	//
	// Notably, this RPC has the following side effects:
	//   - A new user is automatically created if one doesn't exist.
	//   - Server will create a new data container for at least every unique
	//     owner account linked to the user.
	LinkAccount(ctx context.Context, in *LinkAccountRequest, opts ...grpc.CallOption) (*LinkAccountResponse, error)
	// UnlinkAccount removes links from an owner account. It will NOT remove
	// existing associations between users, owner accounts and identifying
	// features.
	//
	// The following associations will remain intact to ensure owner accounts
	// can continue to be used with a consistent login experience:
	//   - the user continues to be associated to existing owner accounts and
	//     identifying features
	//
	// Client can continue mainting their current login session. Their current
	// user and data container will remain the same.
	//
	// The call is guaranteed to be idempotent. It will not fail if the link is
	// already removed by either a previous call to this RPC or by a more recent
	// call to LinkAccount. A failure will only occur if the link between a user
	// and the owner accout or identifying feature never existed.
	UnlinkAccount(ctx context.Context, in *UnlinkAccountRequest, opts ...grpc.CallOption) (*UnlinkAccountResponse, error)
	// GetUser gets user information given a user identifier and an owner account.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type identityClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityClient(cc grpc.ClientConnInterface) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) LinkAccount(ctx context.Context, in *LinkAccountRequest, opts ...grpc.CallOption) (*LinkAccountResponse, error) {
	out := new(LinkAccountResponse)
	err := c.cc.Invoke(ctx, "/code.user.v1.Identity/LinkAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) UnlinkAccount(ctx context.Context, in *UnlinkAccountRequest, opts ...grpc.CallOption) (*UnlinkAccountResponse, error) {
	out := new(UnlinkAccountResponse)
	err := c.cc.Invoke(ctx, "/code.user.v1.Identity/UnlinkAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/code.user.v1.Identity/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
// All implementations must embed UnimplementedIdentityServer
// for forward compatibility
type IdentityServer interface {
	// LinkAccount links an owner account to the user identified and authenticated
	// by a one-time use token.
	//
	// Notably, this RPC has the following side effects:
	//   - A new user is automatically created if one doesn't exist.
	//   - Server will create a new data container for at least every unique
	//     owner account linked to the user.
	LinkAccount(context.Context, *LinkAccountRequest) (*LinkAccountResponse, error)
	// UnlinkAccount removes links from an owner account. It will NOT remove
	// existing associations between users, owner accounts and identifying
	// features.
	//
	// The following associations will remain intact to ensure owner accounts
	// can continue to be used with a consistent login experience:
	//   - the user continues to be associated to existing owner accounts and
	//     identifying features
	//
	// Client can continue mainting their current login session. Their current
	// user and data container will remain the same.
	//
	// The call is guaranteed to be idempotent. It will not fail if the link is
	// already removed by either a previous call to this RPC or by a more recent
	// call to LinkAccount. A failure will only occur if the link between a user
	// and the owner accout or identifying feature never existed.
	UnlinkAccount(context.Context, *UnlinkAccountRequest) (*UnlinkAccountResponse, error)
	// GetUser gets user information given a user identifier and an owner account.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedIdentityServer()
}

// UnimplementedIdentityServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (UnimplementedIdentityServer) LinkAccount(context.Context, *LinkAccountRequest) (*LinkAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkAccount not implemented")
}
func (UnimplementedIdentityServer) UnlinkAccount(context.Context, *UnlinkAccountRequest) (*UnlinkAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkAccount not implemented")
}
func (UnimplementedIdentityServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIdentityServer) mustEmbedUnimplementedIdentityServer() {}

// UnsafeIdentityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServer will
// result in compilation errors.
type UnsafeIdentityServer interface {
	mustEmbedUnimplementedIdentityServer()
}

func RegisterIdentityServer(s grpc.ServiceRegistrar, srv IdentityServer) {
	s.RegisterService(&Identity_ServiceDesc, srv)
}

func _Identity_LinkAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).LinkAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.user.v1.Identity/LinkAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).LinkAccount(ctx, req.(*LinkAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_UnlinkAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).UnlinkAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.user.v1.Identity/UnlinkAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).UnlinkAccount(ctx, req.(*UnlinkAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.user.v1.Identity/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Identity_ServiceDesc is the grpc.ServiceDesc for Identity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Identity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.user.v1.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LinkAccount",
			Handler:    _Identity_LinkAccount_Handler,
		},
		{
			MethodName: "UnlinkAccount",
			Handler:    _Identity_UnlinkAccount_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Identity_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/identity_service.proto",
}
