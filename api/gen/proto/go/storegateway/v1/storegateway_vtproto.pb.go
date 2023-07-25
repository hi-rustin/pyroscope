// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.0.0-20230724064928-69fa34dad347
// source: storegateway/v1/storegateway.proto

package storegatewayv1

import (
	context "context"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/ingester/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StoreGatewayServiceClient is the client API for StoreGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreGatewayServiceClient interface {
	MergeProfilesStacktraces(ctx context.Context, opts ...grpc.CallOption) (StoreGatewayService_MergeProfilesStacktracesClient, error)
	MergeProfilesLabels(ctx context.Context, opts ...grpc.CallOption) (StoreGatewayService_MergeProfilesLabelsClient, error)
	MergeProfilesPprof(ctx context.Context, opts ...grpc.CallOption) (StoreGatewayService_MergeProfilesPprofClient, error)
}

type storeGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreGatewayServiceClient(cc grpc.ClientConnInterface) StoreGatewayServiceClient {
	return &storeGatewayServiceClient{cc}
}

func (c *storeGatewayServiceClient) MergeProfilesStacktraces(ctx context.Context, opts ...grpc.CallOption) (StoreGatewayService_MergeProfilesStacktracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &StoreGatewayService_ServiceDesc.Streams[0], "/storegateway.v1.StoreGatewayService/MergeProfilesStacktraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeGatewayServiceMergeProfilesStacktracesClient{stream}
	return x, nil
}

type StoreGatewayService_MergeProfilesStacktracesClient interface {
	Send(*v1.MergeProfilesStacktracesRequest) error
	Recv() (*v1.MergeProfilesStacktracesResponse, error)
	grpc.ClientStream
}

type storeGatewayServiceMergeProfilesStacktracesClient struct {
	grpc.ClientStream
}

func (x *storeGatewayServiceMergeProfilesStacktracesClient) Send(m *v1.MergeProfilesStacktracesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storeGatewayServiceMergeProfilesStacktracesClient) Recv() (*v1.MergeProfilesStacktracesResponse, error) {
	m := new(v1.MergeProfilesStacktracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storeGatewayServiceClient) MergeProfilesLabels(ctx context.Context, opts ...grpc.CallOption) (StoreGatewayService_MergeProfilesLabelsClient, error) {
	stream, err := c.cc.NewStream(ctx, &StoreGatewayService_ServiceDesc.Streams[1], "/storegateway.v1.StoreGatewayService/MergeProfilesLabels", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeGatewayServiceMergeProfilesLabelsClient{stream}
	return x, nil
}

type StoreGatewayService_MergeProfilesLabelsClient interface {
	Send(*v1.MergeProfilesLabelsRequest) error
	Recv() (*v1.MergeProfilesLabelsResponse, error)
	grpc.ClientStream
}

type storeGatewayServiceMergeProfilesLabelsClient struct {
	grpc.ClientStream
}

func (x *storeGatewayServiceMergeProfilesLabelsClient) Send(m *v1.MergeProfilesLabelsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storeGatewayServiceMergeProfilesLabelsClient) Recv() (*v1.MergeProfilesLabelsResponse, error) {
	m := new(v1.MergeProfilesLabelsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storeGatewayServiceClient) MergeProfilesPprof(ctx context.Context, opts ...grpc.CallOption) (StoreGatewayService_MergeProfilesPprofClient, error) {
	stream, err := c.cc.NewStream(ctx, &StoreGatewayService_ServiceDesc.Streams[2], "/storegateway.v1.StoreGatewayService/MergeProfilesPprof", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeGatewayServiceMergeProfilesPprofClient{stream}
	return x, nil
}

type StoreGatewayService_MergeProfilesPprofClient interface {
	Send(*v1.MergeProfilesPprofRequest) error
	Recv() (*v1.MergeProfilesPprofResponse, error)
	grpc.ClientStream
}

type storeGatewayServiceMergeProfilesPprofClient struct {
	grpc.ClientStream
}

func (x *storeGatewayServiceMergeProfilesPprofClient) Send(m *v1.MergeProfilesPprofRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storeGatewayServiceMergeProfilesPprofClient) Recv() (*v1.MergeProfilesPprofResponse, error) {
	m := new(v1.MergeProfilesPprofResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StoreGatewayServiceServer is the server API for StoreGatewayService service.
// All implementations must embed UnimplementedStoreGatewayServiceServer
// for forward compatibility
type StoreGatewayServiceServer interface {
	MergeProfilesStacktraces(StoreGatewayService_MergeProfilesStacktracesServer) error
	MergeProfilesLabels(StoreGatewayService_MergeProfilesLabelsServer) error
	MergeProfilesPprof(StoreGatewayService_MergeProfilesPprofServer) error
	mustEmbedUnimplementedStoreGatewayServiceServer()
}

// UnimplementedStoreGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoreGatewayServiceServer struct {
}

func (UnimplementedStoreGatewayServiceServer) MergeProfilesStacktraces(StoreGatewayService_MergeProfilesStacktracesServer) error {
	return status.Errorf(codes.Unimplemented, "method MergeProfilesStacktraces not implemented")
}
func (UnimplementedStoreGatewayServiceServer) MergeProfilesLabels(StoreGatewayService_MergeProfilesLabelsServer) error {
	return status.Errorf(codes.Unimplemented, "method MergeProfilesLabels not implemented")
}
func (UnimplementedStoreGatewayServiceServer) MergeProfilesPprof(StoreGatewayService_MergeProfilesPprofServer) error {
	return status.Errorf(codes.Unimplemented, "method MergeProfilesPprof not implemented")
}
func (UnimplementedStoreGatewayServiceServer) mustEmbedUnimplementedStoreGatewayServiceServer() {}

// UnsafeStoreGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreGatewayServiceServer will
// result in compilation errors.
type UnsafeStoreGatewayServiceServer interface {
	mustEmbedUnimplementedStoreGatewayServiceServer()
}

func RegisterStoreGatewayServiceServer(s grpc.ServiceRegistrar, srv StoreGatewayServiceServer) {
	s.RegisterService(&StoreGatewayService_ServiceDesc, srv)
}

func _StoreGatewayService_MergeProfilesStacktraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StoreGatewayServiceServer).MergeProfilesStacktraces(&storeGatewayServiceMergeProfilesStacktracesServer{stream})
}

type StoreGatewayService_MergeProfilesStacktracesServer interface {
	Send(*v1.MergeProfilesStacktracesResponse) error
	Recv() (*v1.MergeProfilesStacktracesRequest, error)
	grpc.ServerStream
}

type storeGatewayServiceMergeProfilesStacktracesServer struct {
	grpc.ServerStream
}

func (x *storeGatewayServiceMergeProfilesStacktracesServer) Send(m *v1.MergeProfilesStacktracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storeGatewayServiceMergeProfilesStacktracesServer) Recv() (*v1.MergeProfilesStacktracesRequest, error) {
	m := new(v1.MergeProfilesStacktracesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StoreGatewayService_MergeProfilesLabels_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StoreGatewayServiceServer).MergeProfilesLabels(&storeGatewayServiceMergeProfilesLabelsServer{stream})
}

type StoreGatewayService_MergeProfilesLabelsServer interface {
	Send(*v1.MergeProfilesLabelsResponse) error
	Recv() (*v1.MergeProfilesLabelsRequest, error)
	grpc.ServerStream
}

type storeGatewayServiceMergeProfilesLabelsServer struct {
	grpc.ServerStream
}

func (x *storeGatewayServiceMergeProfilesLabelsServer) Send(m *v1.MergeProfilesLabelsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storeGatewayServiceMergeProfilesLabelsServer) Recv() (*v1.MergeProfilesLabelsRequest, error) {
	m := new(v1.MergeProfilesLabelsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StoreGatewayService_MergeProfilesPprof_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StoreGatewayServiceServer).MergeProfilesPprof(&storeGatewayServiceMergeProfilesPprofServer{stream})
}

type StoreGatewayService_MergeProfilesPprofServer interface {
	Send(*v1.MergeProfilesPprofResponse) error
	Recv() (*v1.MergeProfilesPprofRequest, error)
	grpc.ServerStream
}

type storeGatewayServiceMergeProfilesPprofServer struct {
	grpc.ServerStream
}

func (x *storeGatewayServiceMergeProfilesPprofServer) Send(m *v1.MergeProfilesPprofResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storeGatewayServiceMergeProfilesPprofServer) Recv() (*v1.MergeProfilesPprofRequest, error) {
	m := new(v1.MergeProfilesPprofRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StoreGatewayService_ServiceDesc is the grpc.ServiceDesc for StoreGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storegateway.v1.StoreGatewayService",
	HandlerType: (*StoreGatewayServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MergeProfilesStacktraces",
			Handler:       _StoreGatewayService_MergeProfilesStacktraces_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MergeProfilesLabels",
			Handler:       _StoreGatewayService_MergeProfilesLabels_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MergeProfilesPprof",
			Handler:       _StoreGatewayService_MergeProfilesPprof_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "storegateway/v1/storegateway.proto",
}