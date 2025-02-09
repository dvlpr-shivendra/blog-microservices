package gateway

import (
	"blog-services/common/discovery"
	"blog-services/common/proto"
	"context"
	"log"

	"go.uber.org/zap"
)

type gateway struct {
	registry    discovery.Registry
	postsClient proto.PostServiceClient
	logger      *zap.Logger
}

func NewGRPCGateway(registry discovery.Registry, logger *zap.Logger) *gateway {
	conn, err := discovery.ServiceConnection(context.Background(), "posts", registry)
	if err != nil {
		log.Fatalf("Failed to connect to posts service: %v", err)
	}

	return &gateway{
		registry:    registry,
		postsClient: proto.NewPostServiceClient(conn),
		logger:      logger,
	}
}

func (g *gateway) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "posts", g.registry)

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		return nil, err
	}

	defer conn.Close()

	c := proto.NewPostServiceClient(conn)

	return c.CreatePost(ctx, req)
}

func (g *gateway) GetPosts(ctx context.Context) ([]*proto.Post, error) {
	res, err := g.postsClient.GetPosts(ctx, &proto.Empty{})

	if err != nil {
		g.logger.Error("Failed to get posts", zap.Error(err))
		return nil, err
	}

	if res == nil {
		return []*proto.Post{}, nil
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

	defer conn.Close()

	c := proto.NewPostServiceClient(conn)

	return c.UpdatePost(ctx, req)
}

func (g *gateway) GetComments(ctx context.Context, postId int64) (*proto.GetCommentsResponse, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "comments", g.registry)

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		return nil, err
	}

	defer conn.Close()

	c := proto.NewCommentServiceClient(conn)

	return c.GetComments(ctx, &proto.GetCommentsRequest{PostId: postId})
}

func (g *gateway) CreateLike(ctx context.Context, postId int64) (*proto.CreateLikeResponse, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "likes", g.registry)

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		return nil, err
	}

	defer conn.Close()

	c := proto.NewLikeServiceClient(conn)

	return c.CreateLike(ctx, &proto.CreateLikeRequest{PostId: postId})
}
