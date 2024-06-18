package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/maypok86/otter"
	"github.com/mikestefanello/pagoda/config"
)

type (
	// CacheClient is the client that allows you to interact with the cache
	CacheClient struct {
		// cache stores the cache interface
		cache *otter.CacheWithVariableTTL[string, any]
	}

	// cacheSet handles chaining a set operation
	cacheSet struct {
		client     *CacheClient
		key        string
		group      string
		data       any
		expiration time.Duration
		tags       []string
	}

	// cacheGet handles chaining a get operation
	cacheGet struct {
		client *CacheClient
		key    string
		group  string
	}

	// cacheFlush handles chaining a flush operation
	cacheFlush struct {
		client *CacheClient
		key    string
		group  string
		tags   []string
	}
)

// NewCacheClient creates a new cache client
func NewCacheClient(cfg *config.Config) (*CacheClient, error) {
	cache, err := otter.MustBuilder[string, any](10000).
		WithVariableTTL().
		DeletionListener(func(key string, value any, cause otter.DeletionCause) {
			// todo
		}).
		Build()

	if err != nil {
		return nil, err
	}

	return &CacheClient{cache: &cache}, nil
}

// Close closes the connection to the cache
func (c *CacheClient) Close() {
	c.cache.Close()
}

// Set creates a cache set operation
func (c *CacheClient) Set() *cacheSet {
	return &cacheSet{
		client: c,
	}
}

// Get creates a cache get operation
func (c *CacheClient) Get() *cacheGet {
	return &cacheGet{
		client: c,
	}
}

// Flush creates a cache flush operation
func (c *CacheClient) Flush() *cacheFlush {
	return &cacheFlush{
		client: c,
	}
}

// cacheKey formats a cache key with an optional group
func (c *CacheClient) cacheKey(group, key string) string {
	if group != "" {
		return fmt.Sprintf("%s::%s", group, key)
	}
	return key
}

// Key sets the cache key
func (c *cacheSet) Key(key string) *cacheSet {
	c.key = key
	return c
}

// Group sets the cache group
func (c *cacheSet) Group(group string) *cacheSet {
	c.group = group
	return c
}

// Data sets the data to cache
func (c *cacheSet) Data(data any) *cacheSet {
	c.data = data
	return c
}

// Expiration sets the expiration duration of the cached data
func (c *cacheSet) Expiration(expiration time.Duration) *cacheSet {
	c.expiration = expiration
	return c
}

// Tags sets the cache tags
func (c *cacheSet) Tags(tags ...string) *cacheSet {
	c.tags = tags
	return c
}

// Save saves the data in the cache
func (c *cacheSet) Save(ctx context.Context) error {
	if c.key == "" {
		return errors.New("no cache key specified")
	}

	if c.data == nil {
		return errors.New("no cache data specified")
	}

	c.client.cache.Set(
		c.client.cacheKey(c.group, c.key),
		c.data,
		c.expiration,
	)

	// TODO tags
	return nil
}

// Key sets the cache key
func (c *cacheGet) Key(key string) *cacheGet {
	c.key = key
	return c
}

// Group sets the cache group
func (c *cacheGet) Group(group string) *cacheGet {
	c.group = group
	return c
}

// Fetch fetches the data from the cache
func (c *cacheGet) Fetch(ctx context.Context) (any, error) {
	if c.key == "" {
		return nil, errors.New("no cache key specified")
	}

	v, exists := c.client.cache.Get(c.client.cacheKey(c.group, c.key))

	if !exists {
		return nil, ErrCacheMiss
	}

	return v, nil
}

var ErrCacheMiss = errors.New("cache miss")

// Key sets the cache key
func (c *cacheFlush) Key(key string) *cacheFlush {
	c.key = key
	return c
}

// Group sets the cache group
func (c *cacheFlush) Group(group string) *cacheFlush {
	c.group = group
	return c
}

// Tags sets the cache tags
func (c *cacheFlush) Tags(tags ...string) *cacheFlush {
	c.tags = tags
	return c
}

// Execute flushes the data from the cache
func (c *cacheFlush) Execute(ctx context.Context) {
	// TODO tags

	c.client.cache.Delete(c.client.cacheKey(c.group, c.key))
}
