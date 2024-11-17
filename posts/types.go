package main

import (
	"blog-services/common/proto"
	"context"
	"time"
)

type Post struct {
	Id        int
	Title     string
	Body      string
	AuthorID  int
	Published bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostsService interface {
	CreatePost(context.Context, *proto.CreatePostRequest) (*proto.Post, error)
	GetPost(context.Context, *proto.GetPostRequest) (*proto.Post, error)
}

type PostsStore interface {
	Create(context.Context, *proto.CreatePostRequest) (*proto.Post, error)
	Get(ctx context.Context, id int) (*proto.Post, error)
}
