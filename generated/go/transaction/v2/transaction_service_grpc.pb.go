// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: transaction/v2/transaction_service.proto

package transaction

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

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	// SubmitIntent is the mechanism for client and server to agree upon a set of
	// client actions to execute on the blockchain using the Code sequencer for
	// fulfillment.
	//
	// Transactions and virtual instructions are never exchanged between client and server.
	// Instead, the required accounts and arguments for instructions known to each actor are
	// exchanged to allow independent and local construction.
	//
	// Client and server are expected to fully validate the intent. Proofs will
	// be provided for any parameter requiring one. Signatures should only be
	// generated after approval.
	//
	// This RPC is not a traditional streaming endpoint. It bundles two unary calls
	// to enable DB-level transaction semantics.
	//
	// The high-level happy path flow for the RPC is as follows:
	//  1. Client initiates a stream and sends SubmitIntentRequest.SubmitActions
	//  2. Server validates the intent, its actions and metadata
	//     3a. If there are transactions or virtual instructions requiring the user's signature,
	//     then server returns SubmitIntentResponse.ServerParameters
	//     3b. Otherwise, server returns SubmitIntentResponse.Success and closes the
	//     stream
	//  4. For each transaction or virtual instruction requiring the user's signature, the client
	//     locally constructs it, performs validation and collects the signature
	//  5. Client sends SubmitIntentRequest.SubmitSignatures with the signature
	//     list generated from 4
	//  6. Server validates all signatures are submitted and are the expected values
	//     using locally constructed transactions or virtual instructions.
	//  7. Server returns SubmitIntentResponse.Success and closes the stream
	//
	// In the error case:
	//   - Server will return SubmitIntentResponse.Error and close the stream
	//   - Client will close the stream
	SubmitIntent(ctx context.Context, opts ...grpc.CallOption) (Transaction_SubmitIntentClient, error)
	// GetIntentMetadata gets basic metadata on an intent. It can also be used
	// to fetch the status of submitted intents. Metadata exists only for intents
	// that have been successfully submitted.
	GetIntentMetadata(ctx context.Context, in *GetIntentMetadataRequest, opts ...grpc.CallOption) (*GetIntentMetadataResponse, error)
	// GetLimits gets limits for money moving intents for an owner account in an
	// identity-aware manner
	GetLimits(ctx context.Context, in *GetLimitsRequest, opts ...grpc.CallOption) (*GetLimitsResponse, error)
	// CanWithdrawToAccount provides hints to clients for submitting withdraw intents.
	// The RPC indicates if a withdrawal is possible, and how it should be performed.
	CanWithdrawToAccount(ctx context.Context, in *CanWithdrawToAccountRequest, opts ...grpc.CallOption) (*CanWithdrawToAccountResponse, error)
	// Airdrop airdrops core mint tokens to the requesting account
	Airdrop(ctx context.Context, in *AirdropRequest, opts ...grpc.CallOption) (*AirdropResponse, error)
	// Swap performs an on-chain swap. The high-level flow mirrors SubmitIntent
	// closely. However, due to the time-sensitive nature and unreliability of
	// swaps, they do not fit within the broader intent system. This results in
	// a few key differences:
	//   - Transactions are submitted on a best-effort basis outside of the Code
	//     Sequencer within the RPC handler
	//   - Balance changes are applied after the transaction has finalized
	//   - Transactions use recent blockhashes over a nonce
	//
	// SubmitIntent also operates on VM virtual instructions, whereas Swap uses
	// Solana transactions.
	//
	// The transaction will have the following instruction format:
	//  1. ComputeBudget::SetComputeUnitLimit
	//  2. ComputeBudget::SetComputeUnitPrice
	//  3. SwapValidator::PreSwap
	//  4. Dynamic swap instruction
	//  5. SwapValidator::PostSwap
	//
	// Note: Currently limited to swapping USDC to core mint tokens.
	// Note: Core mint tokens are deposited into the token account derived from the VM deposit PDA of the owner account.
	Swap(ctx context.Context, opts ...grpc.CallOption) (Transaction_SwapClient, error)
	// DeclareFiatOnrampPurchaseAttempt is called whenever a user attempts to use a fiat
	// onramp to purchase core mint tokens for use in Code.
	DeclareFiatOnrampPurchaseAttempt(ctx context.Context, in *DeclareFiatOnrampPurchaseAttemptRequest, opts ...grpc.CallOption) (*DeclareFiatOnrampPurchaseAttemptResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) SubmitIntent(ctx context.Context, opts ...grpc.CallOption) (Transaction_SubmitIntentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transaction_ServiceDesc.Streams[0], "/code.transaction.v2.Transaction/SubmitIntent", opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionSubmitIntentClient{stream}
	return x, nil
}

type Transaction_SubmitIntentClient interface {
	Send(*SubmitIntentRequest) error
	Recv() (*SubmitIntentResponse, error)
	grpc.ClientStream
}

type transactionSubmitIntentClient struct {
	grpc.ClientStream
}

func (x *transactionSubmitIntentClient) Send(m *SubmitIntentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transactionSubmitIntentClient) Recv() (*SubmitIntentResponse, error) {
	m := new(SubmitIntentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transactionClient) GetIntentMetadata(ctx context.Context, in *GetIntentMetadataRequest, opts ...grpc.CallOption) (*GetIntentMetadataResponse, error) {
	out := new(GetIntentMetadataResponse)
	err := c.cc.Invoke(ctx, "/code.transaction.v2.Transaction/GetIntentMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) GetLimits(ctx context.Context, in *GetLimitsRequest, opts ...grpc.CallOption) (*GetLimitsResponse, error) {
	out := new(GetLimitsResponse)
	err := c.cc.Invoke(ctx, "/code.transaction.v2.Transaction/GetLimits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) CanWithdrawToAccount(ctx context.Context, in *CanWithdrawToAccountRequest, opts ...grpc.CallOption) (*CanWithdrawToAccountResponse, error) {
	out := new(CanWithdrawToAccountResponse)
	err := c.cc.Invoke(ctx, "/code.transaction.v2.Transaction/CanWithdrawToAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) Airdrop(ctx context.Context, in *AirdropRequest, opts ...grpc.CallOption) (*AirdropResponse, error) {
	out := new(AirdropResponse)
	err := c.cc.Invoke(ctx, "/code.transaction.v2.Transaction/Airdrop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) Swap(ctx context.Context, opts ...grpc.CallOption) (Transaction_SwapClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transaction_ServiceDesc.Streams[1], "/code.transaction.v2.Transaction/Swap", opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionSwapClient{stream}
	return x, nil
}

type Transaction_SwapClient interface {
	Send(*SwapRequest) error
	Recv() (*SwapResponse, error)
	grpc.ClientStream
}

type transactionSwapClient struct {
	grpc.ClientStream
}

func (x *transactionSwapClient) Send(m *SwapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transactionSwapClient) Recv() (*SwapResponse, error) {
	m := new(SwapResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transactionClient) DeclareFiatOnrampPurchaseAttempt(ctx context.Context, in *DeclareFiatOnrampPurchaseAttemptRequest, opts ...grpc.CallOption) (*DeclareFiatOnrampPurchaseAttemptResponse, error) {
	out := new(DeclareFiatOnrampPurchaseAttemptResponse)
	err := c.cc.Invoke(ctx, "/code.transaction.v2.Transaction/DeclareFiatOnrampPurchaseAttempt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	// SubmitIntent is the mechanism for client and server to agree upon a set of
	// client actions to execute on the blockchain using the Code sequencer for
	// fulfillment.
	//
	// Transactions and virtual instructions are never exchanged between client and server.
	// Instead, the required accounts and arguments for instructions known to each actor are
	// exchanged to allow independent and local construction.
	//
	// Client and server are expected to fully validate the intent. Proofs will
	// be provided for any parameter requiring one. Signatures should only be
	// generated after approval.
	//
	// This RPC is not a traditional streaming endpoint. It bundles two unary calls
	// to enable DB-level transaction semantics.
	//
	// The high-level happy path flow for the RPC is as follows:
	//  1. Client initiates a stream and sends SubmitIntentRequest.SubmitActions
	//  2. Server validates the intent, its actions and metadata
	//     3a. If there are transactions or virtual instructions requiring the user's signature,
	//     then server returns SubmitIntentResponse.ServerParameters
	//     3b. Otherwise, server returns SubmitIntentResponse.Success and closes the
	//     stream
	//  4. For each transaction or virtual instruction requiring the user's signature, the client
	//     locally constructs it, performs validation and collects the signature
	//  5. Client sends SubmitIntentRequest.SubmitSignatures with the signature
	//     list generated from 4
	//  6. Server validates all signatures are submitted and are the expected values
	//     using locally constructed transactions or virtual instructions.
	//  7. Server returns SubmitIntentResponse.Success and closes the stream
	//
	// In the error case:
	//   - Server will return SubmitIntentResponse.Error and close the stream
	//   - Client will close the stream
	SubmitIntent(Transaction_SubmitIntentServer) error
	// GetIntentMetadata gets basic metadata on an intent. It can also be used
	// to fetch the status of submitted intents. Metadata exists only for intents
	// that have been successfully submitted.
	GetIntentMetadata(context.Context, *GetIntentMetadataRequest) (*GetIntentMetadataResponse, error)
	// GetLimits gets limits for money moving intents for an owner account in an
	// identity-aware manner
	GetLimits(context.Context, *GetLimitsRequest) (*GetLimitsResponse, error)
	// CanWithdrawToAccount provides hints to clients for submitting withdraw intents.
	// The RPC indicates if a withdrawal is possible, and how it should be performed.
	CanWithdrawToAccount(context.Context, *CanWithdrawToAccountRequest) (*CanWithdrawToAccountResponse, error)
	// Airdrop airdrops core mint tokens to the requesting account
	Airdrop(context.Context, *AirdropRequest) (*AirdropResponse, error)
	// Swap performs an on-chain swap. The high-level flow mirrors SubmitIntent
	// closely. However, due to the time-sensitive nature and unreliability of
	// swaps, they do not fit within the broader intent system. This results in
	// a few key differences:
	//   - Transactions are submitted on a best-effort basis outside of the Code
	//     Sequencer within the RPC handler
	//   - Balance changes are applied after the transaction has finalized
	//   - Transactions use recent blockhashes over a nonce
	//
	// SubmitIntent also operates on VM virtual instructions, whereas Swap uses
	// Solana transactions.
	//
	// The transaction will have the following instruction format:
	//  1. ComputeBudget::SetComputeUnitLimit
	//  2. ComputeBudget::SetComputeUnitPrice
	//  3. SwapValidator::PreSwap
	//  4. Dynamic swap instruction
	//  5. SwapValidator::PostSwap
	//
	// Note: Currently limited to swapping USDC to core mint tokens.
	// Note: Core mint tokens are deposited into the token account derived from the VM deposit PDA of the owner account.
	Swap(Transaction_SwapServer) error
	// DeclareFiatOnrampPurchaseAttempt is called whenever a user attempts to use a fiat
	// onramp to purchase core mint tokens for use in Code.
	DeclareFiatOnrampPurchaseAttempt(context.Context, *DeclareFiatOnrampPurchaseAttemptRequest) (*DeclareFiatOnrampPurchaseAttemptResponse, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) SubmitIntent(Transaction_SubmitIntentServer) error {
	return status.Errorf(codes.Unimplemented, "method SubmitIntent not implemented")
}
func (UnimplementedTransactionServer) GetIntentMetadata(context.Context, *GetIntentMetadataRequest) (*GetIntentMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntentMetadata not implemented")
}
func (UnimplementedTransactionServer) GetLimits(context.Context, *GetLimitsRequest) (*GetLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimits not implemented")
}
func (UnimplementedTransactionServer) CanWithdrawToAccount(context.Context, *CanWithdrawToAccountRequest) (*CanWithdrawToAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanWithdrawToAccount not implemented")
}
func (UnimplementedTransactionServer) Airdrop(context.Context, *AirdropRequest) (*AirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Airdrop not implemented")
}
func (UnimplementedTransactionServer) Swap(Transaction_SwapServer) error {
	return status.Errorf(codes.Unimplemented, "method Swap not implemented")
}
func (UnimplementedTransactionServer) DeclareFiatOnrampPurchaseAttempt(context.Context, *DeclareFiatOnrampPurchaseAttemptRequest) (*DeclareFiatOnrampPurchaseAttemptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclareFiatOnrampPurchaseAttempt not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_SubmitIntent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransactionServer).SubmitIntent(&transactionSubmitIntentServer{stream})
}

type Transaction_SubmitIntentServer interface {
	Send(*SubmitIntentResponse) error
	Recv() (*SubmitIntentRequest, error)
	grpc.ServerStream
}

type transactionSubmitIntentServer struct {
	grpc.ServerStream
}

func (x *transactionSubmitIntentServer) Send(m *SubmitIntentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transactionSubmitIntentServer) Recv() (*SubmitIntentRequest, error) {
	m := new(SubmitIntentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Transaction_GetIntentMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntentMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetIntentMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.transaction.v2.Transaction/GetIntentMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetIntentMetadata(ctx, req.(*GetIntentMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_GetLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.transaction.v2.Transaction/GetLimits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetLimits(ctx, req.(*GetLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_CanWithdrawToAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanWithdrawToAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CanWithdrawToAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.transaction.v2.Transaction/CanWithdrawToAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CanWithdrawToAccount(ctx, req.(*CanWithdrawToAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_Airdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirdropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).Airdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.transaction.v2.Transaction/Airdrop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).Airdrop(ctx, req.(*AirdropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_Swap_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransactionServer).Swap(&transactionSwapServer{stream})
}

type Transaction_SwapServer interface {
	Send(*SwapResponse) error
	Recv() (*SwapRequest, error)
	grpc.ServerStream
}

type transactionSwapServer struct {
	grpc.ServerStream
}

func (x *transactionSwapServer) Send(m *SwapResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transactionSwapServer) Recv() (*SwapRequest, error) {
	m := new(SwapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Transaction_DeclareFiatOnrampPurchaseAttempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclareFiatOnrampPurchaseAttemptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).DeclareFiatOnrampPurchaseAttempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.transaction.v2.Transaction/DeclareFiatOnrampPurchaseAttempt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).DeclareFiatOnrampPurchaseAttempt(ctx, req.(*DeclareFiatOnrampPurchaseAttemptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.transaction.v2.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIntentMetadata",
			Handler:    _Transaction_GetIntentMetadata_Handler,
		},
		{
			MethodName: "GetLimits",
			Handler:    _Transaction_GetLimits_Handler,
		},
		{
			MethodName: "CanWithdrawToAccount",
			Handler:    _Transaction_CanWithdrawToAccount_Handler,
		},
		{
			MethodName: "Airdrop",
			Handler:    _Transaction_Airdrop_Handler,
		},
		{
			MethodName: "DeclareFiatOnrampPurchaseAttempt",
			Handler:    _Transaction_DeclareFiatOnrampPurchaseAttempt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitIntent",
			Handler:       _Transaction_SubmitIntent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Swap",
			Handler:       _Transaction_Swap_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "transaction/v2/transaction_service.proto",
}
