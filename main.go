package main

import (
	"net/http"
	"simplehttpserver/controllers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	controllers.SetupDatabase("host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable")
	SetRoutes(r)
	http.ListenAndServe(":3333", r)
}
