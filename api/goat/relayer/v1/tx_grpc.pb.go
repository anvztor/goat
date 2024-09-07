// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: goat/relayer/v1/tx.proto

package relayerv1

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
	Msg_NewVoter_FullMethodName       = "/goat.relayer.v1.Msg/NewVoter"
	Msg_AcceptProposer_FullMethodName = "/goat.relayer.v1.Msg/AcceptProposer"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// NewVoter adds a pending voter with the online proof
	// an address is approved to engage as a relayer voter
	// the voter must send online proof to current proposer
	// then proposer initiates NewVoter tx to add the voter to the active list
	NewVoter(ctx context.Context, in *MsgNewVoterRequest, opts ...grpc.CallOption) (*MsgNewVoterResponse, error)
	// AcceptProposer accepts the proposer role
	// if a voter is elected as a proposer
	// the voter must initiate AcceptProposer tx to prove online status
	// the consensus layer can wait for Param.AcceptProposerTimeout at most
	// Note: if the timeout is zero, the consensus layer won't use it
	AcceptProposer(ctx context.Context, in *MsgAcceptProposerRequest, opts ...grpc.CallOption) (*MsgAcceptProposerResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) NewVoter(ctx context.Context, in *MsgNewVoterRequest, opts ...grpc.CallOption) (*MsgNewVoterResponse, error) {
	out := new(MsgNewVoterResponse)
	err := c.cc.Invoke(ctx, Msg_NewVoter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptProposer(ctx context.Context, in *MsgAcceptProposerRequest, opts ...grpc.CallOption) (*MsgAcceptProposerResponse, error) {
	out := new(MsgAcceptProposerResponse)
	err := c.cc.Invoke(ctx, Msg_AcceptProposer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// NewVoter adds a pending voter with the online proof
	// an address is approved to engage as a relayer voter
	// the voter must send online proof to current proposer
	// then proposer initiates NewVoter tx to add the voter to the active list
	NewVoter(context.Context, *MsgNewVoterRequest) (*MsgNewVoterResponse, error)
	// AcceptProposer accepts the proposer role
	// if a voter is elected as a proposer
	// the voter must initiate AcceptProposer tx to prove online status
	// the consensus layer can wait for Param.AcceptProposerTimeout at most
	// Note: if the timeout is zero, the consensus layer won't use it
	AcceptProposer(context.Context, *MsgAcceptProposerRequest) (*MsgAcceptProposerResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) NewVoter(context.Context, *MsgNewVoterRequest) (*MsgNewVoterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewVoter not implemented")
}
func (UnimplementedMsgServer) AcceptProposer(context.Context, *MsgAcceptProposerRequest) (*MsgAcceptProposerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptProposer not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_NewVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewVoter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewVoter(ctx, req.(*MsgNewVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptProposer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptProposerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptProposer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AcceptProposer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptProposer(ctx, req.(*MsgAcceptProposerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goat.relayer.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewVoter",
			Handler:    _Msg_NewVoter_Handler,
		},
		{
			MethodName: "AcceptProposer",
			Handler:    _Msg_AcceptProposer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goat/relayer/v1/tx.proto",
}
