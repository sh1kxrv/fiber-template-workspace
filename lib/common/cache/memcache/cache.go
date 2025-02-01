package memcache

import (
	"context"
	"sync"
	"time"
)

// Каждые 25 секунд
var CLEANUP_TIMEOUT = 25 * time.Second

type CacheItem struct {
	Value      any
	Expiration int64
}

// NOTE: Кэширование памятью
type MemoryCache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

func NewMemoryCache(ctx context.Context) *MemoryCache {
	cache := &MemoryCache{
		items: make(map[string]CacheItem),
	}
	go cache.cleanup(ctx)
	return cache
}

func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(ttl).Unix()
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found || (item.Expiration > 0 && item.Expiration < time.Now().Unix()) {
		return nil, false
	}
	return item.Value, true
}

func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *MemoryCache) cleanup(ctx context.Context) {
	ticker := time.NewTicker(CLEANUP_TIMEOUT)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now().Unix()
			for key, item := range c.items {
				if item.Expiration > 0 && item.Expiration < now {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
