package cache_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tuxoo/idler/pkg/cache"
	"testing"
)

const (
	notFoundInCacheText = "cache: value not found"
	testUserKey         = 0
)

type user struct {
	id   int
	name string
}

func TestCache(t *testing.T) {
	ctx := context.TODO()

	testCache := cache.NewMemoryCache[int, user]()
	_, err := testCache.Get(ctx, testUserKey)
	assert.Equal(t, notFoundInCacheText, err.Error(),
		fmt.Sprintf("Incorrect result. Expect %s, got %s", notFoundInCacheText, err.Error()))

	testUser := user{
		id:   13,
		name: "Alex",
	}

	testCache.Set(ctx, testUserKey, &testUser)

	cacheableUser, err := testCache.Get(ctx, testUserKey)
	assert.Nil(t, err, "Incorrect result. Expect nil")

	assert.Equal(t, testUser.id, cacheableUser.id,
		fmt.Sprintf("Incorrect result. Expect %d, got %d", testUser.id, cacheableUser.id))

	assert.Equal(t, testUser.name, cacheableUser.name,
		fmt.Sprintf("Incorrect result. Expect %s, got %s", testUser.name, cacheableUser.name))
}
