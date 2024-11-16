package main

import (
	"net/http"

	"blog-services/common"
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
	common.WriteJSON(w, http.StatusOK, "{}")
}
