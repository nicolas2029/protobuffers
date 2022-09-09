// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: classification/classification.proto

package classification

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

// ValidateOrderClient is the client API for ValidateOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidateOrderClient interface {
	ClassImg(ctx context.Context, in *RequestClassImg, opts ...grpc.CallOption) (*ResponseClassImg, error)
}

type validateOrderClient struct {
	cc grpc.ClientConnInterface
}

func NewValidateOrderClient(cc grpc.ClientConnInterface) ValidateOrderClient {
	return &validateOrderClient{cc}
}

func (c *validateOrderClient) ClassImg(ctx context.Context, in *RequestClassImg, opts ...grpc.CallOption) (*ResponseClassImg, error) {
	out := new(ResponseClassImg)
	err := c.cc.Invoke(ctx, "/ValidateOrder/ClassImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidateOrderServer is the server API for ValidateOrder service.
// All implementations must embed UnimplementedValidateOrderServer
// for forward compatibility
type ValidateOrderServer interface {
	ClassImg(context.Context, *RequestClassImg) (*ResponseClassImg, error)
	mustEmbedUnimplementedValidateOrderServer()
}

// UnimplementedValidateOrderServer must be embedded to have forward compatible implementations.
type UnimplementedValidateOrderServer struct {
}

func (UnimplementedValidateOrderServer) ClassImg(context.Context, *RequestClassImg) (*ResponseClassImg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassImg not implemented")
}
func (UnimplementedValidateOrderServer) mustEmbedUnimplementedValidateOrderServer() {}

// UnsafeValidateOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidateOrderServer will
// result in compilation errors.
type UnsafeValidateOrderServer interface {
	mustEmbedUnimplementedValidateOrderServer()
}

func RegisterValidateOrderServer(s grpc.ServiceRegistrar, srv ValidateOrderServer) {
	s.RegisterService(&ValidateOrder_ServiceDesc, srv)
}

func _ValidateOrder_ClassImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestClassImg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidateOrderServer).ClassImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ValidateOrder/ClassImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidateOrderServer).ClassImg(ctx, req.(*RequestClassImg))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidateOrder_ServiceDesc is the grpc.ServiceDesc for ValidateOrder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidateOrder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ValidateOrder",
	HandlerType: (*ValidateOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClassImg",
			Handler:    _ValidateOrder_ClassImg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "classification/classification.proto",
}
