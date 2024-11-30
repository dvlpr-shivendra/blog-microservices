package main

import (
	"blog-services/common/proto"
	"context"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	proto.UnimplementedCommentServiceServer

	service CommentsService
}

func NewGRPCHandler(grpcServer *grpc.Server, service CommentsService) {
	handler := &grpcHandler{
		service: service,
	}
	proto.RegisterCommentServiceServer(grpcServer, handler)
}

func (h *grpcHandler) GetComments(ctx context.Context, req *proto.GetCommentsRequest) (*proto.GetCommentsResponse, error) {
	comments, err := h.service.GetComments(ctx, req.PostId)

	if err != nil {
		return nil, err
	}

	return &proto.GetCommentsResponse{Comments: comments}, nil
}
