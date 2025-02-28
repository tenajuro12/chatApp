package http

import (
	"blogs_service/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	BlogUsecase *usecase.BlogUsecase
}

type createBlogRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var req createBlogRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := h.BlogUsecase.CreateBlog(req.Title, req.Content, req.AuthorName, req.AuthorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Blog created successfully"))

}

func (h *BlogHandler) GetBlog(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing blog id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog id", http.StatusBadRequest)
		return
	}
	blog, err := h.BlogUsecase.GetBlogById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func (h *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing blog id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog id", http.StatusBadRequest)
		return
	}
	err = h.BlogUsecase.DeleteBlog(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog deleted successfully"))
}

type updateBlogRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *BlogHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	var req updateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.BlogUsecase.UpdateBlog(req.ID, req.Title, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog updated successfully"))
}

func (h *BlogHandler) ListBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := h.BlogUsecase.ListBlogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}
