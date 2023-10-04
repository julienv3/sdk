// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: advisory.platform.proto

package v1

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
	SecurityAdvisory_ListDocuments_FullMethodName = "/chainguard.platform.advisory.SecurityAdvisory/ListDocuments"
)

// SecurityAdvisoryClient is the client API for SecurityAdvisory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityAdvisoryClient interface {
	ListDocuments(ctx context.Context, in *DocumentFilter, opts ...grpc.CallOption) (*DocumentList, error)
}

type securityAdvisoryClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityAdvisoryClient(cc grpc.ClientConnInterface) SecurityAdvisoryClient {
	return &securityAdvisoryClient{cc}
}

func (c *securityAdvisoryClient) ListDocuments(ctx context.Context, in *DocumentFilter, opts ...grpc.CallOption) (*DocumentList, error) {
	out := new(DocumentList)
	err := c.cc.Invoke(ctx, SecurityAdvisory_ListDocuments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityAdvisoryServer is the server API for SecurityAdvisory service.
// All implementations must embed UnimplementedSecurityAdvisoryServer
// for forward compatibility
type SecurityAdvisoryServer interface {
	ListDocuments(context.Context, *DocumentFilter) (*DocumentList, error)
	mustEmbedUnimplementedSecurityAdvisoryServer()
}

// UnimplementedSecurityAdvisoryServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityAdvisoryServer struct {
}

func (UnimplementedSecurityAdvisoryServer) ListDocuments(context.Context, *DocumentFilter) (*DocumentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}
func (UnimplementedSecurityAdvisoryServer) mustEmbedUnimplementedSecurityAdvisoryServer() {}

// UnsafeSecurityAdvisoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityAdvisoryServer will
// result in compilation errors.
type UnsafeSecurityAdvisoryServer interface {
	mustEmbedUnimplementedSecurityAdvisoryServer()
}

func RegisterSecurityAdvisoryServer(s grpc.ServiceRegistrar, srv SecurityAdvisoryServer) {
	s.RegisterService(&SecurityAdvisory_ServiceDesc, srv)
}

func _SecurityAdvisory_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityAdvisoryServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityAdvisory_ListDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityAdvisoryServer).ListDocuments(ctx, req.(*DocumentFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityAdvisory_ServiceDesc is the grpc.ServiceDesc for SecurityAdvisory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityAdvisory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.advisory.SecurityAdvisory",
	HandlerType: (*SecurityAdvisoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDocuments",
			Handler:    _SecurityAdvisory_ListDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advisory.platform.proto",
}
