package main

import "blog-services/common/proto"

type store struct{}

func NewStore() *store {
	return &store{}
}

func (s *store) GetAll(postId int) ([]*proto.Comment, error) {
	return []*proto.Comment{{Body: "Test comment"}}, nil
}
