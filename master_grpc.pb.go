// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: master.proto

package kpmasterproto_git

import (
	context "context"
	occupation "github.com/djoonta/kpmasterproto.git/occupation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OccupationService_FindID_FullMethodName  = "/kpmasterproto.OccupationService/FindID"
	OccupationService_FindAll_FullMethodName = "/kpmasterproto.OccupationService/FindAll"
)

// OccupationServiceClient is the client API for OccupationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OccupationServiceClient interface {
	FindID(ctx context.Context, in *occupation.OccupationFindIDRequest, opts ...grpc.CallOption) (*occupation.OccupationFindIDResponse, error)
	FindAll(ctx context.Context, in *occupation.OccupationFindAllRequest, opts ...grpc.CallOption) (*occupation.OccupationFindAllResponse, error)
}

type occupationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOccupationServiceClient(cc grpc.ClientConnInterface) OccupationServiceClient {
	return &occupationServiceClient{cc}
}

func (c *occupationServiceClient) FindID(ctx context.Context, in *occupation.OccupationFindIDRequest, opts ...grpc.CallOption) (*occupation.OccupationFindIDResponse, error) {
	out := new(occupation.OccupationFindIDResponse)
	err := c.cc.Invoke(ctx, OccupationService_FindID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occupationServiceClient) FindAll(ctx context.Context, in *occupation.OccupationFindAllRequest, opts ...grpc.CallOption) (*occupation.OccupationFindAllResponse, error) {
	out := new(occupation.OccupationFindAllResponse)
	err := c.cc.Invoke(ctx, OccupationService_FindAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OccupationServiceServer is the server API for OccupationService service.
// All implementations must embed UnimplementedOccupationServiceServer
// for forward compatibility
type OccupationServiceServer interface {
	FindID(context.Context, *occupation.OccupationFindIDRequest) (*occupation.OccupationFindIDResponse, error)
	FindAll(context.Context, *occupation.OccupationFindAllRequest) (*occupation.OccupationFindAllResponse, error)
	mustEmbedUnimplementedOccupationServiceServer()
}

// UnimplementedOccupationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOccupationServiceServer struct {
}

func (UnimplementedOccupationServiceServer) FindID(context.Context, *occupation.OccupationFindIDRequest) (*occupation.OccupationFindIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindID not implemented")
}
func (UnimplementedOccupationServiceServer) FindAll(context.Context, *occupation.OccupationFindAllRequest) (*occupation.OccupationFindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedOccupationServiceServer) mustEmbedUnimplementedOccupationServiceServer() {}

// UnsafeOccupationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OccupationServiceServer will
// result in compilation errors.
type UnsafeOccupationServiceServer interface {
	mustEmbedUnimplementedOccupationServiceServer()
}

func RegisterOccupationServiceServer(s grpc.ServiceRegistrar, srv OccupationServiceServer) {
	s.RegisterService(&OccupationService_ServiceDesc, srv)
}

func _OccupationService_FindID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(occupation.OccupationFindIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccupationServiceServer).FindID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OccupationService_FindID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccupationServiceServer).FindID(ctx, req.(*occupation.OccupationFindIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccupationService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(occupation.OccupationFindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccupationServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OccupationService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccupationServiceServer).FindAll(ctx, req.(*occupation.OccupationFindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OccupationService_ServiceDesc is the grpc.ServiceDesc for OccupationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OccupationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kpmasterproto.OccupationService",
	HandlerType: (*OccupationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindID",
			Handler:    _OccupationService_FindID_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _OccupationService_FindAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}