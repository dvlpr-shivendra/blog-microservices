package gateway

import (
	"context"

	"blog-services/common/proto"
)

type PostsGateway interface {
	CreatePost(context.Context, *proto.CreatePostRequest) (*proto.Post, error)
	GetPost(ctx context.Context, id int64) (*proto.Post, error)
}