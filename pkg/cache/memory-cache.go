package cache

import (
	"context"
	"errors"
	"sync"
)

type MemoryCache[K comparable, V any] struct {
	cache map[K]V
	sync.RWMutex
}

func NewMemoryCache[K comparable, V any]() *MemoryCache[K, V] {
	return &MemoryCache[K, V]{
		cache: make(map[K]V),
	}
}

func (c *MemoryCache[K, V]) Set(ctx context.Context, key K, value *V) error {
	c.Lock()
	c.cache[key] = *value
	c.Unlock()
	return nil
}

func (c *MemoryCache[K, V]) Get(ctx context.Context, key K) (*V, error) {
	c.RLock()
	value, exist := c.cache[key]
	c.RUnlock()

	if !exist {
		return &value, errors.New("cache: value not found")
	}

	return &value, nil
}

func (c *MemoryCache[K, V]) Delete(ctx context.Context, key K) {
	c.Lock()
	delete(c.cache, key)
	c.Unlock()
}

func (c *MemoryCache[K, V]) Size(ctx context.Context) int {
	c.RLock()
	size := len(c.cache)
	c.RUnlock()

	return size
}
