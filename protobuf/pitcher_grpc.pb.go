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
	GetTracks(ctx context.Context, in *TracksRequest, opts ...grpc.CallOption) (*TracksResponse, error)
	GetCoverArt(ctx context.Context, in *CoverArtRequest, opts ...grpc.CallOption) (*CoverArtResponse, error)
	GetArtist(ctx context.Context, in *ArtistRequest, opts ...grpc.CallOption) (*ArtistResponse, error)
	GetArtists(ctx context.Context, in *ArtistsRequest, opts ...grpc.CallOption) (*ArtistsResponse, error)
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

func (c *pitcherClient) GetTracks(ctx context.Context, in *TracksRequest, opts ...grpc.CallOption) (*TracksResponse, error) {
	out := new(TracksResponse)
	err := c.cc.Invoke(ctx, "/Pitcher/GetTracks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pitcherClient) GetCoverArt(ctx context.Context, in *CoverArtRequest, opts ...grpc.CallOption) (*CoverArtResponse, error) {
	out := new(CoverArtResponse)
	err := c.cc.Invoke(ctx, "/Pitcher/GetCoverArt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pitcherClient) GetArtist(ctx context.Context, in *ArtistRequest, opts ...grpc.CallOption) (*ArtistResponse, error) {
	out := new(ArtistResponse)
	err := c.cc.Invoke(ctx, "/Pitcher/GetArtist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pitcherClient) GetArtists(ctx context.Context, in *ArtistsRequest, opts ...grpc.CallOption) (*ArtistsResponse, error) {
	out := new(ArtistsResponse)
	err := c.cc.Invoke(ctx, "/Pitcher/GetArtists", in, out, opts...)
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
	GetTracks(context.Context, *TracksRequest) (*TracksResponse, error)
	GetCoverArt(context.Context, *CoverArtRequest) (*CoverArtResponse, error)
	GetArtist(context.Context, *ArtistRequest) (*ArtistResponse, error)
	GetArtists(context.Context, *ArtistsRequest) (*ArtistsResponse, error)
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
func (UnimplementedPitcherServer) GetTracks(context.Context, *TracksRequest) (*TracksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTracks not implemented")
}
func (UnimplementedPitcherServer) GetCoverArt(context.Context, *CoverArtRequest) (*CoverArtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoverArt not implemented")
}
func (UnimplementedPitcherServer) GetArtist(context.Context, *ArtistRequest) (*ArtistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtist not implemented")
}
func (UnimplementedPitcherServer) GetArtists(context.Context, *ArtistsRequest) (*ArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtists not implemented")
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

func _Pitcher_GetTracks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TracksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitcherServer).GetTracks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pitcher/GetTracks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitcherServer).GetTracks(ctx, req.(*TracksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pitcher_GetCoverArt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoverArtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitcherServer).GetCoverArt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pitcher/GetCoverArt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitcherServer).GetCoverArt(ctx, req.(*CoverArtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pitcher_GetArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitcherServer).GetArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pitcher/GetArtist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitcherServer).GetArtist(ctx, req.(*ArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pitcher_GetArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitcherServer).GetArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Pitcher/GetArtists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitcherServer).GetArtists(ctx, req.(*ArtistsRequest))
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
		{
			MethodName: "GetTracks",
			Handler:    _Pitcher_GetTracks_Handler,
		},
		{
			MethodName: "GetCoverArt",
			Handler:    _Pitcher_GetCoverArt_Handler,
		},
		{
			MethodName: "GetArtist",
			Handler:    _Pitcher_GetArtist_Handler,
		},
		{
			MethodName: "GetArtists",
			Handler:    _Pitcher_GetArtists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/pitcher.proto",
}
