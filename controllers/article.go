package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"simplehttpserver/models"

	"github.com/go-chi/chi/v5"
)

// ListArticles retrieves all articles from the database
func ListArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article

	// Query all articles from the database
	if err := db.Find(&articles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ListArticlesByDate(w http.ResponseWriter, r *http.Request) {
	// Extract date parameters from URL
	month := chi.URLParam(r, "month")
	day := chi.URLParam(r, "day")
	year := chi.URLParam(r, "year")

	// Construct the date string (format: YYYY-MM-DD)
	dateStr := year + "-" + month + "-" + day

	var articles []models.Article

	// Query articles by date
	if err := db.Where("DATE(created_at) = ?", dateStr).Find(&articles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	// Decode JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate required fields (optional but recommended)
	if article.Title == "" || article.Content == "" {
		http.Error(w, "Title and Content are required", http.StatusBadRequest)
		return
	}

	// Create the article in the database
	if err := db.Create(&article).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Return the created article
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	// Get article ID from URL parameter
	articleID := chi.URLParam(r, "articleID")

	// Find existing article
	var article models.Article
	if err := db.First(&article, articleID).Error; err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Decode JSON from request body for updates
	var updates models.Article
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Update only the fields that are provided
	if updates.Title != "" {
		article.Title = updates.Title
	}
	if updates.Content != "" {
		article.Content = updates.Content
	}
	if updates.Slug != "" {
		article.Slug = updates.Slug
	}
	if updates.Author != "" {
		article.Author = updates.Author
	}

	// Save the updated article to the database
	if err := db.Save(&article).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Return the updated article
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// Get article ID from URL parameter
	articleID := chi.URLParam(r, "articleID")

	// Find existing article
	var article models.Article
	if err := db.First(&article, articleID).Error; err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Delete the article from the database
	if err := db.Delete(&article).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return 204 No Content on successful deletion
	w.WriteHeader(http.StatusNoContent)
}

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin"))
	})
	return r
}
