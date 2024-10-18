// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AutomationService_WhoAmI_FullMethodName       = "/agent.AutomationService/WhoAmI"
	AutomationService_SelectStack_FullMethodName  = "/agent.AutomationService/SelectStack"
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
	// WhoAmI returns detailed information about the currently logged-in Pulumi
	// identity.
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResult, error)
	// *
	// Info returns information about the given stack.
	SelectStack(ctx context.Context, in *SelectStackRequest, opts ...grpc.CallOption) (*SelectStackResult, error)
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
	// Preview performs a dry-run update to a stack, returning the expected
	// changes.
	Preview(ctx context.Context, in *PreviewRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PreviewStream], error)
	// *
	// Refresh updates the resources in a stack to match the current state of the
	// cloud provider.
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RefreshStream], error)
	// *
	// Up creates or updates the resources in a stack, returning the actual
	// changes.
	Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UpStream], error)
	// *
	// Destroy deletes all resources in a stack.
	Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DestroyStream], error)
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

func (c *automationServiceClient) SelectStack(ctx context.Context, in *SelectStackRequest, opts ...grpc.CallOption) (*SelectStackResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SelectStackResult)
	err := c.cc.Invoke(ctx, AutomationService_SelectStack_FullMethodName, in, out, cOpts...)
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

func (c *automationServiceClient) Preview(ctx context.Context, in *PreviewRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PreviewStream], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[0], AutomationService_Preview_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PreviewRequest, PreviewStream]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_PreviewClient = grpc.ServerStreamingClient[PreviewStream]

func (c *automationServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RefreshStream], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[1], AutomationService_Refresh_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RefreshRequest, RefreshStream]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_RefreshClient = grpc.ServerStreamingClient[RefreshStream]

func (c *automationServiceClient) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UpStream], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[2], AutomationService_Up_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UpRequest, UpStream]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_UpClient = grpc.ServerStreamingClient[UpStream]

func (c *automationServiceClient) Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DestroyStream], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AutomationService_ServiceDesc.Streams[3], AutomationService_Destroy_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DestroyRequest, DestroyStream]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_DestroyClient = grpc.ServerStreamingClient[DestroyStream]

// AutomationServiceServer is the server API for AutomationService service.
// All implementations must embed UnimplementedAutomationServiceServer
// for forward compatibility.
type AutomationServiceServer interface {
	// *
	// WhoAmI returns detailed information about the currently logged-in Pulumi
	// identity.
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResult, error)
	// *
	// Info returns information about the given stack.
	SelectStack(context.Context, *SelectStackRequest) (*SelectStackResult, error)
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
	// Preview performs a dry-run update to a stack, returning the expected
	// changes.
	Preview(*PreviewRequest, grpc.ServerStreamingServer[PreviewStream]) error
	// *
	// Refresh updates the resources in a stack to match the current state of the
	// cloud provider.
	Refresh(*RefreshRequest, grpc.ServerStreamingServer[RefreshStream]) error
	// *
	// Up creates or updates the resources in a stack, returning the actual
	// changes.
	Up(*UpRequest, grpc.ServerStreamingServer[UpStream]) error
	// *
	// Destroy deletes all resources in a stack.
	Destroy(*DestroyRequest, grpc.ServerStreamingServer[DestroyStream]) error
	mustEmbedUnimplementedAutomationServiceServer()
}

// UnimplementedAutomationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAutomationServiceServer struct{}

func (UnimplementedAutomationServiceServer) WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedAutomationServiceServer) SelectStack(context.Context, *SelectStackRequest) (*SelectStackResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectStack not implemented")
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
func (UnimplementedAutomationServiceServer) Preview(*PreviewRequest, grpc.ServerStreamingServer[PreviewStream]) error {
	return status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedAutomationServiceServer) Refresh(*RefreshRequest, grpc.ServerStreamingServer[RefreshStream]) error {
	return status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAutomationServiceServer) Up(*UpRequest, grpc.ServerStreamingServer[UpStream]) error {
	return status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (UnimplementedAutomationServiceServer) Destroy(*DestroyRequest, grpc.ServerStreamingServer[DestroyStream]) error {
	return status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedAutomationServiceServer) mustEmbedUnimplementedAutomationServiceServer() {}
func (UnimplementedAutomationServiceServer) testEmbeddedByValue()                           {}

// UnsafeAutomationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutomationServiceServer will
// result in compilation errors.
type UnsafeAutomationServiceServer interface {
	mustEmbedUnimplementedAutomationServiceServer()
}

func RegisterAutomationServiceServer(s grpc.ServiceRegistrar, srv AutomationServiceServer) {
	// If the following call pancis, it indicates UnimplementedAutomationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _AutomationService_SelectStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectStackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).SelectStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutomationService_SelectStack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).SelectStack(ctx, req.(*SelectStackRequest))
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
	return srv.(AutomationServiceServer).Preview(m, &grpc.GenericServerStream[PreviewRequest, PreviewStream]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_PreviewServer = grpc.ServerStreamingServer[PreviewStream]

func _AutomationService_Refresh_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RefreshRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomationServiceServer).Refresh(m, &grpc.GenericServerStream[RefreshRequest, RefreshStream]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_RefreshServer = grpc.ServerStreamingServer[RefreshStream]

func _AutomationService_Up_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomationServiceServer).Up(m, &grpc.GenericServerStream[UpRequest, UpStream]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_UpServer = grpc.ServerStreamingServer[UpStream]

func _AutomationService_Destroy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DestroyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomationServiceServer).Destroy(m, &grpc.GenericServerStream[DestroyRequest, DestroyStream]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AutomationService_DestroyServer = grpc.ServerStreamingServer[DestroyStream]

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
			MethodName: "SelectStack",
			Handler:    _AutomationService_SelectStack_Handler,
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
