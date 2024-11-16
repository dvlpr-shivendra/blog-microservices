package gateway

import (
	"blog-services/common/proto"
	"context"
)

type gateway struct {
}

func NewGRPCGateway() *gateway {
	return &gateway{}
}

func (g *gateway) CreatePost(ctx context.Context, p *proto.CreatePostRequest) (*proto.Post, error) {
	return nil, nil
}

func (g *gateway) GetPost(ctx context.Context, id int64) (*proto.Post, error) {
	return nil, nil
}
