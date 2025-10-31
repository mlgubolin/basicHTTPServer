package simplehttpserver

import (
	"net/http"

	"simplehttpserver/controllers"

	"github.com/go-chi/chi/v5"
)

// paginate is a stub middleware handler
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pagination logic would go here
		next.ServeHTTP(w, r)
	})
}

func SetRoutes(r *chi.Mux) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	// RESTy routes for "articles" resource
	r.Route("/articles", func(r chi.Router) {
		r.With(paginate).Get("/", controllers.ListArticles)                           // GET /articles
		r.With(paginate).Get("/{month}-{day}-{year}", controllers.ListArticlesByDate) // GET /articles/01-16-2017

		r.Post("/", controllers.CreateArticle)       // POST /articles
		r.Get("/search", controllers.SearchArticles) // GET /articles/search

		// Regexp url parameters:
		r.Get("/{articleSlug:[a-z-]+}", controllers.GetArticleBySlug) // GET /articles/home-is-toronto

		// Subrouters:
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(controllers.ArticleCtx)
			r.Get("/", controllers.GetArticle)       // GET /articles/123
			r.Put("/", controllers.UpdateArticle)    // PUT /articles/123
			r.Delete("/", controllers.DeleteArticle) // DELETE /articles/123
		})
	})

	// Mount the admin sub-router
	r.Mount("/admin", controllers.AdminRouter())
}
