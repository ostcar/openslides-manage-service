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

// ManageClient is the client API for Manage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageClient interface {
	CheckServer(ctx context.Context, in *CheckServerRequest, opts ...grpc.CallOption) (*CheckServerResponse, error)
	InitialData(ctx context.Context, in *InitialDataRequest, opts ...grpc.CallOption) (*InitialDataResponse, error)
	Migrations(ctx context.Context, in *MigrationsRequest, opts ...grpc.CallOption) (*MigrationsResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type manageClient struct {
	cc grpc.ClientConnInterface
}

func NewManageClient(cc grpc.ClientConnInterface) ManageClient {
	return &manageClient{cc}
}

func (c *manageClient) CheckServer(ctx context.Context, in *CheckServerRequest, opts ...grpc.CallOption) (*CheckServerResponse, error) {
	out := new(CheckServerResponse)
	err := c.cc.Invoke(ctx, "/Manage/CheckServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) InitialData(ctx context.Context, in *InitialDataRequest, opts ...grpc.CallOption) (*InitialDataResponse, error) {
	out := new(InitialDataResponse)
	err := c.cc.Invoke(ctx, "/Manage/InitialData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) Migrations(ctx context.Context, in *MigrationsRequest, opts ...grpc.CallOption) (*MigrationsResponse, error) {
	out := new(MigrationsResponse)
	err := c.cc.Invoke(ctx, "/Manage/Migrations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/Manage/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error) {
	out := new(SetPasswordResponse)
	err := c.cc.Invoke(ctx, "/Manage/SetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/Manage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/Manage/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/Manage/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/Manage/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManageServer is the server API for Manage service.
// All implementations should embed UnimplementedManageServer
// for forward compatibility
type ManageServer interface {
	CheckServer(context.Context, *CheckServerRequest) (*CheckServerResponse, error)
	InitialData(context.Context, *InitialDataRequest) (*InitialDataResponse, error)
	Migrations(context.Context, *MigrationsRequest) (*MigrationsResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

// UnimplementedManageServer should be embedded to have forward compatible implementations.
type UnimplementedManageServer struct {
}

func (UnimplementedManageServer) CheckServer(context.Context, *CheckServerRequest) (*CheckServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckServer not implemented")
}
func (UnimplementedManageServer) InitialData(context.Context, *InitialDataRequest) (*InitialDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialData not implemented")
}
func (UnimplementedManageServer) Migrations(context.Context, *MigrationsRequest) (*MigrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Migrations not implemented")
}
func (UnimplementedManageServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedManageServer) SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedManageServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedManageServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedManageServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedManageServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

// UnsafeManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageServer will
// result in compilation errors.
type UnsafeManageServer interface {
	mustEmbedUnimplementedManageServer()
}

func RegisterManageServer(s grpc.ServiceRegistrar, srv ManageServer) {
	s.RegisterService(&Manage_ServiceDesc, srv)
}

func _Manage_CheckServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).CheckServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/CheckServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).CheckServer(ctx, req.(*CheckServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_InitialData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitialDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).InitialData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/InitialData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).InitialData(ctx, req.(*InitialDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_Migrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).Migrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/Migrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).Migrations(ctx, req.(*MigrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manage_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manage/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manage_ServiceDesc is the grpc.ServiceDesc for Manage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Manage",
	HandlerType: (*ManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckServer",
			Handler:    _Manage_CheckServer_Handler,
		},
		{
			MethodName: "InitialData",
			Handler:    _Manage_InitialData_Handler,
		},
		{
			MethodName: "Migrations",
			Handler:    _Manage_Migrations_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Manage_CreateUser_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _Manage_SetPassword_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Manage_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Manage_Set_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Manage_Version_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Manage_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/manage.proto",
}
