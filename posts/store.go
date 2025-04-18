package main

import (
	"blog-services/common/proto"
	"context"
	"database/sql"
	"log"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db}
}

func (s *store) Create(ctx context.Context, r *proto.CreatePostRequest) (*proto.Post, error) {
	post := &proto.Post{}

	query := `
		INSERT INTO posts (title, body, author_id, published)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, body, author_id, published, created_at, updated_at
	`

	err := s.db.QueryRowContext(ctx, query, r.Title, r.Body, 1, true).Scan(
		&post.Id,
		&post.Title,
		&post.Body,
		&post.AuthorId,
		&post.Published,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return post, nil
}

func (s *store) Update(ctx context.Context, r *proto.UpdatePostRequest) (*proto.Post, error) {
	post := &proto.Post{}
	query := `
	UPDATE posts
	SET 
		title = $2,
		body = $3,
		updated_at = NOW()
	WHERE id = $1
	RETURNING id, title, body, author_id, published, created_at, updated_at;
	`
	err := s.db.QueryRowContext(ctx, query, r.Id, r.Title, r.Body).Scan(
		&post.Id,
		&post.Title,
		&post.Body,
		&post.AuthorId,
		&post.Published,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *store) Get(ctx context.Context, id int64) (*proto.Post, error) {
	post := &proto.Post{}
	err := s.db.QueryRowContext(ctx, "SELECT id, title, body, author_id, published, created_at, updated_at FROM posts WHERE id = $1", id).Scan(
		&post.Id,
		&post.Title,
		&post.Body,
		&post.AuthorId,
		&post.Published,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *store) GetList(ctx context.Context) ([]*proto.Post, error) {
	var posts []*proto.Post
	rows, err := s.db.QueryContext(ctx, "SELECT id, title, body, author_id, published, created_at, updated_at FROM posts")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := &proto.Post{}
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Body,
			&post.AuthorId,
			&post.Published,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *store) UpdateLikesCount(ctx context.Context, postId int64) (*proto.Post, error) {
	post := &proto.Post{}

	var currentLikes int64
	err := s.db.QueryRowContext(ctx, "SELECT likes_count FROM posts WHERE id = $1", postId).Scan(&currentLikes)
	if err != nil {
		return nil, err
	}

	newLikesCount := currentLikes + 1

	query := `
	UPDATE posts
	SET 
		likes_count = $2,
		updated_at = NOW()
	WHERE id = $1
	RETURNING id, title, body, author_id, published, created_at, updated_at, likes_count;
	`
	err = s.db.QueryRowContext(ctx, query, postId, newLikesCount).Scan(
		&post.Id,
		&post.Title,
		&post.Body,
		&post.AuthorId,
		&post.Published,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.LikesCount,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}
