package services

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/maypok86/otter"
)

// ErrCacheMiss indicates that the requested key does not exist in the cache
var ErrCacheMiss = errors.New("cache miss")

type (
	// CacheStore provides an interface for cache storage
	CacheStore interface {
		// get attempts to get a cached value
		get(context.Context, *CacheGetOp) (any, error)

		// set attempts to set an entry in the cache
		set(context.Context, *CacheSetOp) error

		// flush removes a given key and/or tags from the cache
		flush(context.Context, *CacheFlushOp) error

		// close shuts down the cache storage
		close()
	}

	// CacheClient is the client that allows you to interact with the cache
	CacheClient struct {
		// store holds the Cache storage
		store CacheStore
	}

	// CacheSetOp handles chaining a set operation
	CacheSetOp struct {
		client     *CacheClient
		key        string
		group      string
		data       any
		expiration time.Duration
		tags       []string
	}

	// CacheGetOp handles chaining a get operation
	CacheGetOp struct {
		client *CacheClient
		key    string
		group  string
	}

	// CacheFlushOp handles chaining a flush operation
	CacheFlushOp struct {
		client *CacheClient
		key    string
		group  string
		tags   []string
	}

	// inMemoryCacheStore is a cache store implementation in memory
	inMemoryCacheStore struct {
		store    *otter.CacheWithVariableTTL[string, any]
		tagIndex *tagIndex
	}

	// tagIndex maintains an index to support cache tags for in-memory cache stores.
	// There is a performance and memory impact to using cache tags since set and get operations using tags will require
	// locking, and we need to keep track of this index in order to keep everything in sync.
	// If using something like Redis for caching, you can leverage sets to store the index.
	// Cache tags can be useful and convenient, so you should decide if your app benefits enough from this.
	// As it stands there, there is no limiting how much memory this will consume and it will track all keys
	// and tags added and removed from the cache.
	tagIndex struct {
		sync.Mutex
		tags map[string]map[string]struct{} // tag->keys
		keys map[string]map[string]struct{} // key->tags
	}
)

// NewCacheClient creates a new cache client
func NewCacheClient(store CacheStore) *CacheClient {
	return &CacheClient{store: store}
}

// Close closes the connection to the cache
func (c *CacheClient) Close() {
	c.store.close()
}

// Set creates a cache set operation
func (c *CacheClient) Set() *CacheSetOp {
	return &CacheSetOp{
		client: c,
	}
}

// Get creates a cache get operation
func (c *CacheClient) Get() *CacheGetOp {
	return &CacheGetOp{
		client: c,
	}
}

// Flush creates a cache flush operation
func (c *CacheClient) Flush() *CacheFlushOp {
	return &CacheFlushOp{
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
func (c *CacheSetOp) Key(key string) *CacheSetOp {
	c.key = key
	return c
}

// Group sets the cache group
func (c *CacheSetOp) Group(group string) *CacheSetOp {
	c.group = group
	return c
}

// Data sets the data to cache
func (c *CacheSetOp) Data(data any) *CacheSetOp {
	c.data = data
	return c
}

// Expiration sets the expiration duration of the cached data
func (c *CacheSetOp) Expiration(expiration time.Duration) *CacheSetOp {
	c.expiration = expiration
	return c
}

// Tags sets the cache tags
func (c *CacheSetOp) Tags(tags ...string) *CacheSetOp {
	c.tags = tags
	return c
}

// Save saves the data in the cache
func (c *CacheSetOp) Save(ctx context.Context) error {
	switch {
	case c.key == "":
		return errors.New("no cache key specified")
	case c.data == nil:
		return errors.New("no cache data specified")
	case c.expiration == 0:
		return errors.New("no cache expiration specified")
	}

	return c.client.store.set(ctx, c)
}

// Key sets the cache key
func (c *CacheGetOp) Key(key string) *CacheGetOp {
	c.key = key
	return c
}

// Group sets the cache group
func (c *CacheGetOp) Group(group string) *CacheGetOp {
	c.group = group
	return c
}

// Fetch fetches the data from the cache
func (c *CacheGetOp) Fetch(ctx context.Context) (any, error) {
	if c.key == "" {
		return nil, errors.New("no cache key specified")
	}

	return c.client.store.get(ctx, c)
}

// Key sets the cache key
func (c *CacheFlushOp) Key(key string) *CacheFlushOp {
	c.key = key
	return c
}

// Group sets the cache group
func (c *CacheFlushOp) Group(group string) *CacheFlushOp {
	c.group = group
	return c
}

// Tags sets the cache tags
func (c *CacheFlushOp) Tags(tags ...string) *CacheFlushOp {
	c.tags = tags
	return c
}

// Execute flushes the data from the cache
func (c *CacheFlushOp) Execute(ctx context.Context) error {
	return c.client.store.flush(ctx, c)
}

// newInMemoryCache creates a new in-memory CacheStore
func newInMemoryCache(capacity int) (CacheStore, error) {
	s := &inMemoryCacheStore{
		tagIndex: newTagIndex(),
	}

	store, err := otter.MustBuilder[string, any](capacity).
		WithVariableTTL().
		DeletionListener(func(key string, value any, cause otter.DeletionCause) {
			s.tagIndex.purgeKeys(key)
		}).
		Build()

	if err != nil {
		return nil, err
	}

	s.store = &store

	return s, nil
}

func (s *inMemoryCacheStore) get(_ context.Context, op *CacheGetOp) (any, error) {
	v, exists := s.store.Get(op.client.cacheKey(op.group, op.key))

	if !exists {
		return nil, ErrCacheMiss
	}

	return v, nil
}

func (s *inMemoryCacheStore) set(_ context.Context, op *CacheSetOp) error {
	key := op.client.cacheKey(op.group, op.key)

	added := s.store.Set(
		key,
		op.data,
		op.expiration,
	)

	if len(op.tags) > 0 {
		s.tagIndex.setTags(key, op.tags...)
	}

	if !added {
		return errors.New("cache set failed")
	}

	return nil
}

func (s *inMemoryCacheStore) flush(_ context.Context, op *CacheFlushOp) error {
	keys := make([]string, 0)

	if key := op.client.cacheKey(op.group, op.key); key != "" {
		keys = append(keys, key)
	}

	if len(op.tags) > 0 {
		keys = append(keys, s.tagIndex.purgeTags(op.tags...)...)
	}

	for _, key := range keys {
		s.store.Delete(key)
	}

	s.tagIndex.purgeKeys(keys...)

	return nil
}

func (s *inMemoryCacheStore) close() {
	s.store.Close()
}

func newTagIndex() *tagIndex {
	return &tagIndex{
		tags: make(map[string]map[string]struct{}),
		keys: make(map[string]map[string]struct{}),
	}
}

func (i *tagIndex) setTags(key string, tags ...string) {
	i.Lock()
	defer i.Unlock()

	if _, exists := i.keys[key]; !exists {
		i.keys[key] = make(map[string]struct{})
	}

	for _, tag := range tags {
		if _, exists := i.tags[tag]; !exists {
			i.tags[tag] = make(map[string]struct{})
		}
		i.tags[tag][key] = struct{}{}
		i.keys[key][tag] = struct{}{}
	}
}

func (i *tagIndex) purgeTags(tags ...string) []string {
	i.Lock()
	defer i.Unlock()

	keys := make([]string, 0)

	for _, tag := range tags {
		tagKeys := i.tags[tag]
		delete(i.tags, tag)

		for key := range tagKeys {
			delete(i.keys[key], tag)
			if len(i.keys[key]) == 0 {
				delete(i.keys, key)
			}

			keys = append(keys, key)
		}
	}

	return keys
}

func (i *tagIndex) purgeKeys(keys ...string) {
	i.Lock()
	defer i.Unlock()

	for _, key := range keys {
		keyTags := i.keys[key]
		delete(i.keys, key)

		for tag := range keyTags {
			delete(i.tags[tag], key)
			if len(i.tags[tag]) == 0 {
				delete(i.tags, tag)
			}
		}
	}
}
