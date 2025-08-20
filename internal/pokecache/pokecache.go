package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}

type Cache struct {
	cache 	map[string]cacheEntry
	mux 	*sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: 		make(map[string]cacheEntry),
		mux: 		&sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry := cacheEntry{
		createdAt: 	time.Now(),
		val: 		val,
	}
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, exists := c.cache[key]
	if exists {
		return val.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		cutoffTime := time.Now().Add(-interval)
		for key, value := range c.cache {
			if value.createdAt.Before(cutoffTime) {
				delete(c.cache, key)
			}
		}
	}
}
