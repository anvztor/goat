// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: goat/bitcoin/v1/query.proto

package bitcoinv1

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
	Query_Params_FullMethodName         = "/goat.bitcoin.v1.Query/Params"
	Query_Pubkey_FullMethodName         = "/goat.bitcoin.v1.Query/Pubkey"
	Query_DepositAddress_FullMethodName = "/goat.bitcoin.v1.Query/DepositAddress"
	Query_HasDeposited_FullMethodName   = "/goat.bitcoin.v1.Query/HasDeposited"
	Query_Withdrawal_FullMethodName     = "/goat.bitcoin.v1.Query/Withdrawal"
	Query_BlockTip_FullMethodName       = "/goat.bitcoin.v1.Query/BlockTip"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Pubkeys queries current public key for deposit
	Pubkey(ctx context.Context, in *QueryPubkeyRequest, opts ...grpc.CallOption) (*QueryPubkeyResponse, error)
	// DepositAddress queries current address for deposit
	DepositAddress(ctx context.Context, in *QueryDepositAddress, opts ...grpc.CallOption) (*QueryDepositAddressResponse, error)
	// HasDeposited checks if a deposit transaction is confirmed on chain
	HasDeposited(ctx context.Context, in *QueryHasDeposited, opts ...grpc.CallOption) (*QueryHasDepositedResponse, error)
	// Withdrawal queries all current public keys
	Withdrawal(ctx context.Context, in *QueryWithdrawalRequest, opts ...grpc.CallOption) (*QueryWithdrawalResponse, error)
	// BlockTip queries current the latest confirmed bitcoin height by relayer
	BlockTip(ctx context.Context, in *QueryBlockTipRequest, opts ...grpc.CallOption) (*QueryBlockTipResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pubkey(ctx context.Context, in *QueryPubkeyRequest, opts ...grpc.CallOption) (*QueryPubkeyResponse, error) {
	out := new(QueryPubkeyResponse)
	err := c.cc.Invoke(ctx, Query_Pubkey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DepositAddress(ctx context.Context, in *QueryDepositAddress, opts ...grpc.CallOption) (*QueryDepositAddressResponse, error) {
	out := new(QueryDepositAddressResponse)
	err := c.cc.Invoke(ctx, Query_DepositAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HasDeposited(ctx context.Context, in *QueryHasDeposited, opts ...grpc.CallOption) (*QueryHasDepositedResponse, error) {
	out := new(QueryHasDepositedResponse)
	err := c.cc.Invoke(ctx, Query_HasDeposited_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Withdrawal(ctx context.Context, in *QueryWithdrawalRequest, opts ...grpc.CallOption) (*QueryWithdrawalResponse, error) {
	out := new(QueryWithdrawalResponse)
	err := c.cc.Invoke(ctx, Query_Withdrawal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BlockTip(ctx context.Context, in *QueryBlockTipRequest, opts ...grpc.CallOption) (*QueryBlockTipResponse, error) {
	out := new(QueryBlockTipResponse)
	err := c.cc.Invoke(ctx, Query_BlockTip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Params queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Pubkeys queries current public key for deposit
	Pubkey(context.Context, *QueryPubkeyRequest) (*QueryPubkeyResponse, error)
	// DepositAddress queries current address for deposit
	DepositAddress(context.Context, *QueryDepositAddress) (*QueryDepositAddressResponse, error)
	// HasDeposited checks if a deposit transaction is confirmed on chain
	HasDeposited(context.Context, *QueryHasDeposited) (*QueryHasDepositedResponse, error)
	// Withdrawal queries all current public keys
	Withdrawal(context.Context, *QueryWithdrawalRequest) (*QueryWithdrawalResponse, error)
	// BlockTip queries current the latest confirmed bitcoin height by relayer
	BlockTip(context.Context, *QueryBlockTipRequest) (*QueryBlockTipResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Pubkey(context.Context, *QueryPubkeyRequest) (*QueryPubkeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pubkey not implemented")
}
func (UnimplementedQueryServer) DepositAddress(context.Context, *QueryDepositAddress) (*QueryDepositAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositAddress not implemented")
}
func (UnimplementedQueryServer) HasDeposited(context.Context, *QueryHasDeposited) (*QueryHasDepositedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasDeposited not implemented")
}
func (UnimplementedQueryServer) Withdrawal(context.Context, *QueryWithdrawalRequest) (*QueryWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdrawal not implemented")
}
func (UnimplementedQueryServer) BlockTip(context.Context, *QueryBlockTipRequest) (*QueryBlockTipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockTip not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pubkey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPubkeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pubkey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pubkey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pubkey(ctx, req.(*QueryPubkeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DepositAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDepositAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DepositAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DepositAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DepositAddress(ctx, req.(*QueryDepositAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HasDeposited_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHasDeposited)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HasDeposited(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_HasDeposited_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HasDeposited(ctx, req.(*QueryHasDeposited))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Withdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Withdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Withdrawal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Withdrawal(ctx, req.(*QueryWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BlockTip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlockTipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlockTip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BlockTip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlockTip(ctx, req.(*QueryBlockTipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goat.bitcoin.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Pubkey",
			Handler:    _Query_Pubkey_Handler,
		},
		{
			MethodName: "DepositAddress",
			Handler:    _Query_DepositAddress_Handler,
		},
		{
			MethodName: "HasDeposited",
			Handler:    _Query_HasDeposited_Handler,
		},
		{
			MethodName: "Withdrawal",
			Handler:    _Query_Withdrawal_Handler,
		},
		{
			MethodName: "BlockTip",
			Handler:    _Query_BlockTip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goat/bitcoin/v1/query.proto",
}
