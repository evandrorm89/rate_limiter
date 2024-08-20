package limiter

import (
	"log"
	"time"
)

type Store interface {
	Incr(key string) (int64, error)
	Expire(key string, duration time.Duration) error
	TTL(key string) (time.Duration, error)
}

type RateLimiter struct {
	store         Store
	ipLimit       int64
	tokenLimit    int64
	blockDuration time.Duration
}

func NewRateLimiter(store Store, ipLimit, tokenLimit int64, blockDuration time.Duration) *RateLimiter {
	return &RateLimiter{
		store:         store,
		ipLimit:       ipLimit,
		tokenLimit:    tokenLimit,
		blockDuration: blockDuration,
	}
}

func (rl *RateLimiter) Allow(key string, isToken bool) (bool, error) {
	limit := rl.ipLimit
	if isToken {
		limit = rl.tokenLimit
	}

	count, err := rl.store.Incr(key)
	if err != nil {
		return false, err
	}

	// if count == 1 {
	// 	err = rl.store.Expire(key, time.Second)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// }

	log.Printf("count: %d, limit: %d", count, limit)

	if count > limit {
		ttl, err := rl.store.TTL(key)
		if err != nil {
			return false, err
		}
		if ttl < 0 {
			err = rl.store.Expire(key, rl.blockDuration)
			if err != nil {
				return false, err
			}
		}
		return false, nil
	}

	return true, nil
}
