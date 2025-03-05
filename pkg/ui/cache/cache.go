package cache

import (
	"bytes"
	"sync"

	"maragu.dev/gomponents"
)

var (
	// cache stores a cache of assembled components by key.
	cache = make(map[string]gomponents.Node)

	// mu handles concurrent access to the cache.
	mu sync.RWMutex
)

// Set sets a given renderable node in the cache with a given key.
// You should only cache nodes that are entirely static.
// This will panic if the node fails to render.
//
// To optimize performance, the node will be rendered and converted to a Raw component so the assembly and rendering
// of the entire, nested node only has to execute once. This can eliminate countless function calls and string building.
// It's worth noting that this performance optimization is, in most cases, entirely unnecessary, but it's easy to do
// and realize some performance gains. In my very limited testing, gomponents actually outperformed Go templates in
// many areas; but not all. The results were still very close and my limited testing is in no way definitive.
//
// In most applications, these slight differences in nanoseconds and bytes allocated will almost never matter or even
// be noticeable, but it's good to be aware of them; and it's fun to address them. In looking at the example layouts
// provided, I noticed that a lot of nested function calls and string building was happening on every single page load
// just to re-render static HTML such as the navbar and the search form/modal. Benchmarks quickly revealed that caching
// those high-level nodes made a significant difference in speed and memory allocations. Going further, I thought that
// with the entire node cached, you still have to render the entire nested structure each time it's used, so that is why
// this will render them upfront, then cache. If my few examples have a handful of static nodes, I assume most full
// applications will have many, so maybe this is useful.
func Set(key string, node gomponents.Node) {
	buf := bytes.NewBuffer(nil)
	if err := node.Render(buf); err != nil {
		panic(err)
	}

	mu.Lock()
	defer mu.Unlock()
	cache[key] = gomponents.Raw(buf.String())
}

// Get returns the node cached under the provided key, if one exists.
func Get(key string) gomponents.Node {
	mu.RLock()
	defer mu.RUnlock()
	return cache[key]
}

// SetIfNotExists will return the cached Node for the key, if it exists, otherwise it will use the provided callback
// function to generate the node and cache it.
func SetIfNotExists(key string, gen func() gomponents.Node) gomponents.Node {
	if n := Get(key); n != nil {
		return n
	}

	n := gen()
	Set(key, n)
	return n
}
