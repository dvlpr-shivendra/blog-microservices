package main

import (
	"blog-services/common/proto"
	"context"
)

type CommentsService interface {
	GetComments(context.Context) ([]*proto.Comment, error)
}

type CommentsStore interface {
	GetAll(int) ([]*proto.Comment, error)
}
