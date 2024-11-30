package main

import (
	"blog-services/common/proto"
	"context"
)

type service struct {
	store CommentsStore
}

func NewService(store CommentsStore) *service {
	return &service{store}
}

func (s *service) GetComments(ctx context.Context, postId int64) ([]*proto.Comment, error) {
	return s.store.GetAll(ctx, postId)
}
