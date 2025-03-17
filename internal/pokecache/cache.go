package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}
	c.reapLoop(duration)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exist := c.entries[key]
	return entry.val, exist
}

// duration * time.Second
func (c *Cache) reapLoop(duration time.Duration) {
	go func() {
		for t := range time.Tick(duration) {
			c.mu.Lock()
			for key, entry := range c.entries {
				if t.Sub(entry.createdAt) > duration {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
