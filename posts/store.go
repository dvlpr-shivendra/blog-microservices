package main

import (
	"blog-services/common/proto"
	"context"
)

type store struct {
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context, request *proto.CreatePostRequest) (*proto.Post, error) {
	return &proto.Post{
		Id: 1,
		Title: request.Title,
		Body: request.Body,
		AuthorId: 1,
		Published: true,
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}, nil
}

func (s *store) Get(ctx context.Context, id int) (*proto.Post, error) {
	return &proto.Post{
		Id:        int64(id),
		Title:     "Test Post",
		Body:      "Test Body",
		AuthorId:  1,
		Published: true,
		CreatedAt: "2020-01-01",
		UpdatedAt: "2020-01-01",
	}, nil
}
