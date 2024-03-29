// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: sorm/sorm.proto

package sorm

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

// SormDataCollectionServiceClient is the client API for SormDataCollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SormDataCollectionServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type sormDataCollectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSormDataCollectionServiceClient(cc grpc.ClientConnInterface) SormDataCollectionServiceClient {
	return &sormDataCollectionServiceClient{cc}
}

func (c *sormDataCollectionServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/sorm.SormDataCollectionService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SormDataCollectionServiceServer is the server API for SormDataCollectionService service.
// All implementations must embed UnimplementedSormDataCollectionServiceServer
// for forward compatibility
type SormDataCollectionServiceServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	mustEmbedUnimplementedSormDataCollectionServiceServer()
}

// UnimplementedSormDataCollectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSormDataCollectionServiceServer struct {
}

func (UnimplementedSormDataCollectionServiceServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSormDataCollectionServiceServer) mustEmbedUnimplementedSormDataCollectionServiceServer() {
}

// UnsafeSormDataCollectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SormDataCollectionServiceServer will
// result in compilation errors.
type UnsafeSormDataCollectionServiceServer interface {
	mustEmbedUnimplementedSormDataCollectionServiceServer()
}

func RegisterSormDataCollectionServiceServer(s grpc.ServiceRegistrar, srv SormDataCollectionServiceServer) {
	s.RegisterService(&SormDataCollectionService_ServiceDesc, srv)
}

func _SormDataCollectionService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SormDataCollectionServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sorm.SormDataCollectionService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SormDataCollectionServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SormDataCollectionService_ServiceDesc is the grpc.ServiceDesc for SormDataCollectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SormDataCollectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sorm.SormDataCollectionService",
	HandlerType: (*SormDataCollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SormDataCollectionService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sorm/sorm.proto",
}
