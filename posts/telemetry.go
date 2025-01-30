package main

import (
	"context"

	"fmt"

	"blog-services/common/proto"

	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
	next PostsService
}

func NewTelemetryMiddleware(next PostsService) PostsService {
	return &TelemetryMiddleware{next}
}

func (s *TelemetryMiddleware) CreatePost(ctx context.Context, request *proto.CreatePostRequest) (*proto.Post, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("CreatePost: %v", request))
	return s.next.CreatePost(ctx, request)
}

func (s *TelemetryMiddleware) GetPost(ctx context.Context, request *proto.GetPostRequest) (*proto.Post, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("GetPost: %v", request))
	return s.next.GetPost(ctx, request)
}

func (s *TelemetryMiddleware) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("UpdatePost: %v", req))
	return s.next.UpdatePost(ctx, req)
}

func (s *TelemetryMiddleware) GetPosts(ctx context.Context) ([]*proto.Post, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("GetPosts invoked")
	return s.next.GetPosts(ctx)
}

func (s *TelemetryMiddleware) IncrementLikeCount(ctx context.Context, postId string) ([]*proto.Post, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("GetPosts invoked")
	return s.next.GetPosts(ctx)
}
