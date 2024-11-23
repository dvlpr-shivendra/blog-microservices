package main

import (
	"blog-services/common/proto"
	"context"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	proto.UnimplementedPostServiceServer

	service PostsService
}

func NewGRPCHandler(grpcServer *grpc.Server, service PostsService) {
	handler := &grpcHandler{
		service: service,
	}
	proto.RegisterPostServiceServer(grpcServer, handler)
}

func (h *grpcHandler) GetPost(ctx context.Context, req *proto.GetPostRequest) (*proto.Post, error) {
	return h.service.GetPost(ctx, req)
}

func (h *grpcHandler) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	return h.service.CreatePost(ctx, req)
}

func (h *grpcHandler) GetPosts(ctx context.Context, req *proto.Empty) (*proto.GetPostsResponse, error) {
	posts, err := h.service.GetPosts(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.GetPostsResponse{Posts: posts}, nil
}
