// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: stream.proto

package pb

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

// VideoStreamClient is the client API for VideoStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoStreamClient interface {
	ReceiveVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoStream_ReceiveVideoClient, error)
}

type videoStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStreamClient(cc grpc.ClientConnInterface) VideoStreamClient {
	return &videoStreamClient{cc}
}

func (c *videoStreamClient) ReceiveVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoStream_ReceiveVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &VideoStream_ServiceDesc.Streams[0], "/v_stream.VideoStream/ReceiveVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoStreamReceiveVideoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VideoStream_ReceiveVideoClient interface {
	Recv() (*VideoResponse, error)
	grpc.ClientStream
}

type videoStreamReceiveVideoClient struct {
	grpc.ClientStream
}

func (x *videoStreamReceiveVideoClient) Recv() (*VideoResponse, error) {
	m := new(VideoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoStreamServer is the server API for VideoStream service.
// All implementations must embed UnimplementedVideoStreamServer
// for forward compatibility
type VideoStreamServer interface {
	ReceiveVideo(*VideoRequest, VideoStream_ReceiveVideoServer) error
	mustEmbedUnimplementedVideoStreamServer()
}

// UnimplementedVideoStreamServer must be embedded to have forward compatible implementations.
type UnimplementedVideoStreamServer struct {
}

func (UnimplementedVideoStreamServer) ReceiveVideo(*VideoRequest, VideoStream_ReceiveVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveVideo not implemented")
}
func (UnimplementedVideoStreamServer) mustEmbedUnimplementedVideoStreamServer() {}

// UnsafeVideoStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoStreamServer will
// result in compilation errors.
type UnsafeVideoStreamServer interface {
	mustEmbedUnimplementedVideoStreamServer()
}

func RegisterVideoStreamServer(s grpc.ServiceRegistrar, srv VideoStreamServer) {
	s.RegisterService(&VideoStream_ServiceDesc, srv)
}

func _VideoStream_ReceiveVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VideoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoStreamServer).ReceiveVideo(m, &videoStreamReceiveVideoServer{stream})
}

type VideoStream_ReceiveVideoServer interface {
	Send(*VideoResponse) error
	grpc.ServerStream
}

type videoStreamReceiveVideoServer struct {
	grpc.ServerStream
}

func (x *videoStreamReceiveVideoServer) Send(m *VideoResponse) error {
	return x.ServerStream.SendMsg(m)
}

// VideoStream_ServiceDesc is the grpc.ServiceDesc for VideoStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v_stream.VideoStream",
	HandlerType: (*VideoStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveVideo",
			Handler:       _VideoStream_ReceiveVideo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}
