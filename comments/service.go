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

func (s *service) GetComments(context.Context) ([]*proto.Comment, error) {
	return s.store.GetAll(1)
}
