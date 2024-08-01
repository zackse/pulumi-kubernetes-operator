// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: agent.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AutomationService_WhoAmI_FullMethodName       = "/agent.AutomationService/WhoAmI"
	AutomationService_Info_FullMethodName         = "/agent.AutomationService/Info"
	AutomationService_SetAllConfig_FullMethodName = "/agent.AutomationService/SetAllConfig"
	AutomationService_Install_FullMethodName      = "/agent.AutomationService/Install"
	AutomationService_Preview_FullMethodName      = "/agent.AutomationService/Preview"
	AutomationService_Refresh_FullMethodName      = "/agent.AutomationService/Refresh"
	AutomationService_Up_FullMethodName           = "/agent.AutomationService/Up"
	AutomationService_Destroy_FullMethodName      = "/agent.AutomationService/Destroy"
)

// AutomationServiceClient is the client API for AutomationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutomationServiceClient interface {
	// *
	// WhoAmI returns detailed information about the currently logged-in Pulumi identity.
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResult, error)
	// *
	// Info returns information about the given stack.
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResult, error)
	// *
	// SetAllConfig sets multiple configuration values for a stack.
	SetAllConfig(ctx context.Context, in *SetAllConfigRequest, opts ...grpc.CallOption) (*SetAllConfigResult, error)
	// *
	// Install installs the Pulumi plugins and dependencies.
	Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*InstallResult, error)
	// *
	// Preview performs a dry-run update to a stack, returning the expected changes.
	Preview(ctx context.Context, in *PreviewRequest, opts ...grpc.CallOption) (AutomationService_PreviewClient, error)
	// *
	// Refresh updates the resources in a stack to match the current state of the cloud provider.
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (AutomationService_RefreshClient, error)
	// *
	// Up creates or updates the resources in a stack, returning the actual changes.
	Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (AutomationService_UpClient, error)
	// *
	// Destroy deletes all resources in a stack.
	Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (AutomationService_DestroyClient, error)
}

type automationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutomationServiceClient(cc grpc.ClientConnInterface) AutomationServiceClient {
	return &automationServiceClient{cc}
}

func (c *automationServiceClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WhoAmIResult)
	err := c.cc.Invoke(ctx, AutomationService_WhoAmI_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InfoResult)
	err := c.cc.Invoke(ctx, AutomationService_Info_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) SetAllConfig(ctx context.Context, in *SetAllConfigRequest, opts ...grpc.CallOption) (*SetAllConfigResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetAllConfigResult)
	err := c.cc.Invoke(ctx, AutomationService_SetAllConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*InstallResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InstallResult)
	err := c.cc.Invoke(ctx, AutomationService_Install_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) Preview(ctx context.Context, in *PreviewRequest, opts ...grpc.CallOption) (AutomationService_PreviewClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[0], AutomationService_Preview_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &automationServicePreviewClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AutomationService_PreviewClient interface {
	Recv() (*PreviewStream, error)
	grpc.ClientStream
}

type automationServicePreviewClient struct {
	grpc.ClientStream
}

func (x *automationServicePreviewClient) Recv() (*PreviewStream, error) {
	m := new(PreviewStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *automationServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (AutomationService_RefreshClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[1], AutomationService_Refresh_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &automationServiceRefreshClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AutomationService_RefreshClient interface {
	Recv() (*RefreshStream, error)
	grpc.ClientStream
}

type automationServiceRefreshClient struct {
	grpc.ClientStream
}

func (x *automationServiceRefreshClient) Recv() (*RefreshStream, error) {
	m := new(RefreshStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *automationServiceClient) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (AutomationService_UpClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[2], AutomationService_Up_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &automationServiceUpClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AutomationService_UpClient interface {
	Recv() (*UpStream, error)
	grpc.ClientStream
}

type automationServiceUpClient struct {
	grpc.ClientStream
}

func (x *automationServiceUpClient) Recv() (*UpStream, error) {
	m := new(UpStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *automationServiceClient) Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (AutomationService_DestroyClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[3], AutomationService_Destroy_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &automationServiceDestroyClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AutomationService_DestroyClient interface {
	Recv() (*DestroyStream, error)
	grpc.ClientStream
}

type automationServiceDestroyClient struct {
	grpc.ClientStream
}

func (x *automationServiceDestroyClient) Recv() (*DestroyStream, error) {
	m := new(DestroyStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AutomationServiceServer is the server API for AutomationService service.
// All implementations must embed UnimplementedAutomationServiceServer
// for forward compatibility
type AutomationServiceServer interface {
	// *
	// WhoAmI returns detailed information about the currently logged-in Pulumi identity.
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResult, error)
	// *
	// Info returns information about the given stack.
	Info(context.Context, *InfoRequest) (*InfoResult, error)
	// *
	// SetAllConfig sets multiple configuration values for a stack.
	SetAllConfig(context.Context, *SetAllConfigRequest) (*SetAllConfigResult, error)
	// *
	// Install installs the Pulumi plugins and dependencies.
	Install(context.Context, *InstallRequest) (*InstallResult, error)
	// *
	// Preview performs a dry-run update to a stack, returning the expected changes.
	Preview(*PreviewRequest, AutomationService_PreviewServer) error
	// *
	// Refresh updates the resources in a stack to match the current state of the cloud provider.
	Refresh(*RefreshRequest, AutomationService_RefreshServer) error
	// *
	// Up creates or updates the resources in a stack, returning the actual changes.
	Up(*UpRequest, AutomationService_UpServer) error
	// *
	// Destroy deletes all resources in a stack.
	Destroy(*DestroyRequest, AutomationService_DestroyServer) error
	mustEmbedUnimplementedAutomationServiceServer()
}

// UnimplementedAutomationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAutomationServiceServer struct {
}

func (UnimplementedAutomationServiceServer) WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedAutomationServiceServer) Info(context.Context, *InfoRequest) (*InfoResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedAutomationServiceServer) SetAllConfig(context.Context, *SetAllConfigRequest) (*SetAllConfigResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAllConfig not implemented")
}
func (UnimplementedAutomationServiceServer) Install(context.Context, *InstallRequest) (*InstallResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedAutomationServiceServer) Preview(*PreviewRequest, AutomationService_PreviewServer) error {
	return status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedAutomationServiceServer) Refresh(*RefreshRequest, AutomationService_RefreshServer) error {
	return status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAutomationServiceServer) Up(*UpRequest, AutomationService_UpServer) error {
	return status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (UnimplementedAutomationServiceServer) Destroy(*DestroyRequest, AutomationService_DestroyServer) error {
	return status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedAutomationServiceServer) mustEmbedUnimplementedAutomationServiceServer() {}

// UnsafeAutomationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutomationServiceServer will
// result in compilation errors.
type UnsafeAutomationServiceServer interface {
	mustEmbedUnimplementedAutomationServiceServer()
}

func RegisterAutomationServiceServer(s grpc.ServiceRegistrar, srv AutomationServiceServer) {
	s.RegisterService(&AutomationService_ServiceDesc, srv)
}

func _AutomationService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutomationService_WhoAmI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutomationService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_SetAllConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAllConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).SetAllConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutomationService_SetAllConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).SetAllConfig(ctx, req.(*SetAllConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutomationService_Install_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).Install(ctx, req.(*InstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_Preview_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PreviewRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomationServiceServer).Preview(m, &automationServicePreviewServer{ServerStream: stream})
}

type AutomationService_PreviewServer interface {
	Send(*PreviewStream) error
	grpc.ServerStream
}

type automationServicePreviewServer struct {
	grpc.ServerStream
}

func (x *automationServicePreviewServer) Send(m *PreviewStream) error {
	return x.ServerStream.SendMsg(m)
}

func _AutomationService_Refresh_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RefreshRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomationServiceServer).Refresh(m, &automationServiceRefreshServer{ServerStream: stream})
}

type AutomationService_RefreshServer interface {
	Send(*RefreshStream) error
	grpc.ServerStream
}

type automationServiceRefreshServer struct {
	grpc.ServerStream
}

func (x *automationServiceRefreshServer) Send(m *RefreshStream) error {
	return x.ServerStream.SendMsg(m)
}

func _AutomationService_Up_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomationServiceServer).Up(m, &automationServiceUpServer{ServerStream: stream})
}

type AutomationService_UpServer interface {
	Send(*UpStream) error
	grpc.ServerStream
}

type automationServiceUpServer struct {
	grpc.ServerStream
}

func (x *automationServiceUpServer) Send(m *UpStream) error {
	return x.ServerStream.SendMsg(m)
}

func _AutomationService_Destroy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DestroyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomationServiceServer).Destroy(m, &automationServiceDestroyServer{ServerStream: stream})
}

type AutomationService_DestroyServer interface {
	Send(*DestroyStream) error
	grpc.ServerStream
}

type automationServiceDestroyServer struct {
	grpc.ServerStream
}

func (x *automationServiceDestroyServer) Send(m *DestroyStream) error {
	return x.ServerStream.SendMsg(m)
}

// AutomationService_ServiceDesc is the grpc.ServiceDesc for AutomationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutomationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.AutomationService",
	HandlerType: (*AutomationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _AutomationService_WhoAmI_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _AutomationService_Info_Handler,
		},
		{
			MethodName: "SetAllConfig",
			Handler:    _AutomationService_SetAllConfig_Handler,
		},
		{
			MethodName: "Install",
			Handler:    _AutomationService_Install_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Preview",
			Handler:       _AutomationService_Preview_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Refresh",
			Handler:       _AutomationService_Refresh_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Up",
			Handler:       _AutomationService_Up_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Destroy",
			Handler:       _AutomationService_Destroy_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "agent.proto",
}
