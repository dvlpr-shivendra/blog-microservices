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
	conn, err := discovery.ServiceConnection(context.Background(), "posts", g.registry)

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		return nil, err
	}

	c := proto.NewPostServiceClient(conn)

	return c.CreatePost(ctx, req)
}

func (g *gateway) GetPosts(ctx context.Context) ([]*proto.Post, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "posts", g.registry)

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		return nil, err
	}

	c := proto.NewPostServiceClient(conn)

	res, err := c.GetPosts(ctx, &proto.Empty{})

	if err != nil {
		log.Printf("Could not get posts: %v", err)
		return nil, err
	}

	if res == nil {
		return make([]*proto.Post, 0), nil
	}

	return res.Posts, nil
}

func (g *gateway) GetPost(ctx context.Context, id int64) (*proto.Post, error) {
	return nil, nil
}

func (g *gateway) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "posts", g.registry)

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		return nil, err
	}

	c := proto.NewPostServiceClient(conn)

	return c.UpdatePost(ctx, req)
}

func (g *gateway) GetComments(ctx context.Context, postId int64) (*proto.GetCommentsResponse, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "comments", g.registry)

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		return nil, err
	}

	c := proto.NewCommentServiceClient(conn)

	return c.GetComments(ctx, &proto.GetCommentsRequest{PostId: postId})
}
