package main

import (
	"blog-services/common/proto"
	"context"
)

type service struct {
	store PostsStore
}

func NewService(store PostsStore) *service {
	return &service{store}
}

func (s *service) GetPost(ctx context.Context, request *proto.GetPostRequest) (*proto.Post, error) {
	return s.store.Get(ctx, int(request.PostId))
}

func (s *service) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	return s.store.Create(ctx, req)
}