// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: meridian.proto

package api

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

const (
	Meridian_AddNamespace_FullMethodName          = "/proto.Meridian/AddNamespace"
	Meridian_RemoveNamespace_FullMethodName       = "/proto.Meridian/RemoveNamespace"
	Meridian_AddApp_FullMethodName                = "/proto.Meridian/AddApp"
	Meridian_RemoveApp_FullMethodName             = "/proto.Meridian/RemoveApp"
	Meridian_GetNamespace_FullMethodName          = "/proto.Meridian/GetNamespace"
	Meridian_GetNamespaceHierarchy_FullMethodName = "/proto.Meridian/GetNamespaceHierarchy"
	Meridian_SetNamespaceResources_FullMethodName = "/proto.Meridian/SetNamespaceResources"
	Meridian_SetAppResources_FullMethodName       = "/proto.Meridian/SetAppResources"
	Meridian_SendMessage_FullMethodName           = "/proto.Meridian/SendMessage"
	Meridian_BorrowResources_FullMethodName       = "/proto.Meridian/BorrowResources"
)

// MeridianClient is the client API for Meridian service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeridianClient interface {
	AddNamespace(ctx context.Context, in *AddNamespaceReq, opts ...grpc.CallOption) (*AddNamespaceResp, error)
	RemoveNamespace(ctx context.Context, in *RemoveNamespaceReq, opts ...grpc.CallOption) (*RemoveNamespaceResp, error)
	AddApp(ctx context.Context, in *AddAppReq, opts ...grpc.CallOption) (*AddAppResp, error)
	RemoveApp(ctx context.Context, in *RemoveAppReq, opts ...grpc.CallOption) (*RemoveAppResp, error)
	GetNamespace(ctx context.Context, in *GetNamespaceReq, opts ...grpc.CallOption) (*GetNamespaceResp, error)
	GetNamespaceHierarchy(ctx context.Context, in *GetNamespaceHierarchyReq, opts ...grpc.CallOption) (*GetNamespaceHierarchyResp, error)
	SetNamespaceResources(ctx context.Context, in *SetNamespaceResourcesReq, opts ...grpc.CallOption) (*SetNamespaceResourcesResp, error)
	SetAppResources(ctx context.Context, in *SetAppResourcesReq, opts ...grpc.CallOption) (*SetAppResourcesResp, error)
	SendMessage(ctx context.Context, in *SendMess, opts ...grpc.CallOption) (*SendMessResp, error)
	BorrowResources(ctx context.Context, in *BorrowResourcesReq, opts ...grpc.CallOption) (*BorrowResourcesResp, error)
}

type meridianClient struct {
	cc grpc.ClientConnInterface
}

func NewMeridianClient(cc grpc.ClientConnInterface) MeridianClient {
	return &meridianClient{cc}
}

func (c *meridianClient) AddNamespace(ctx context.Context, in *AddNamespaceReq, opts ...grpc.CallOption) (*AddNamespaceResp, error) {
	out := new(AddNamespaceResp)
	err := c.cc.Invoke(ctx, Meridian_AddNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) RemoveNamespace(ctx context.Context, in *RemoveNamespaceReq, opts ...grpc.CallOption) (*RemoveNamespaceResp, error) {
	out := new(RemoveNamespaceResp)
	err := c.cc.Invoke(ctx, Meridian_RemoveNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) AddApp(ctx context.Context, in *AddAppReq, opts ...grpc.CallOption) (*AddAppResp, error) {
	out := new(AddAppResp)
	err := c.cc.Invoke(ctx, Meridian_AddApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) RemoveApp(ctx context.Context, in *RemoveAppReq, opts ...grpc.CallOption) (*RemoveAppResp, error) {
	out := new(RemoveAppResp)
	err := c.cc.Invoke(ctx, Meridian_RemoveApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) GetNamespace(ctx context.Context, in *GetNamespaceReq, opts ...grpc.CallOption) (*GetNamespaceResp, error) {
	out := new(GetNamespaceResp)
	err := c.cc.Invoke(ctx, Meridian_GetNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) GetNamespaceHierarchy(ctx context.Context, in *GetNamespaceHierarchyReq, opts ...grpc.CallOption) (*GetNamespaceHierarchyResp, error) {
	out := new(GetNamespaceHierarchyResp)
	err := c.cc.Invoke(ctx, Meridian_GetNamespaceHierarchy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) SetNamespaceResources(ctx context.Context, in *SetNamespaceResourcesReq, opts ...grpc.CallOption) (*SetNamespaceResourcesResp, error) {
	out := new(SetNamespaceResourcesResp)
	err := c.cc.Invoke(ctx, Meridian_SetNamespaceResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) SetAppResources(ctx context.Context, in *SetAppResourcesReq, opts ...grpc.CallOption) (*SetAppResourcesResp, error) {
	out := new(SetAppResourcesResp)
	err := c.cc.Invoke(ctx, Meridian_SetAppResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) SendMessage(ctx context.Context, in *SendMess, opts ...grpc.CallOption) (*SendMessResp, error) {
	out := new(SendMessResp)
	err := c.cc.Invoke(ctx, Meridian_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meridianClient) BorrowResources(ctx context.Context, in *BorrowResourcesReq, opts ...grpc.CallOption) (*BorrowResourcesResp, error) {
	out := new(BorrowResourcesResp)
	err := c.cc.Invoke(ctx, Meridian_BorrowResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeridianServer is the server API for Meridian service.
// All implementations must embed UnimplementedMeridianServer
// for forward compatibility
type MeridianServer interface {
	AddNamespace(context.Context, *AddNamespaceReq) (*AddNamespaceResp, error)
	RemoveNamespace(context.Context, *RemoveNamespaceReq) (*RemoveNamespaceResp, error)
	AddApp(context.Context, *AddAppReq) (*AddAppResp, error)
	RemoveApp(context.Context, *RemoveAppReq) (*RemoveAppResp, error)
	GetNamespace(context.Context, *GetNamespaceReq) (*GetNamespaceResp, error)
	GetNamespaceHierarchy(context.Context, *GetNamespaceHierarchyReq) (*GetNamespaceHierarchyResp, error)
	SetNamespaceResources(context.Context, *SetNamespaceResourcesReq) (*SetNamespaceResourcesResp, error)
	SetAppResources(context.Context, *SetAppResourcesReq) (*SetAppResourcesResp, error)
	SendMessage(context.Context, *SendMess) (*SendMessResp, error)
	BorrowResources(context.Context, *BorrowResourcesReq) (*BorrowResourcesResp, error)
	mustEmbedUnimplementedMeridianServer()
}

// UnimplementedMeridianServer must be embedded to have forward compatible implementations.
type UnimplementedMeridianServer struct {
}

func (UnimplementedMeridianServer) AddNamespace(context.Context, *AddNamespaceReq) (*AddNamespaceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNamespace not implemented")
}
func (UnimplementedMeridianServer) RemoveNamespace(context.Context, *RemoveNamespaceReq) (*RemoveNamespaceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNamespace not implemented")
}
func (UnimplementedMeridianServer) AddApp(context.Context, *AddAppReq) (*AddAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddApp not implemented")
}
func (UnimplementedMeridianServer) RemoveApp(context.Context, *RemoveAppReq) (*RemoveAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveApp not implemented")
}
func (UnimplementedMeridianServer) GetNamespace(context.Context, *GetNamespaceReq) (*GetNamespaceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespace not implemented")
}
func (UnimplementedMeridianServer) GetNamespaceHierarchy(context.Context, *GetNamespaceHierarchyReq) (*GetNamespaceHierarchyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespaceHierarchy not implemented")
}
func (UnimplementedMeridianServer) SetNamespaceResources(context.Context, *SetNamespaceResourcesReq) (*SetNamespaceResourcesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNamespaceResources not implemented")
}
func (UnimplementedMeridianServer) SetAppResources(context.Context, *SetAppResourcesReq) (*SetAppResourcesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAppResources not implemented")
}
func (UnimplementedMeridianServer) SendMessage(context.Context, *SendMess) (*SendMessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMeridianServer) BorrowResources(context.Context, *BorrowResourcesReq) (*BorrowResourcesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowResources not implemented")
}
func (UnimplementedMeridianServer) mustEmbedUnimplementedMeridianServer() {}

// UnsafeMeridianServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeridianServer will
// result in compilation errors.
type UnsafeMeridianServer interface {
	mustEmbedUnimplementedMeridianServer()
}

func RegisterMeridianServer(s grpc.ServiceRegistrar, srv MeridianServer) {
	s.RegisterService(&Meridian_ServiceDesc, srv)
}

func _Meridian_AddNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNamespaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).AddNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_AddNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).AddNamespace(ctx, req.(*AddNamespaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_RemoveNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNamespaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).RemoveNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_RemoveNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).RemoveNamespace(ctx, req.(*RemoveNamespaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_AddApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).AddApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_AddApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).AddApp(ctx, req.(*AddAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_RemoveApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).RemoveApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_RemoveApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).RemoveApp(ctx, req.(*RemoveAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_GetNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).GetNamespace(ctx, req.(*GetNamespaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_GetNamespaceHierarchy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceHierarchyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).GetNamespaceHierarchy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_GetNamespaceHierarchy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).GetNamespaceHierarchy(ctx, req.(*GetNamespaceHierarchyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_SetNamespaceResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNamespaceResourcesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).SetNamespaceResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_SetNamespaceResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).SetNamespaceResources(ctx, req.(*SetNamespaceResourcesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_SetAppResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAppResourcesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).SetAppResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_SetAppResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).SetAppResources(ctx, req.(*SetAppResourcesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).SendMessage(ctx, req.(*SendMess))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meridian_BorrowResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowResourcesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeridianServer).BorrowResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meridian_BorrowResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeridianServer).BorrowResources(ctx, req.(*BorrowResourcesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Meridian_ServiceDesc is the grpc.ServiceDesc for Meridian service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Meridian_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Meridian",
	HandlerType: (*MeridianServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNamespace",
			Handler:    _Meridian_AddNamespace_Handler,
		},
		{
			MethodName: "RemoveNamespace",
			Handler:    _Meridian_RemoveNamespace_Handler,
		},
		{
			MethodName: "AddApp",
			Handler:    _Meridian_AddApp_Handler,
		},
		{
			MethodName: "RemoveApp",
			Handler:    _Meridian_RemoveApp_Handler,
		},
		{
			MethodName: "GetNamespace",
			Handler:    _Meridian_GetNamespace_Handler,
		},
		{
			MethodName: "GetNamespaceHierarchy",
			Handler:    _Meridian_GetNamespaceHierarchy_Handler,
		},
		{
			MethodName: "SetNamespaceResources",
			Handler:    _Meridian_SetNamespaceResources_Handler,
		},
		{
			MethodName: "SetAppResources",
			Handler:    _Meridian_SetAppResources_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Meridian_SendMessage_Handler,
		},
		{
			MethodName: "BorrowResources",
			Handler:    _Meridian_BorrowResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meridian.proto",
}
