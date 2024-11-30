// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/blog.proto

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
	PostService_CreatePost_FullMethodName = "/proto.PostService/CreatePost"
	PostService_GetPost_FullMethodName    = "/proto.PostService/GetPost"
	PostService_GetPosts_FullMethodName   = "/proto.PostService/GetPosts"
	PostService_UpdatePost_FullMethodName = "/proto.PostService/UpdatePost"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
	GetPosts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPostsResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_GetPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPosts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, PostService_GetPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_UpdatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility.
type PostServiceServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	GetPosts(context.Context, *Empty) (*GetPostsResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*Post, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostServiceServer struct{}

func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) GetPost(context.Context, *GetPostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) GetPosts(context.Context, *Empty) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedPostServiceServer) UpdatePost(context.Context, *UpdatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}
func (UnimplementedPostServiceServer) testEmbeddedByValue()                     {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	// If the following call pancis, it indicates UnimplementedPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPosts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "GetPosts",
			Handler:    _PostService_GetPosts_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blog.proto",
}

const (
	CommentService_CreateComment_FullMethodName = "/proto.CommentService/CreateComment"
	CommentService_GetComments_FullMethodName   = "/proto.CommentService/GetComments"
	CommentService_UpdateComment_FullMethodName = "/proto.CommentService/UpdateComment"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*Comment, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*Comment, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, CommentService_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_GetComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, CommentService_UpdateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility.
type CommentServiceServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*Comment, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*Comment, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentServiceServer struct{}

func (UnimplementedCommentServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedCommentServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}
func (UnimplementedCommentServiceServer) testEmbeddedByValue()                        {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _CommentService_GetComments_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _CommentService_UpdateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blog.proto",
}
