package cache

import (
	"context"
	"github.com/bluele/gcache"
	"time"
)

type GCache[K comparable, V any] struct {
	cache gcache.Cache
}

func NewGCache[K comparable, V any](size int, expires time.Duration) *GCache[K, V] {
	return &GCache[K, V]{
		cache: gcache.New(size).Expiration(expires).LRU().Build(),
	}
}

func (c *GCache[K, V]) Get(ctx context.Context, key K) (*V, error) {
	value, err := c.cache.Get(key)
	if err != nil {
		return nil, err
	}
	return value.(*V), nil
}

func (c *GCache[K, V]) Set(ctx context.Context, key K, value *V) error {
	return c.cache.Set(key, value)
}
