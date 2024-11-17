package gateway

import (
	"blog-services/common/discovery"
	"blog-services/common/proto"
	"context"
	"log"
)

type gateway struct {
	registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
	return &gateway{registry}
}

func (g *gateway) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "orders", g.registry)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	c := proto.NewPostServiceClient(conn)

	return c.CreatePost(ctx, &proto.CreatePostRequest{
		Title: req.Title,
		Body:  req.Body,
	})
}

func (g *gateway) GetPost(ctx context.Context, id int64) (*proto.Post, error) {
	return nil, nil
}
