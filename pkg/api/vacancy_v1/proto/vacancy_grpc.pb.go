// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/vacancy.proto

package vacancy

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

// VacancyClient is the client API for Vacancy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacancyClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*VacancyResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*VacancyResponse, error)
}

type vacancyClient struct {
	cc grpc.ClientConnInterface
}

func NewVacancyClient(cc grpc.ClientConnInterface) VacancyClient {
	return &vacancyClient{cc}
}

func (c *vacancyClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*VacancyResponse, error) {
	out := new(VacancyResponse)
	err := c.cc.Invoke(ctx, "/vacancy_v1.Vacancy/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*VacancyResponse, error) {
	out := new(VacancyResponse)
	err := c.cc.Invoke(ctx, "/vacancy_v1.Vacancy/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacancyServer is the server API for Vacancy service.
// All implementations must embed UnimplementedVacancyServer
// for forward compatibility
type VacancyServer interface {
	Create(context.Context, *CreateRequest) (*VacancyResponse, error)
	GetById(context.Context, *GetByIdRequest) (*VacancyResponse, error)
	mustEmbedUnimplementedVacancyServer()
}

// UnimplementedVacancyServer must be embedded to have forward compatible implementations.
type UnimplementedVacancyServer struct {
}

func (UnimplementedVacancyServer) Create(context.Context, *CreateRequest) (*VacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVacancyServer) GetById(context.Context, *GetByIdRequest) (*VacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedVacancyServer) mustEmbedUnimplementedVacancyServer() {}

// UnsafeVacancyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacancyServer will
// result in compilation errors.
type UnsafeVacancyServer interface {
	mustEmbedUnimplementedVacancyServer()
}

func RegisterVacancyServer(s grpc.ServiceRegistrar, srv VacancyServer) {
	s.RegisterService(&Vacancy_ServiceDesc, srv)
}

func _Vacancy_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacancy_v1.Vacancy/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vacancy_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vacancy_v1.Vacancy/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vacancy_ServiceDesc is the grpc.ServiceDesc for Vacancy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vacancy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vacancy_v1.Vacancy",
	HandlerType: (*VacancyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Vacancy_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _Vacancy_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vacancy.proto",
}