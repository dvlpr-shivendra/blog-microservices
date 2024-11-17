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

func (h *grpcHandler) GetOrder(ctx context.Context, req *proto.GetPostRequest) (*proto.Post, error) {
	return h.service.GetPost(ctx, req)
}

func (h *grpcHandler) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	return h.service.CreatePost(ctx, req)
}
