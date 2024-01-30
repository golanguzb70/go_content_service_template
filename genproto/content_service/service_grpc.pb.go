// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: service.proto

package content_service

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
	PositionService_Create_FullMethodName = "/content_service.PositionService/Create"
	PositionService_Get_FullMethodName    = "/content_service.PositionService/Get"
	PositionService_Find_FullMethodName   = "/content_service.PositionService/Find"
	PositionService_Update_FullMethodName = "/content_service.PositionService/Update"
	PositionService_Delete_FullMethodName = "/content_service.PositionService/Delete"
)

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	Create(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Position, error)
	Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Positions, error)
	Update(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) Create(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, PositionService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, PositionService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Positions, error) {
	out := new(Positions)
	err := c.cc.Invoke(ctx, PositionService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Update(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := c.cc.Invoke(ctx, PositionService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, PositionService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations must embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	Create(context.Context, *Position) (*Position, error)
	Get(context.Context, *Id) (*Position, error)
	Find(context.Context, *GetListFilter) (*Positions, error)
	Update(context.Context, *Position) (*Position, error)
	Delete(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedPositionServiceServer()
}

// UnimplementedPositionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) Create(context.Context, *Position) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPositionServiceServer) Get(context.Context, *Id) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPositionServiceServer) Find(context.Context, *GetListFilter) (*Positions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedPositionServiceServer) Update(context.Context, *Position) (*Position, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPositionServiceServer) Delete(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPositionServiceServer) mustEmbedUnimplementedPositionServiceServer() {}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Create(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Find(ctx, req.(*GetListFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Update(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PositionService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_service.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PositionService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PositionService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _PositionService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PositionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PositionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

const (
	StaffService_Create_FullMethodName = "/content_service.StaffService/Create"
	StaffService_Get_FullMethodName    = "/content_service.StaffService/Get"
	StaffService_Find_FullMethodName   = "/content_service.StaffService/Find"
	StaffService_Update_FullMethodName = "/content_service.StaffService/Update"
	StaffService_Delete_FullMethodName = "/content_service.StaffService/Delete"
)

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	Create(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Staff, error)
	Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Staffs, error)
	Update(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) Create(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, StaffService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, StaffService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Staffs, error) {
	out := new(Staffs)
	err := c.cc.Invoke(ctx, StaffService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Update(ctx context.Context, in *Staff, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, StaffService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, StaffService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations must embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	Create(context.Context, *Staff) (*Staff, error)
	Get(context.Context, *Id) (*Staff, error)
	Find(context.Context, *GetListFilter) (*Staffs, error)
	Update(context.Context, *Staff) (*Staff, error)
	Delete(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedStaffServiceServer()
}

// UnimplementedStaffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) Create(context.Context, *Staff) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStaffServiceServer) Get(context.Context, *Id) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStaffServiceServer) Find(context.Context, *GetListFilter) (*Staffs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedStaffServiceServer) Update(context.Context, *Staff) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStaffServiceServer) Delete(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Staff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Create(ctx, req.(*Staff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Find(ctx, req.(*GetListFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Staff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Update(ctx, req.(*Staff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_service.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StaffService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StaffService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _StaffService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

const (
	TagService_Create_FullMethodName = "/content_service.TagService/Create"
	TagService_Get_FullMethodName    = "/content_service.TagService/Get"
	TagService_Find_FullMethodName   = "/content_service.TagService/Find"
	TagService_Update_FullMethodName = "/content_service.TagService/Update"
	TagService_Delete_FullMethodName = "/content_service.TagService/Delete"
)

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	Create(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Tag, error)
	Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Tags, error)
	Update(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) Create(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, TagService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, TagService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Tags, error) {
	out := new(Tags)
	err := c.cc.Invoke(ctx, TagService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) Update(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, TagService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TagService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	Create(context.Context, *Tag) (*Tag, error)
	Get(context.Context, *Id) (*Tag, error)
	Find(context.Context, *GetListFilter) (*Tags, error)
	Update(context.Context, *Tag) (*Tag, error)
	Delete(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) Create(context.Context, *Tag) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTagServiceServer) Get(context.Context, *Id) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTagServiceServer) Find(context.Context, *GetListFilter) (*Tags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedTagServiceServer) Update(context.Context, *Tag) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTagServiceServer) Delete(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).Create(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).Find(ctx, req.(*GetListFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).Update(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_service.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TagService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TagService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _TagService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TagService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TagService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

const (
	GenreService_Create_FullMethodName = "/content_service.GenreService/Create"
	GenreService_Get_FullMethodName    = "/content_service.GenreService/Get"
	GenreService_Find_FullMethodName   = "/content_service.GenreService/Find"
	GenreService_Update_FullMethodName = "/content_service.GenreService/Update"
	GenreService_Delete_FullMethodName = "/content_service.GenreService/Delete"
)

// GenreServiceClient is the client API for GenreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenreServiceClient interface {
	Create(ctx context.Context, in *Genre, opts ...grpc.CallOption) (*Genre, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Genre, error)
	Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Genres, error)
	Update(ctx context.Context, in *Genre, opts ...grpc.CallOption) (*Genre, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type genreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenreServiceClient(cc grpc.ClientConnInterface) GenreServiceClient {
	return &genreServiceClient{cc}
}

func (c *genreServiceClient) Create(ctx context.Context, in *Genre, opts ...grpc.CallOption) (*Genre, error) {
	out := new(Genre)
	err := c.cc.Invoke(ctx, GenreService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreServiceClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Genre, error) {
	out := new(Genre)
	err := c.cc.Invoke(ctx, GenreService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreServiceClient) Find(ctx context.Context, in *GetListFilter, opts ...grpc.CallOption) (*Genres, error) {
	out := new(Genres)
	err := c.cc.Invoke(ctx, GenreService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreServiceClient) Update(ctx context.Context, in *Genre, opts ...grpc.CallOption) (*Genre, error) {
	out := new(Genre)
	err := c.cc.Invoke(ctx, GenreService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genreServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, GenreService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenreServiceServer is the server API for GenreService service.
// All implementations must embed UnimplementedGenreServiceServer
// for forward compatibility
type GenreServiceServer interface {
	Create(context.Context, *Genre) (*Genre, error)
	Get(context.Context, *Id) (*Genre, error)
	Find(context.Context, *GetListFilter) (*Genres, error)
	Update(context.Context, *Genre) (*Genre, error)
	Delete(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedGenreServiceServer()
}

// UnimplementedGenreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGenreServiceServer struct {
}

func (UnimplementedGenreServiceServer) Create(context.Context, *Genre) (*Genre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGenreServiceServer) Get(context.Context, *Id) (*Genre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGenreServiceServer) Find(context.Context, *GetListFilter) (*Genres, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedGenreServiceServer) Update(context.Context, *Genre) (*Genre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGenreServiceServer) Delete(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGenreServiceServer) mustEmbedUnimplementedGenreServiceServer() {}

// UnsafeGenreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenreServiceServer will
// result in compilation errors.
type UnsafeGenreServiceServer interface {
	mustEmbedUnimplementedGenreServiceServer()
}

func RegisterGenreServiceServer(s grpc.ServiceRegistrar, srv GenreServiceServer) {
	s.RegisterService(&GenreService_ServiceDesc, srv)
}

func _GenreService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Genre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenreService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).Create(ctx, req.(*Genre))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenreService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenreService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).Find(ctx, req.(*GetListFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Genre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenreService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).Update(ctx, req.(*Genre))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenreService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenreService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// GenreService_ServiceDesc is the grpc.ServiceDesc for GenreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_service.GenreService",
	HandlerType: (*GenreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GenreService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GenreService_Get_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _GenreService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GenreService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GenreService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
