package main

import (
	"log"
	"time"

	"github.com/evandrorm89/rate_limiter/internal/config"
	"github.com/evandrorm89/rate_limiter/internal/limiter"
	"github.com/evandrorm89/rate_limiter/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	store, err := limiter.NewRedisStore(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to create Redis store: %v", err)
	}

	rl := limiter.NewRateLimiter(store, cfg.RateLimitIP, cfg.RateLimitToken, time.Duration(cfg.BlockDuration)*time.Second)

	srv := server.NewServer(rl)

	log.Printf("Server is running on http://localhost%s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
