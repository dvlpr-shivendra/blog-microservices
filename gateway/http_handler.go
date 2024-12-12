package main

import (
	"net/http"
	"strconv"

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
	mux.HandleFunc("GET /api/posts", h.handleGetPosts)
	mux.HandleFunc("POST /api/posts", h.handleCreatePost)
	mux.HandleFunc("PUT /api/posts/{id}", h.handleUpdatePost)

	mux.HandleFunc("GET /api/comments/{postId}", h.handleGetComments)

	mux.HandleFunc("POST /api/posts/{postId}/likes", h.handleCreateLike)
}

func (h *handler) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.gateway.GetPosts(r.Context())
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, posts)
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

func (h *handler) handleUpdatePost(w http.ResponseWriter, r *http.Request) {

	var body UpdatePostRequest

	if err := common.ReadJSON(r, &body); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.gateway.UpdatePost(r.Context(), &proto.UpdatePostRequest{
		Id:        id,
		Title:     body.Title,
		Body:      body.Body,
		Published: body.Published,
	})

	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusOK, post)
}

func (h *handler) handleGetComments(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.ParseInt(r.PathValue("postId"), 10, 64)

	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, "Something went wrong")
	}

	comments, err := h.gateway.GetComments(r.Context(), postId)
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, comments)
}

func (h *handler) handleCreateLike(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.ParseInt(r.PathValue("postId"), 10, 64)

	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, "Something went wrong")
	}

	comments, err := h.gateway.CreateLike(r.Context(), postId)
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, comments)
}
