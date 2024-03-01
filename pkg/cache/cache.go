package cache

import "context"

type Cache[K comparable, V any] interface {
	Set(ctx context.Context, key K, value *V) error
	Get(ctx context.Context, key K) (*V, error)
}
