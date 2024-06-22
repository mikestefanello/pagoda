package services

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCacheClient(t *testing.T) {
	type cacheTest struct {
		Value string
	}

	// Cache some data
	data := cacheTest{Value: "abcdef"}
	group := "testgroup"
	key := "testkey"
	err := c.Cache.
		Set().
		Group(group).
		Key(key).
		Data(data).
		Expiration(500 * time.Millisecond).
		Save(context.Background())
	require.NoError(t, err)

	// Get the data
	fromCache, err := c.Cache.
		Get().
		Group(group).
		Key(key).
		Fetch(context.Background())
	require.NoError(t, err)
	cast, ok := fromCache.(cacheTest)
	require.True(t, ok)
	assert.Equal(t, data, cast)

	// The same key with the wrong group should fail
	_, err = c.Cache.
		Get().
		Key(key).
		Fetch(context.Background())
	assert.Equal(t, ErrCacheMiss, err)

	// Flush the data
	err = c.Cache.
		Flush().
		Group(group).
		Key(key).
		Execute(context.Background())
	require.NoError(t, err)

	// The data should be gone
	assertFlushed := func(key string) {
		// The data should be gone
		_, err = c.Cache.
			Get().
			Group(group).
			Key(key).
			Fetch(context.Background())
		assert.Equal(t, ErrCacheMiss, err)
	}
	assertFlushed(key)

	// Set with tags
	key = "testkey2"
	err = c.Cache.
		Set().
		Group(group).
		Key(key).
		Data(data).
		Tags("tag1", "tag2").
		Expiration(time.Hour).
		Save(context.Background())
	require.NoError(t, err)

	// Check the tag index
	index := c.Cache.store.(*inMemoryCacheStore).tagIndex
	gk := c.Cache.cacheKey(group, key)
	_, exists := index.tags["tag1"][gk]
	assert.True(t, exists)
	_, exists = index.tags["tag2"][gk]
	assert.True(t, exists)
	_, exists = index.keys[gk]["tag1"]
	assert.True(t, exists)
	_, exists = index.keys[gk]["tag2"]
	assert.True(t, exists)

	// Flush one of tags
	err = c.Cache.
		Flush().
		Tags("tag1").
		Execute(context.Background())
	require.NoError(t, err)

	// The data should be gone
	assertFlushed(key)

	// The index should be empty
	assert.Empty(t, index.tags)
	assert.Empty(t, index.keys)
}
