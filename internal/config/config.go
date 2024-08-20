package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RateLimitIP    int64
	RateLimitToken int64
	BlockDuration  int
	RedisURL       string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// rateLimitIP, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	rateLimitIP, _ := strconv.ParseInt(os.Getenv("RATE_LIMIT_IP"), 0, 64)
	rateLimitToken, _ := strconv.ParseInt(os.Getenv("RATE_LIMIT_TOKEN"), 0, 64)
	// rateLimitToken, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	blockDuration, _ := strconv.Atoi(os.Getenv("BLOCK_DURATION"))

	return &Config{
		RateLimitIP:    rateLimitIP,
		RateLimitToken: rateLimitToken,
		BlockDuration:  blockDuration,
		RedisURL:       os.Getenv("REDIS_URL"),
	}, nil
}
