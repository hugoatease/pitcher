// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PitcherClient is the client API for Pitcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PitcherClient interface {
	MatchTrack(ctx context.Context, in *MatchingRequest, opts ...grpc.CallOption) (*MatchingResponse, error)
	GetTrack(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*TrackResponse, error)
}

type pitcherClient struct {
	cc grpc.ClientConnInterface
}

func NewPitcherClient(cc grpc.ClientConnInterface) PitcherClient {
	return &pitcherClient{cc}
}

func (c *pitcherClient) MatchTrack(ctx context.Context, in *MatchingRequest, opts ...grpc.CallOption) (*MatchingResponse, error) {
	out := new(MatchingResponse)
	err := c.cc.Invoke(ctx, "/Pitcher/MatchTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pitcherClient) GetTrack(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*TrackResponse, error) {
	out := new(TrackResponse)
	err := c.cc.Invoke(ctx, "/Pitcher/GetTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PitcherServer is the server API for Pitcher service.
// All implementations must embed UnimplementedPitcherServer
// for forward compatibility
type PitcherServer interface {
	MatchTrack(context.Context, *MatchingRequest) (*MatchingResponse, error)
	GetTrack(context.Context, *TrackRequest) (*TrackResponse, error)
	mustEmbedUnimplementedPitcherServer()
}

// UnimplementedPitcherServer must be embedded to have forward compatible implementations.
type UnimplementedPitcherServer struct {
}

func (UnimplementedPitcherServer) MatchTrack(context.Context, *MatchingRequest) (*MatchingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatchTrack not implemented")
}
func (UnimplementedPitcherServer) GetTrack(context.Context, *TrackRequest) (*TrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrack not implemented")
}
func (UnimplementedPitcherServer) mustEmbedUnimplementedPitcherServer() {}

// UnsafePitcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PitcherServer will
// result in compilation errors.
type UnsafePitcherServer interface {
	mustEmbedUnimplementedPitcherServer()
}

func RegisterPitcherServer(s grpc.ServiceRegistrar, srv PitcherServer) {
	s.RegisterService(&_Pitcher_serviceDesc, srv)
}

func _Pitcher_MatchTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitcherServer).MatchTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pitcher/MatchTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitcherServer).MatchTrack(ctx, req.(*MatchingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pitcher_GetTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitcherServer).GetTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pitcher/GetTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitcherServer).GetTrack(ctx, req.(*TrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pitcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Pitcher",
	HandlerType: (*PitcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MatchTrack",
			Handler:    _Pitcher_MatchTrack_Handler,
		},
		{
			MethodName: "GetTrack",
			Handler:    _Pitcher_GetTrack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/pitcher.proto",
}