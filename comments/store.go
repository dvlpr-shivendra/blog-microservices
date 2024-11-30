package main

import (
	"blog-services/common/proto"
	"context"
	"database/sql"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db}
}

func (s *store) GetAll(ctx context.Context, postId int64) ([]*proto.Comment, error) {
	var comments []*proto.Comment

	rows, err := s.db.QueryContext(
		ctx,
		"SELECT id, body, user_id, post_id, created_at, updated_at FROM comments WHERE post_id = $1",
		postId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := &proto.Comment{}
		err := rows.Scan(
			&comment.Id,
			&comment.Body,
			&comment.UserId,
			&comment.PostId,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
