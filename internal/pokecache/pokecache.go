package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache 	map[string]cacheEntry
	mu 		*sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	var newCache Cache
	newCache.cache = make(map[string]cacheEntry)
	newCache.mu = &sync.Mutex{}
	ptr_newCache := &newCache
	go ptr_newCache.reapLoop(interval)
	return ptr_newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{createdAt : time.Now(), val : val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cache_entry, ok := c.cache[key]
	return cache_entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.mu.Lock()
		for key, value := range c.cache {
			if time.Since(value.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}