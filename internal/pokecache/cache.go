package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
  createdAt time.Time
  val []byte
}

type Cache struct {
  interval time.Duration
  cache    map[string]cacheEntry
  mu       sync.RWMutex
}

func NewCache(duration time.Duration) *Cache {
  new_cache := &Cache{
    interval: duration,
    cache: make(map[string]cacheEntry),
  }

  go new_cache.reapLoop()

  return new_cache
}

func (c *Cache) Add(key string, val []byte) {
  c.mu.Lock()
  defer c.mu.Unlock()  

  c.cache[key] = cacheEntry{
    createdAt: time.Now(), 
    val:       val,
  }
}

func (c *Cache) Get(key string) ([]byte, bool) {
  c.mu.RLock()
  defer c.mu.RUnlock()

  if c.cache[key].val != nil {
    return c.cache[key].val, true
  }
  return []byte{}, false
}

func (c *Cache) reapLoop() {
  ticker := time.NewTicker(1 * time.Millisecond)
  defer ticker.Stop()
  
  for range ticker.C {
    c.mu.Lock()
    for key, entry := range c.cache {
      if time.Since(entry.createdAt) > c.interval {
        delete(c.cache, key)
      }
    }
    c.mu.Unlock()
  }
}

