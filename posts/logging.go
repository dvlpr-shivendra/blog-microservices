package main

import (
	"context"
	"time"

	"blog-services/common/proto"

	"go.uber.org/zap"
)

type LoggingMiddleware struct {
	next PostsService
}

func NewLoggingMiddleware(next PostsService) PostsService {
	return &LoggingMiddleware{next}
}

func (s *LoggingMiddleware) CreatePost(ctx context.Context, request *proto.CreatePostRequest) (*proto.Post, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("CreatePost", zap.Duration("took", time.Since(start)))
	}()
	return s.next.CreatePost(ctx, request)
}

func (s *LoggingMiddleware) GetPost(ctx context.Context, request *proto.GetPostRequest) (*proto.Post, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("GetPost", zap.Duration("took", time.Since(start)))
	}()
	return s.next.GetPost(ctx, request)
}

func (s *LoggingMiddleware) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("UpdatePost", zap.Duration("took", time.Since(start)))
	}()
	return s.next.UpdatePost(ctx, req)
}

func (s *LoggingMiddleware) GetPosts(ctx context.Context) ([]*proto.Post, error) {
	start := time.Now()
	defer func() {
		zap.L().Info("GetPosts", zap.Duration("took", time.Since(start)))
	}()
	return s.next.GetPosts(ctx)
}
