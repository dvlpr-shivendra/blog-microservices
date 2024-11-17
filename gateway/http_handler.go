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

	var body CreatePostRequest
	
	if err := common.ReadJSON(r, &body); err != nil {
		common.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.gateway.CreatePost(r.Context(), &proto.CreatePostRequest{
		Title: body.Title,
		Body:  body.Body,
	})

	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, post)
}
