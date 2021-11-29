// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package snowflakeadminreporting

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

// SnowflakeAdminReportingServiceClient is the client API for SnowflakeAdminReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnowflakeAdminReportingServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetTransactionReport(ctx context.Context, in *Financialrequest, opts ...grpc.CallOption) (*Financialresponse, error)
	Externalvoucherrequest(ctx context.Context, in *Financialrequest, opts ...grpc.CallOption) (*Externalvoucherresponse, error)
}

type snowflakeAdminReportingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnowflakeAdminReportingServiceClient(cc grpc.ClientConnInterface) SnowflakeAdminReportingServiceClient {
	return &snowflakeAdminReportingServiceClient{cc}
}

func (c *snowflakeAdminReportingServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/snowflakeadminreporting.SnowflakeAdminReportingService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snowflakeAdminReportingServiceClient) GetTransactionReport(ctx context.Context, in *Financialrequest, opts ...grpc.CallOption) (*Financialresponse, error) {
	out := new(Financialresponse)
	err := c.cc.Invoke(ctx, "/snowflakeadminreporting.SnowflakeAdminReportingService/GetTransactionReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snowflakeAdminReportingServiceClient) Externalvoucherrequest(ctx context.Context, in *Financialrequest, opts ...grpc.CallOption) (*Externalvoucherresponse, error) {
	out := new(Externalvoucherresponse)
	err := c.cc.Invoke(ctx, "/snowflakeadminreporting.SnowflakeAdminReportingService/Externalvoucherrequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnowflakeAdminReportingServiceServer is the server API for SnowflakeAdminReportingService service.
// All implementations must embed UnimplementedSnowflakeAdminReportingServiceServer
// for forward compatibility
type SnowflakeAdminReportingServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	GetTransactionReport(context.Context, *Financialrequest) (*Financialresponse, error)
	Externalvoucherrequest(context.Context, *Financialrequest) (*Externalvoucherresponse, error)
	mustEmbedUnimplementedSnowflakeAdminReportingServiceServer()
}

// UnimplementedSnowflakeAdminReportingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSnowflakeAdminReportingServiceServer struct {
}

func (UnimplementedSnowflakeAdminReportingServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedSnowflakeAdminReportingServiceServer) GetTransactionReport(context.Context, *Financialrequest) (*Financialresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionReport not implemented")
}
func (UnimplementedSnowflakeAdminReportingServiceServer) Externalvoucherrequest(context.Context, *Financialrequest) (*Externalvoucherresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Externalvoucherrequest not implemented")
}
func (UnimplementedSnowflakeAdminReportingServiceServer) mustEmbedUnimplementedSnowflakeAdminReportingServiceServer() {
}

// UnsafeSnowflakeAdminReportingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnowflakeAdminReportingServiceServer will
// result in compilation errors.
type UnsafeSnowflakeAdminReportingServiceServer interface {
	mustEmbedUnimplementedSnowflakeAdminReportingServiceServer()
}

func RegisterSnowflakeAdminReportingServiceServer(s grpc.ServiceRegistrar, srv SnowflakeAdminReportingServiceServer) {
	s.RegisterService(&SnowflakeAdminReportingService_ServiceDesc, srv)
}

func _SnowflakeAdminReportingService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowflakeAdminReportingServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snowflakeadminreporting.SnowflakeAdminReportingService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowflakeAdminReportingServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnowflakeAdminReportingService_GetTransactionReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Financialrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowflakeAdminReportingServiceServer).GetTransactionReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snowflakeadminreporting.SnowflakeAdminReportingService/GetTransactionReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowflakeAdminReportingServiceServer).GetTransactionReport(ctx, req.(*Financialrequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnowflakeAdminReportingService_Externalvoucherrequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Financialrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowflakeAdminReportingServiceServer).Externalvoucherrequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snowflakeadminreporting.SnowflakeAdminReportingService/Externalvoucherrequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowflakeAdminReportingServiceServer).Externalvoucherrequest(ctx, req.(*Financialrequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SnowflakeAdminReportingService_ServiceDesc is the grpc.ServiceDesc for SnowflakeAdminReportingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnowflakeAdminReportingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snowflakeadminreporting.SnowflakeAdminReportingService",
	HandlerType: (*SnowflakeAdminReportingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SnowflakeAdminReportingService_SayHello_Handler,
		},
		{
			MethodName: "GetTransactionReport",
			Handler:    _SnowflakeAdminReportingService_GetTransactionReport_Handler,
		},
		{
			MethodName: "Externalvoucherrequest",
			Handler:    _SnowflakeAdminReportingService_Externalvoucherrequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/snowflakeadminreporting.proto",
}
