package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache[K string, V any] struct {
	RedisClient *redis.Client
	Expires     time.Duration
}

func NewRedisCache[K string, V any](client *redis.Client, expires time.Duration) *RedisCache[K, V] {
	return &RedisCache[K, V]{
		RedisClient: client,
		Expires:     expires,
	}
}

func (c *RedisCache[K, V]) Set(ctx context.Context, key K, value *V) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}

	c.RedisClient.Set(ctx, string(key), val, c.Expires*time.Second)
	return nil
}

func (c *RedisCache[K, V]) Get(ctx context.Context, key K) (*V, error) {
	strValue, err := c.RedisClient.Get(ctx, string(key)).Result()
	if err != nil {
		return nil, err
	}

	var value V
	if err := json.Unmarshal([]byte(strValue), &value); err != nil {
		panic(err)
	}

	return &value, err
}
