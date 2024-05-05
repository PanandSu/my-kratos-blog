// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: api/comment/comment.proto

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
	Comment_AddComment_FullMethodName            = "/api.Comment/AddComment"
	Comment_AddReward_FullMethodName             = "/api.Comment/AddReward"
	Comment_ExtractParentComments_FullMethodName = "/api.Comment/ExtractParentComments"
)

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	AddComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentReply, error)
	AddReward(ctx context.Context, in *CreateRewardRequest, opts ...grpc.CallOption) (*CreateRewardReply, error)
	ExtractParentComments(ctx context.Context, in *ExtractParentCommentsRequest, opts ...grpc.CallOption) (*ExtractParentCommentsReply, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) AddComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentReply, error) {
	out := new(CreateCommentReply)
	err := c.cc.Invoke(ctx, Comment_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) AddReward(ctx context.Context, in *CreateRewardRequest, opts ...grpc.CallOption) (*CreateRewardReply, error) {
	out := new(CreateRewardReply)
	err := c.cc.Invoke(ctx, Comment_AddReward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) ExtractParentComments(ctx context.Context, in *ExtractParentCommentsRequest, opts ...grpc.CallOption) (*ExtractParentCommentsReply, error) {
	out := new(ExtractParentCommentsReply)
	err := c.cc.Invoke(ctx, Comment_ExtractParentComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	AddComment(context.Context, *CreateCommentRequest) (*CreateCommentReply, error)
	AddReward(context.Context, *CreateRewardRequest) (*CreateRewardReply, error)
	ExtractParentComments(context.Context, *ExtractParentCommentsRequest) (*ExtractParentCommentsReply, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) AddComment(context.Context, *CreateCommentRequest) (*CreateCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCommentServer) AddReward(context.Context, *CreateRewardRequest) (*CreateRewardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReward not implemented")
}
func (UnimplementedCommentServer) ExtractParentComments(context.Context, *ExtractParentCommentsRequest) (*ExtractParentCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractParentComments not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).AddComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_AddReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).AddReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_AddReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).AddReward(ctx, req.(*CreateRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_ExtractParentComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractParentCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).ExtractParentComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_ExtractParentComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).ExtractParentComments(ctx, req.(*ExtractParentCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _Comment_AddComment_Handler,
		},
		{
			MethodName: "AddReward",
			Handler:    _Comment_AddReward_Handler,
		},
		{
			MethodName: "ExtractParentComments",
			Handler:    _Comment_ExtractParentComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/comment/comment.proto",
}
