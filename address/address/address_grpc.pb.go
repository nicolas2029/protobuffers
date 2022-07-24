// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: address/address.proto

package address

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

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressServiceClient interface {
	CreateDelivery(ctx context.Context, in *Delivery, opts ...grpc.CallOption) (*ID, error)
	GetAllByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResponseAll, error)
	DeleteByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResponseDelete, error)
	GetByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*Address, error)
	CreateEstablishment(ctx context.Context, in *Address, opts ...grpc.CallOption) (*ID, error)
	DeleteEstablishment(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ResponseDelete, error)
	Search(ctx context.Context, in *SearchAddress, opts ...grpc.CallOption) (*ResponseAll, error)
	Nearest(ctx context.Context, in *User, opts ...grpc.CallOption) (*ID, error)
}

type addressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressServiceClient(cc grpc.ClientConnInterface) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) CreateDelivery(ctx context.Context, in *Delivery, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/CreateDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) GetAllByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResponseAll, error) {
	out := new(ResponseAll)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/GetAllByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) DeleteByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResponseDelete, error) {
	out := new(ResponseDelete)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/DeleteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) GetByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) CreateEstablishment(ctx context.Context, in *Address, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/CreateEstablishment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) DeleteEstablishment(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ResponseDelete, error) {
	out := new(ResponseDelete)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/DeleteEstablishment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Search(ctx context.Context, in *SearchAddress, opts ...grpc.CallOption) (*ResponseAll, error) {
	out := new(ResponseAll)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Nearest(ctx context.Context, in *User, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/proto.address.address.AddressService/Nearest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
// All implementations must embed UnimplementedAddressServiceServer
// for forward compatibility
type AddressServiceServer interface {
	CreateDelivery(context.Context, *Delivery) (*ID, error)
	GetAllByUser(context.Context, *User) (*ResponseAll, error)
	DeleteByID(context.Context, *User) (*ResponseDelete, error)
	GetByID(context.Context, *User) (*Address, error)
	CreateEstablishment(context.Context, *Address) (*ID, error)
	DeleteEstablishment(context.Context, *ID) (*ResponseDelete, error)
	Search(context.Context, *SearchAddress) (*ResponseAll, error)
	Nearest(context.Context, *User) (*ID, error)
	mustEmbedUnimplementedAddressServiceServer()
}

// UnimplementedAddressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (UnimplementedAddressServiceServer) CreateDelivery(context.Context, *Delivery) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDelivery not implemented")
}
func (UnimplementedAddressServiceServer) GetAllByUser(context.Context, *User) (*ResponseAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByUser not implemented")
}
func (UnimplementedAddressServiceServer) DeleteByID(context.Context, *User) (*ResponseDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedAddressServiceServer) GetByID(context.Context, *User) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedAddressServiceServer) CreateEstablishment(context.Context, *Address) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEstablishment not implemented")
}
func (UnimplementedAddressServiceServer) DeleteEstablishment(context.Context, *ID) (*ResponseDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEstablishment not implemented")
}
func (UnimplementedAddressServiceServer) Search(context.Context, *SearchAddress) (*ResponseAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAddressServiceServer) Nearest(context.Context, *User) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nearest not implemented")
}
func (UnimplementedAddressServiceServer) mustEmbedUnimplementedAddressServiceServer() {}

// UnsafeAddressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServiceServer will
// result in compilation errors.
type UnsafeAddressServiceServer interface {
	mustEmbedUnimplementedAddressServiceServer()
}

func RegisterAddressServiceServer(s grpc.ServiceRegistrar, srv AddressServiceServer) {
	s.RegisterService(&AddressService_ServiceDesc, srv)
}

func _AddressService_CreateDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Delivery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).CreateDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/CreateDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).CreateDelivery(ctx, req.(*Delivery))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_GetAllByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GetAllByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/GetAllByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GetAllByUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/DeleteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).DeleteByID(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GetByID(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_CreateEstablishment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).CreateEstablishment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/CreateEstablishment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).CreateEstablishment(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_DeleteEstablishment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).DeleteEstablishment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/DeleteEstablishment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).DeleteEstablishment(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Search(ctx, req.(*SearchAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Nearest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Nearest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.address.address.AddressService/Nearest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Nearest(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressService_ServiceDesc is the grpc.ServiceDesc for AddressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.address.address.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDelivery",
			Handler:    _AddressService_CreateDelivery_Handler,
		},
		{
			MethodName: "GetAllByUser",
			Handler:    _AddressService_GetAllByUser_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _AddressService_DeleteByID_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _AddressService_GetByID_Handler,
		},
		{
			MethodName: "CreateEstablishment",
			Handler:    _AddressService_CreateEstablishment_Handler,
		},
		{
			MethodName: "DeleteEstablishment",
			Handler:    _AddressService_DeleteEstablishment_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _AddressService_Search_Handler,
		},
		{
			MethodName: "Nearest",
			Handler:    _AddressService_Nearest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address/address.proto",
}
