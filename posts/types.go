package main

import (
	"blog-services/common/proto"
	"context"
	"time"
)

type Post struct {
	Id         int64
	Title      string
	Body       string
	AuthorId   int64
	Published  bool
	LikesCount int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PostsService interface {
	CreatePost(context.Context, *proto.CreatePostRequest) (*proto.Post, error)
	UpdatePost(context.Context, *proto.UpdatePostRequest) (*proto.Post, error)
	GetPost(context.Context, *proto.GetPostRequest) (*proto.Post, error)
	GetPosts(context.Context) ([]*proto.Post, error)
	IncrementLikeCount(context.Context, int64) (*proto.Post, error)
}

type PostsStore interface {
	Create(context.Context, *proto.CreatePostRequest) (*proto.Post, error)
	Update(context.Context, *proto.UpdatePostRequest) (*proto.Post, error)
	Get(context.Context, int64) (*proto.Post, error)
	GetList(context.Context) ([]*proto.Post, error)
	UpdateLikesCount(ctx context.Context, postId int64) (*proto.Post, error)
}
