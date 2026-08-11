package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheContent map[string]cacheEntry
	mu           sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	// Create a new Cache instance with an empty cacheContent map and start the reapLoop in a separate goroutine
	c := &Cache{
		cacheContent: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Add a new cache entry with the current time and the provided value to the cacheContent map
	c.cacheContent[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Retrieve a cache entry from the cacheContent map based on the provided key
	entry, ok := c.cacheContent[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Start a ticker that triggers every 'interval' duration to clean up expired cache entries
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheContent {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheContent, key)
			}
		}
		c.mu.Unlock()
	}
}
