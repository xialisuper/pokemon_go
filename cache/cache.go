package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	Val       []byte
	CreatedAt time.Time
}

type Cache struct {
	Entries map[string]CacheEntry
	// mutex
	mux *sync.Mutex
}

// NewCache creates a new cache with the given reap interval
func NewCache(interval time.Duration) Cache {
	c := Cache{
		Entries: make(map[string]CacheEntry),
		mux:     &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

// reapLoop periodically cleans cache entries that are older than the given interval
func (c *Cache) reapLoop(interval time.Duration) {
	// use time.Ticker to clean cache every interval
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.CreatedAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.mux.Unlock()
	}
}

// Add adds a new entry to the cache with the given key and value
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.Entries[key] = CacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

// Get retrieves the value of the entry with the given key from the cache
// and returns it along with a boolean indicating whether the entry was found or not.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	if time.Since(entry.CreatedAt) > 10*time.Minute {
		delete(c.Entries, key)
		return nil, false
	}
	return entry.Val, true
}
