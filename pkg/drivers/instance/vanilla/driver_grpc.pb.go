// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vanilla

import (
	context "context"
	proto "github.com/slntopp/nocloud/pkg/services/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DriverServiceClient is the client API for DriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverServiceClient interface {
	ValidateConfigSyntax(ctx context.Context, in *proto.ValidateServiceConfigRequest, opts ...grpc.CallOption) (*proto.ValidateServiceConfigResponse, error)
	GetType(ctx context.Context, in *GetTypeRequest, opts ...grpc.CallOption) (*GetTypeResponse, error)
}

type driverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverServiceClient(cc grpc.ClientConnInterface) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) ValidateConfigSyntax(ctx context.Context, in *proto.ValidateServiceConfigRequest, opts ...grpc.CallOption) (*proto.ValidateServiceConfigResponse, error) {
	out := new(proto.ValidateServiceConfigResponse)
	err := c.cc.Invoke(ctx, "/nocloud.instance.driver.vanilla.DriverService/ValidateConfigSyntax", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) GetType(ctx context.Context, in *GetTypeRequest, opts ...grpc.CallOption) (*GetTypeResponse, error) {
	out := new(GetTypeResponse)
	err := c.cc.Invoke(ctx, "/nocloud.instance.driver.vanilla.DriverService/GetType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServiceServer is the server API for DriverService service.
// All implementations must embed UnimplementedDriverServiceServer
// for forward compatibility
type DriverServiceServer interface {
	ValidateConfigSyntax(context.Context, *proto.ValidateServiceConfigRequest) (*proto.ValidateServiceConfigResponse, error)
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)
	mustEmbedUnimplementedDriverServiceServer()
}

// UnimplementedDriverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDriverServiceServer struct {
}

func (UnimplementedDriverServiceServer) ValidateConfigSyntax(context.Context, *proto.ValidateServiceConfigRequest) (*proto.ValidateServiceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateConfigSyntax not implemented")
}
func (UnimplementedDriverServiceServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetType not implemented")
}
func (UnimplementedDriverServiceServer) mustEmbedUnimplementedDriverServiceServer() {}

// UnsafeDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServiceServer will
// result in compilation errors.
type UnsafeDriverServiceServer interface {
	mustEmbedUnimplementedDriverServiceServer()
}

func RegisterDriverServiceServer(s grpc.ServiceRegistrar, srv DriverServiceServer) {
	s.RegisterService(&DriverService_ServiceDesc, srv)
}

func _DriverService_ValidateConfigSyntax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.ValidateServiceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).ValidateConfigSyntax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.instance.driver.vanilla.DriverService/ValidateConfigSyntax",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).ValidateConfigSyntax(ctx, req.(*proto.ValidateServiceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_GetType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.instance.driver.vanilla.DriverService/GetType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetType(ctx, req.(*GetTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverService_ServiceDesc is the grpc.ServiceDesc for DriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nocloud.instance.driver.vanilla.DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateConfigSyntax",
			Handler:    _DriverService_ValidateConfigSyntax_Handler,
		},
		{
			MethodName: "GetType",
			Handler:    _DriverService_GetType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/drivers/instance/vanilla/driver.proto",
}
