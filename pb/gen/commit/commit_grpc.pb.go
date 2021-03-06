// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package commit

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

// CommitServiceClient is the client API for CommitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommitServiceClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	Get(ctx context.Context, in *CommitID, opts ...grpc.CallOption) (*Commit, error)
}

type commitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommitServiceClient(cc grpc.ClientConnInterface) CommitServiceClient {
	return &commitServiceClient{cc}
}

func (c *commitServiceClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, "/commit.CommitService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commitServiceClient) Get(ctx context.Context, in *CommitID, opts ...grpc.CallOption) (*Commit, error) {
	out := new(Commit)
	err := c.cc.Invoke(ctx, "/commit.CommitService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommitServiceServer is the server API for CommitService service.
// All implementations must embed UnimplementedCommitServiceServer
// for forward compatibility
type CommitServiceServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	Get(context.Context, *CommitID) (*Commit, error)
	mustEmbedUnimplementedCommitServiceServer()
}

// UnimplementedCommitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommitServiceServer struct {
}

func (UnimplementedCommitServiceServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCommitServiceServer) Get(context.Context, *CommitID) (*Commit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCommitServiceServer) mustEmbedUnimplementedCommitServiceServer() {}

// UnsafeCommitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommitServiceServer will
// result in compilation errors.
type UnsafeCommitServiceServer interface {
	mustEmbedUnimplementedCommitServiceServer()
}

func RegisterCommitServiceServer(s grpc.ServiceRegistrar, srv CommitServiceServer) {
	s.RegisterService(&CommitService_ServiceDesc, srv)
}

func _CommitService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commit.CommitService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitServiceServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommitService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commit.CommitService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitServiceServer).Get(ctx, req.(*CommitID))
	}
	return interceptor(ctx, in, info, handler)
}

// CommitService_ServiceDesc is the grpc.ServiceDesc for CommitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commit.CommitService",
	HandlerType: (*CommitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CommitService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CommitService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commit.proto",
}
