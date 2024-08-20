package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(redisURL string) (*RedisStore, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	return &RedisStore{client: client}, nil
}

func (rs *RedisStore) Incr(key string) (int64, error) {
	ctx := context.Background()
	return rs.client.Incr(ctx, key).Result()
}

func (rs *RedisStore) Expire(key string, duration time.Duration) error {
	ctx := context.Background()
	return rs.client.Expire(ctx, key, duration).Err()
}

func (rs *RedisStore) TTL(key string) (time.Duration, error) {
	ctx := context.Background()
	return rs.client.TTL(ctx, key).Result()
}
