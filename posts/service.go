package main

import (
	"blog-services/common/cache"
	"blog-services/common/proto"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type service struct {
	store PostsStore
	cache cache.Cache
}

func NewService(store PostsStore, cache cache.Cache) *service {
	return &service{store, cache}
}

func (s *service) GetPost(ctx context.Context, request *proto.GetPostRequest) (*proto.Post, error) {
	cacheKey := fmt.Sprintf("post:%d", request.PostId)

	cachedData, err := s.cache.Get(ctx, cacheKey)
	if err == nil {
		var post proto.Post
		if err := json.Unmarshal([]byte(cachedData), &post); err == nil {
			return &post, nil
		}
	}

	post, err := s.store.Get(ctx, request.PostId)
	if err != nil {
		return nil, err
	}

	postData, err := json.Marshal(post)
	if err == nil {
		s.cache.Set(ctx, cacheKey, string(postData), 15*time.Minute)
	}

	return post, nil
}

func (s *service) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	return s.store.Create(ctx, req)
}

func (s *service) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	post, err := s.store.Update(ctx, req)

	if err != nil {
		return nil, err
	}

	s.cache.Delete(ctx, fmt.Sprintf("post:%d", post.Id))

	return post, nil
}

func (s *service) GetPosts(ctx context.Context) ([]*proto.Post, error) {
	return s.store.GetList(ctx)
}

func (s *service) IncrementLikeCount(ctx context.Context, postId int64) (*proto.Post, error) {
	return s.store.UpdateLikesCount(ctx, postId)
}
