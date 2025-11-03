package controllers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler stubs
func ListArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of articles"))
}

func ListArticlesByDate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of articles by date"))
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create new article"))
}

func SearchArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search articles"))
}

func GetArticleBySlug(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get article by slug"))
}

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "article", "article data")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get article"))
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update article"))
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete article"))
}

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin"))
	})
	return r
}
