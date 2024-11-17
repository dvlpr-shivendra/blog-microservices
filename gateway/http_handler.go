package main

import (
	"net/http"

	"blog-services/common"
	"blog-services/common/proto"
	"blog-services/gateway/gateway"
)

type handler struct {
	gateway gateway.PostsGateway
}

func NewHandler(gateway gateway.PostsGateway) *handler {
	return &handler{gateway}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/posts", h.handleCreatePost)
}

func (h *handler) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	h.gateway.CreatePost(r.Context(), &proto.CreatePostRequest{
		Title: "Test Post",
		Body:  "Test Body",
	})
	common.WriteJSON(w, http.StatusOK, "{}")
}
