package server

import (
	"net/http"

	"github.com/evandrorm89/rate_limiter/internal/limiter"
	middlewarerl "github.com/evandrorm89/rate_limiter/internal/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewServer(rl *limiter.RateLimiter) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewarerl.RateLimiterMiddleware(rl))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// return &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: middlewarerl.RateLimiterMiddleware(rl)(mux),
	// }
}
