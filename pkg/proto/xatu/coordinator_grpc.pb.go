// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/proto/xatu/coordinator.proto

package xatu

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

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorClient interface {
	CreateNodeRecords(ctx context.Context, in *CreateNodeRecordsRequest, opts ...grpc.CallOption) (*CreateNodeRecordsResponse, error)
	ListStalledExecutionNodeRecords(ctx context.Context, in *ListStalledExecutionNodeRecordsRequest, opts ...grpc.CallOption) (*ListStalledExecutionNodeRecordsResponse, error)
	CreateExecutionNodeRecordStatus(ctx context.Context, in *CreateExecutionNodeRecordStatusRequest, opts ...grpc.CallOption) (*CreateExecutionNodeRecordStatusResponse, error)
	CoordinateExecutionNodeRecords(ctx context.Context, in *CoordinateExecutionNodeRecordsRequest, opts ...grpc.CallOption) (*CoordinateExecutionNodeRecordsResponse, error)
}

type coordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorClient(cc grpc.ClientConnInterface) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) CreateNodeRecords(ctx context.Context, in *CreateNodeRecordsRequest, opts ...grpc.CallOption) (*CreateNodeRecordsResponse, error) {
	out := new(CreateNodeRecordsResponse)
	err := c.cc.Invoke(ctx, "/xatu.Coordinator/CreateNodeRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) ListStalledExecutionNodeRecords(ctx context.Context, in *ListStalledExecutionNodeRecordsRequest, opts ...grpc.CallOption) (*ListStalledExecutionNodeRecordsResponse, error) {
	out := new(ListStalledExecutionNodeRecordsResponse)
	err := c.cc.Invoke(ctx, "/xatu.Coordinator/ListStalledExecutionNodeRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CreateExecutionNodeRecordStatus(ctx context.Context, in *CreateExecutionNodeRecordStatusRequest, opts ...grpc.CallOption) (*CreateExecutionNodeRecordStatusResponse, error) {
	out := new(CreateExecutionNodeRecordStatusResponse)
	err := c.cc.Invoke(ctx, "/xatu.Coordinator/CreateExecutionNodeRecordStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinateExecutionNodeRecords(ctx context.Context, in *CoordinateExecutionNodeRecordsRequest, opts ...grpc.CallOption) (*CoordinateExecutionNodeRecordsResponse, error) {
	out := new(CoordinateExecutionNodeRecordsResponse)
	err := c.cc.Invoke(ctx, "/xatu.Coordinator/CoordinateExecutionNodeRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
// All implementations must embed UnimplementedCoordinatorServer
// for forward compatibility
type CoordinatorServer interface {
	CreateNodeRecords(context.Context, *CreateNodeRecordsRequest) (*CreateNodeRecordsResponse, error)
	ListStalledExecutionNodeRecords(context.Context, *ListStalledExecutionNodeRecordsRequest) (*ListStalledExecutionNodeRecordsResponse, error)
	CreateExecutionNodeRecordStatus(context.Context, *CreateExecutionNodeRecordStatusRequest) (*CreateExecutionNodeRecordStatusResponse, error)
	CoordinateExecutionNodeRecords(context.Context, *CoordinateExecutionNodeRecordsRequest) (*CoordinateExecutionNodeRecordsResponse, error)
	mustEmbedUnimplementedCoordinatorServer()
}

// UnimplementedCoordinatorServer must be embedded to have forward compatible implementations.
type UnimplementedCoordinatorServer struct {
}

func (UnimplementedCoordinatorServer) CreateNodeRecords(context.Context, *CreateNodeRecordsRequest) (*CreateNodeRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeRecords not implemented")
}
func (UnimplementedCoordinatorServer) ListStalledExecutionNodeRecords(context.Context, *ListStalledExecutionNodeRecordsRequest) (*ListStalledExecutionNodeRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStalledExecutionNodeRecords not implemented")
}
func (UnimplementedCoordinatorServer) CreateExecutionNodeRecordStatus(context.Context, *CreateExecutionNodeRecordStatusRequest) (*CreateExecutionNodeRecordStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExecutionNodeRecordStatus not implemented")
}
func (UnimplementedCoordinatorServer) CoordinateExecutionNodeRecords(context.Context, *CoordinateExecutionNodeRecordsRequest) (*CoordinateExecutionNodeRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinateExecutionNodeRecords not implemented")
}
func (UnimplementedCoordinatorServer) mustEmbedUnimplementedCoordinatorServer() {}

// UnsafeCoordinatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServer will
// result in compilation errors.
type UnsafeCoordinatorServer interface {
	mustEmbedUnimplementedCoordinatorServer()
}

func RegisterCoordinatorServer(s grpc.ServiceRegistrar, srv CoordinatorServer) {
	s.RegisterService(&Coordinator_ServiceDesc, srv)
}

func _Coordinator_CreateNodeRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CreateNodeRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xatu.Coordinator/CreateNodeRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CreateNodeRecords(ctx, req.(*CreateNodeRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_ListStalledExecutionNodeRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStalledExecutionNodeRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).ListStalledExecutionNodeRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xatu.Coordinator/ListStalledExecutionNodeRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).ListStalledExecutionNodeRecords(ctx, req.(*ListStalledExecutionNodeRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CreateExecutionNodeRecordStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExecutionNodeRecordStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CreateExecutionNodeRecordStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xatu.Coordinator/CreateExecutionNodeRecordStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CreateExecutionNodeRecordStatus(ctx, req.(*CreateExecutionNodeRecordStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinateExecutionNodeRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoordinateExecutionNodeRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinateExecutionNodeRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xatu.Coordinator/CoordinateExecutionNodeRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinateExecutionNodeRecords(ctx, req.(*CoordinateExecutionNodeRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Coordinator_ServiceDesc is the grpc.ServiceDesc for Coordinator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coordinator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xatu.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNodeRecords",
			Handler:    _Coordinator_CreateNodeRecords_Handler,
		},
		{
			MethodName: "ListStalledExecutionNodeRecords",
			Handler:    _Coordinator_ListStalledExecutionNodeRecords_Handler,
		},
		{
			MethodName: "CreateExecutionNodeRecordStatus",
			Handler:    _Coordinator_CreateExecutionNodeRecordStatus_Handler,
		},
		{
			MethodName: "CoordinateExecutionNodeRecords",
			Handler:    _Coordinator_CoordinateExecutionNodeRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/xatu/coordinator.proto",
}
