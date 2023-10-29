// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: provision-pod.proto

package proto

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

// ProvisionPodServiceClient is the client API for ProvisionPodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvisionPodServiceClient interface {
	ProvisionPod(ctx context.Context, in *ProvisionPodRequest, opts ...grpc.CallOption) (*ProvisionPodResponse, error)
	UpdatePod(ctx context.Context, in *UpdatePodRequest, opts ...grpc.CallOption) (*ProvisionPodResponse, error)
	DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error)
}

type provisionPodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisionPodServiceClient(cc grpc.ClientConnInterface) ProvisionPodServiceClient {
	return &provisionPodServiceClient{cc}
}

func (c *provisionPodServiceClient) ProvisionPod(ctx context.Context, in *ProvisionPodRequest, opts ...grpc.CallOption) (*ProvisionPodResponse, error) {
	out := new(ProvisionPodResponse)
	err := c.cc.Invoke(ctx, "/apocryph.proto.v0.provisionPod.ProvisionPodService/ProvisionPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionPodServiceClient) UpdatePod(ctx context.Context, in *UpdatePodRequest, opts ...grpc.CallOption) (*ProvisionPodResponse, error) {
	out := new(ProvisionPodResponse)
	err := c.cc.Invoke(ctx, "/apocryph.proto.v0.provisionPod.ProvisionPodService/UpdatePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionPodServiceClient) DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error) {
	out := new(DeletePodResponse)
	err := c.cc.Invoke(ctx, "/apocryph.proto.v0.provisionPod.ProvisionPodService/DeletePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionPodServiceServer is the server API for ProvisionPodService service.
// All implementations must embed UnimplementedProvisionPodServiceServer
// for forward compatibility
type ProvisionPodServiceServer interface {
	ProvisionPod(context.Context, *ProvisionPodRequest) (*ProvisionPodResponse, error)
	UpdatePod(context.Context, *UpdatePodRequest) (*ProvisionPodResponse, error)
	DeletePod(context.Context, *DeletePodRequest) (*DeletePodResponse, error)
	mustEmbedUnimplementedProvisionPodServiceServer()
}

// UnimplementedProvisionPodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProvisionPodServiceServer struct {
}

func (UnimplementedProvisionPodServiceServer) ProvisionPod(context.Context, *ProvisionPodRequest) (*ProvisionPodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionPod not implemented")
}
func (UnimplementedProvisionPodServiceServer) UpdatePod(context.Context, *UpdatePodRequest) (*ProvisionPodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePod not implemented")
}
func (UnimplementedProvisionPodServiceServer) DeletePod(context.Context, *DeletePodRequest) (*DeletePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}
func (UnimplementedProvisionPodServiceServer) mustEmbedUnimplementedProvisionPodServiceServer() {}

// UnsafeProvisionPodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvisionPodServiceServer will
// result in compilation errors.
type UnsafeProvisionPodServiceServer interface {
	mustEmbedUnimplementedProvisionPodServiceServer()
}

func RegisterProvisionPodServiceServer(s grpc.ServiceRegistrar, srv ProvisionPodServiceServer) {
	s.RegisterService(&ProvisionPodService_ServiceDesc, srv)
}

func _ProvisionPodService_ProvisionPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionPodServiceServer).ProvisionPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apocryph.proto.v0.provisionPod.ProvisionPodService/ProvisionPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionPodServiceServer).ProvisionPod(ctx, req.(*ProvisionPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvisionPodService_UpdatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionPodServiceServer).UpdatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apocryph.proto.v0.provisionPod.ProvisionPodService/UpdatePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionPodServiceServer).UpdatePod(ctx, req.(*UpdatePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvisionPodService_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionPodServiceServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apocryph.proto.v0.provisionPod.ProvisionPodService/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionPodServiceServer).DeletePod(ctx, req.(*DeletePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProvisionPodService_ServiceDesc is the grpc.ServiceDesc for ProvisionPodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProvisionPodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apocryph.proto.v0.provisionPod.ProvisionPodService",
	HandlerType: (*ProvisionPodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProvisionPod",
			Handler:    _ProvisionPodService_ProvisionPod_Handler,
		},
		{
			MethodName: "UpdatePod",
			Handler:    _ProvisionPodService_UpdatePod_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _ProvisionPodService_DeletePod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provision-pod.proto",
}
