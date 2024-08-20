package limiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockRedisStore simula o comportamento de RedisStore
type MockRedisStore struct {
	store map[string]int64
}

func (m *MockRedisStore) Incr(key string) (int64, error) {
	m.store[key]++
	return m.store[key], nil
}

func (m *MockRedisStore) TTL(key string) (time.Duration, error) {
	// Simular o TTL retornando, por exemplo, 10 segundos
	return 10 * time.Second, nil
}

func TestRateLimiterAllow(t *testing.T) {
	mockStore := &MockRedisStore{store: make(map[string]int64)}
	rl := NewRateLimiter(mockStore, 5, 10, 60*time.Second) // ipLimit, tokenLimit, blockDuration

	// Teste de caso onde a contagem é permitida
	allowed, err := rl.Allow("ip:127.0.0.1", false)
	assert.NoError(t, err)
	assert.True(t, allowed)

	// Teste de caso onde o limite é atingido
	for i := 0; i < 5; i++ {
		rl.Allow("ip:127.0.0.1", false)
	}
	allowed, err = rl.Allow("ip:127.0.0.1", false)
	assert.NoError(t, err)
	assert.False(t, allowed)
}
