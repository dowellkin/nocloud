// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ServicesServiceClient is the client API for ServicesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesServiceClient interface {
	TestConfig(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*TestConfigResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Service, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Service, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Service, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpResponse, error)
	Down(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (*DownResponse, error)
	PerformServiceAction(ctx context.Context, in *PerformActionRequest, opts ...grpc.CallOption) (*PerformActionResponse, error)
}

type servicesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesServiceClient(cc grpc.ClientConnInterface) ServicesServiceClient {
	return &servicesServiceClient{cc}
}

func (c *servicesServiceClient) TestConfig(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*TestConfigResponse, error) {
	out := new(TestConfigResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/TestConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpResponse, error) {
	out := new(UpResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/Up", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) Down(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (*DownResponse, error) {
	out := new(DownResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/Down", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) PerformServiceAction(ctx context.Context, in *PerformActionRequest, opts ...grpc.CallOption) (*PerformActionResponse, error) {
	out := new(PerformActionResponse)
	err := c.cc.Invoke(ctx, "/nocloud.services.ServicesService/PerformServiceAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServiceServer is the server API for ServicesService service.
// All implementations must embed UnimplementedServicesServiceServer
// for forward compatibility
type ServicesServiceServer interface {
	TestConfig(context.Context, *CreateRequest) (*TestConfigResponse, error)
	Create(context.Context, *CreateRequest) (*Service, error)
	Update(context.Context, *UpdateRequest) (*Service, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Get(context.Context, *GetRequest) (*Service, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Up(context.Context, *UpRequest) (*UpResponse, error)
	Down(context.Context, *DownRequest) (*DownResponse, error)
	PerformServiceAction(context.Context, *PerformActionRequest) (*PerformActionResponse, error)
	mustEmbedUnimplementedServicesServiceServer()
}

// UnimplementedServicesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServicesServiceServer struct {
}

func (UnimplementedServicesServiceServer) TestConfig(context.Context, *CreateRequest) (*TestConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConfig not implemented")
}
func (UnimplementedServicesServiceServer) Create(context.Context, *CreateRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServicesServiceServer) Update(context.Context, *UpdateRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServicesServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedServicesServiceServer) Get(context.Context, *GetRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServicesServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServicesServiceServer) Up(context.Context, *UpRequest) (*UpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (UnimplementedServicesServiceServer) Down(context.Context, *DownRequest) (*DownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Down not implemented")
}
func (UnimplementedServicesServiceServer) PerformServiceAction(context.Context, *PerformActionRequest) (*PerformActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformServiceAction not implemented")
}
func (UnimplementedServicesServiceServer) mustEmbedUnimplementedServicesServiceServer() {}

// UnsafeServicesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServiceServer will
// result in compilation errors.
type UnsafeServicesServiceServer interface {
	mustEmbedUnimplementedServicesServiceServer()
}

func RegisterServicesServiceServer(s grpc.ServiceRegistrar, srv ServicesServiceServer) {
	s.RegisterService(&ServicesService_ServiceDesc, srv)
}

func _ServicesService_TestConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).TestConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/TestConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).TestConfig(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).Up(ctx, req.(*UpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_Down_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).Down(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/Down",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).Down(ctx, req.(*DownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_PerformServiceAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).PerformServiceAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.services.ServicesService/PerformServiceAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).PerformServiceAction(ctx, req.(*PerformActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicesService_ServiceDesc is the grpc.ServiceDesc for ServicesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nocloud.services.ServicesService",
	HandlerType: (*ServicesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestConfig",
			Handler:    _ServicesService_TestConfig_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ServicesService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServicesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ServicesService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServicesService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServicesService_List_Handler,
		},
		{
			MethodName: "Up",
			Handler:    _ServicesService_Up_Handler,
		},
		{
			MethodName: "Down",
			Handler:    _ServicesService_Down_Handler,
		},
		{
			MethodName: "PerformServiceAction",
			Handler:    _ServicesService_PerformServiceAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/proto/services.proto",
}
