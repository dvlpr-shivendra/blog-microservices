package main

import (
	"blog-services/common/proto"
	"context"
)

type CommentsService interface {
	GetComments(context.Context, int64) ([]*proto.Comment, error)
}

type CommentsStore interface {
	GetAll(context.Context, int64) ([]*proto.Comment, error)
}
