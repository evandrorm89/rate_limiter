package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/evandrorm89/rate_limiter/internal/limiter"
)

func RateLimiterMiddleware(rl *limiter.RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var key string
			var isToken bool

			token := r.Header.Get("API_KEY")
			if token != "" {
				key = "token:" + token
				isToken = true
			} else {
				key = "ip:" + strings.Split(r.RemoteAddr, ":")[0]
			}

			allowed, err := rl.Allow(key, isToken)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println(err)
				return
			}

			if !allowed {
				http.Error(w, "You have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
