package main

import (
	"blog-services/common"
	"blog-services/common/proto"
	"blog-services/gateway/gateway"
	"fmt"
	"net/http"
	"strconv"

	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("http")
	ctx, span := tr.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()
	posts, err := h.gateway.GetPosts(ctx)
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, posts)
}

func (h *handler) handleCreatePost(w http.ResponseWriter, r *http.Request) {

	tr := otel.Tracer("http")
	ctx, span := tr.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	var body CreatePostRequest

	if err := common.ReadJSON(r, &body); err != nil {
		common.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.gateway.CreatePost(ctx, &proto.CreatePostRequest{
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
	tr := otel.Tracer("http")
	ctx, span := tr.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()
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

	post, err := h.gateway.UpdatePost(ctx, &proto.UpdatePostRequest{
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
	tr := otel.Tracer("http")
	ctx, span := tr.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()
	postId, err := strconv.ParseInt(r.PathValue("postId"), 10, 64)

	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, "Something went wrong")
	}

	comments, err := h.gateway.GetComments(ctx, postId)
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, comments)
}

func (h *handler) handleCreateLike(w http.ResponseWriter, r *http.Request) {
	tr := otel.Tracer("http")
	ctx, span := tr.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()
	postId, err := strconv.ParseInt(r.PathValue("postId"), 10, 64)

	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, "Something went wrong")
	}

	comments, err := h.gateway.CreateLike(ctx, postId)
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSON(w, http.StatusOK, comments)
}
