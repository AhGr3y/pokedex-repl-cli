package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
	stop     chan struct{}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return nil, ok
	}
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {

	// Create a ticker that returns the current time at the point of every interval
	ticker := time.NewTicker(interval)

	// Infinite loop
	for t := range ticker.C {

		func() {
			c.mu.Lock()
			defer c.mu.Unlock()

			// Delete entries that are older than the set interval
			for key, entry := range c.cacheMap {
				entryCreatedOn := entry.createdAt
				timeDiff := t.Sub(entryCreatedOn)
				if timeDiff >= interval {
					delete(c.cacheMap, key)
				}
			}
		}()

	}

}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheMap: map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		stop:     make(chan struct{}),
	}
	go cache.reapLoop(interval)
	return &cache
}
