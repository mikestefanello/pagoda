package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/marshaler"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"github.com/mikestefanello/pagoda/config"
)

type (
	// CacheClient is the client that allows you to interact with the cache
	CacheClient struct {
		// Client stores the client to the underlying cache service
		Client *redis.Client

		// cache stores the cache interface
		cache *cache.Cache
	}

	// cacheSet handles chaining a set operation
	cacheSet struct {
		client     *CacheClient
		key        string
		group      string
		data       interface{}
		expiration time.Duration
		tags       []string
	}

	// cacheGet handles chaining a get operation
	cacheGet struct {
		client   *CacheClient
		key      string
		group    string
		dataType interface{}
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
	// Determine the database based on the environment
	db := cfg.Cache.Database
	if cfg.App.Environment == config.EnvTest {
		db = cfg.Cache.TestDatabase
	}

	// Connect to the cache
	c := &CacheClient{}
	c.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Hostname, cfg.Cache.Port),
		Password: cfg.Cache.Password,
		DB:       db,
	})
	if _, err := c.Client.Ping(context.Background()).Result(); err != nil {
		return c, err
	}

	// Flush the database if this is the test environment
	if cfg.App.Environment == config.EnvTest {
		if err := c.Client.FlushDB(context.Background()).Err(); err != nil {
			return c, err
		}
	}

	cacheStore := store.NewRedis(c.Client, nil)
	c.cache = cache.New(cacheStore)
	return c, nil
}

// Close closes the connection to the cache
func (c *CacheClient) Close() error {
	return c.Client.Close()
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
func (c *cacheSet) Data(data interface{}) *cacheSet {
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

	opts := &store.Options{
		Expiration: c.expiration,
		Tags:       c.tags,
	}

	return marshaler.
		New(c.client.cache).
		Set(ctx, c.client.cacheKey(c.group, c.key), c.data, opts)
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

// Type sets the expected Go type of the data being retrieved from the cache
func (c *cacheGet) Type(expectedType interface{}) *cacheGet {
	c.dataType = expectedType
	return c
}

// Fetch fetches the data from the cache
func (c *cacheGet) Fetch(ctx context.Context) (interface{}, error) {
	if c.key == "" {
		return nil, errors.New("no cache key specified")
	}

	return marshaler.New(c.client.cache).Get(
		ctx,
		c.client.cacheKey(c.group, c.key),
		c.dataType,
	)
}

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
func (c *cacheFlush) Execute(ctx context.Context) error {
	if len(c.tags) > 0 {
		if err := c.client.cache.Invalidate(ctx, store.InvalidateOptions{
			Tags: c.tags,
		}); err != nil {
			return err
		}
	}

	if c.key != "" {
		return c.client.cache.Delete(ctx, c.client.cacheKey(c.group, c.key))
	}

	return nil
}
